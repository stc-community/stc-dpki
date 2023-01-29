package internal

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/stc-community/stc-dpki/internal/config"
	"github.com/stc-community/stc-dpki/internal/initx"
	"github.com/stc-community/stc-dpki/pkg/errors"
	"github.com/stc-community/stc-dpki/pkg/logger"
)

const AppTlsType = "tls"
const AppOcspType = "ocsp"

type options struct {
	App        string
	ConfigFile string
	Version    string
}

type Option func(*options)

func SetConfigFile(s string) Option {
	return func(o *options) {
		o.ConfigFile = s
	}
}

func SetVersion(s string) Option {
	return func(o *options) {
		o.Version = s
	}
}

func SetAppType(s string) Option {
	return func(o *options) {
		o.App = s
	}
}

func Init(ctx context.Context, opts ...Option) (func(), error) {
	var o options
	for _, opt := range opts {
		opt(&o)
	}

	err := config.MustLoad(o.ConfigFile)
	if err != nil {
		return nil, err
	}
	logger.WithContext(ctx).Printf("Start server,#run_mode %s,#version %s,#pid %d", config.C.RunMode, o.Version, os.Getpid())
	loggerCleanFunc, err := initx.InitLogger()
	if err != nil {
		return nil, err
	}

	err = initx.InitRootCert()
	if err != nil {
		return nil, err
	}

	injector, injectorCleanFunc, err := BuildInjector(ctx)
	if err != nil {
		return nil, err
	}
	// 根证书上链
	err = injector.Router.CertificateAPI.CertificateSrv.CreateRootCa(ctx)
	if err != nil {
		return nil, err
	}
	tlsServerCleanFunc := func() {}
	ocspServerCleanFunc := func() {}
	switch o.App {
	case AppTlsType:
		tlsServerCleanFunc = InitTLSServer(ctx, injector.Engine)
	case AppOcspType:
		ocspServerCleanFunc = InitOCSPServer(ctx, injector.OcspEngine)
	default:
		return nil, errors.New("Unknown app type")
	}
	return func() {
		ocspServerCleanFunc()
		tlsServerCleanFunc()
		injectorCleanFunc()
		loggerCleanFunc()
	}, nil
}

// InitTLSServer TLS 服务
func InitTLSServer(ctx context.Context, handler http.Handler) func() {
	cfg := config.C.TLS
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	srv := &http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	go func() {
		logger.WithContext(ctx).Printf("TLS server is running at %s.", addr)

		var err error
		if cfg.CertFile != "" && cfg.KeyFile != "" {
			srv.TLSConfig = &tls.Config{MinVersion: tls.VersionTLS12}
			err = srv.ListenAndServeTLS(cfg.CertFile, cfg.KeyFile)
		} else {
			err = srv.ListenAndServe()
		}
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}

	}()

	return func() {
		ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(cfg.ShutdownTimeout))
		defer cancel()

		srv.SetKeepAlivesEnabled(false)
		if err := srv.Shutdown(ctx); err != nil {
			logger.WithContext(ctx).Errorf(err.Error())
		}
	}
}

// InitOCSPServer OCSP 服务
func InitOCSPServer(ctx context.Context, handler http.Handler) func() {
	cfg := config.C.OCSP
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	srv := &http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	go func() {
		logger.WithContext(ctx).Printf("OCSP server is running at %s.", addr)
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()

	return func() {
		ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(cfg.ShutdownTimeout))
		defer cancel()

		srv.SetKeepAlivesEnabled(false)
		if err := srv.Shutdown(ctx); err != nil {
			logger.WithContext(ctx).Errorf(err.Error())
		}
	}
}

func Run(ctx context.Context, opts ...Option) error {
	state := 1
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	cleanFunc, err := Init(ctx, opts...)
	if err != nil {
		return err
	}

EXIT:
	for {
		sig := <-sc
		logger.WithContext(ctx).Infof("Receive signal[%s]", sig.String())
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			state = 0
			break EXIT
		case syscall.SIGHUP:
		default:
			break EXIT
		}
	}

	cleanFunc()
	logger.WithContext(ctx).Infof("Server exit")
	time.Sleep(time.Second)
	os.Exit(state)
	return nil
}
