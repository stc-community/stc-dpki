package main

import (
	"context"
	"github.com/stc-community/stc-dpki/pkg/contract"
	"os"

	"github.com/stc-community/stc-dpki/internal"
	"github.com/stc-community/stc-dpki/pkg/logger"
	"github.com/urfave/cli/v2"
)

var VERSION = "0.0.1"

func main() {
	ctx := logger.NewTagContext(context.Background(), "__main__")

	app := cli.NewApp()
	app.Name = "ca"
	app.Version = VERSION
	app.Usage = "CA PKI"
	app.Commands = []*cli.Command{
		newTlsCmd(ctx),
		newDeployContractCmd(ctx),
		//newOcspCmd(ctx),
	}
	err := app.Run(os.Args)
	if err != nil {
		logger.WithContext(ctx).Errorf(err.Error())
	}
}

func newTlsCmd(ctx context.Context) *cli.Command {
	return &cli.Command{
		Name:  "tls",
		Usage: "Run TLS server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "conf",
				Aliases:  []string{"c"},
				Usage:    "App configuration file(.json,.yaml,.toml)",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			return internal.Run(ctx,
				internal.SetConfigFile(c.String("conf")),
				internal.SetVersion(VERSION),
				internal.SetAppType(c.Command.Name),
			)
		},
	}
}

// newDeployContractCmd 部署合约
func newDeployContractCmd(ctx context.Context) *cli.Command {
	return &cli.Command{
		Name:  "deploy",
		Usage: "deploy contract",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "url",
				Aliases:  []string{"u"},
				Usage:    "rpc url",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "key",
				Aliases:  []string{"k"},
				Usage:    "private key",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			return contract.DeployContract(ctx, &contract.Config{
				PrivateKey: c.String("key"),
				RpcUrl:     c.String("url"),
			})
		},
	}
}

func newOcspCmd(ctx context.Context) *cli.Command {
	return &cli.Command{
		Name:  "ocsp",
		Usage: "Run OCSP server",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "conf",
				Aliases:  []string{"c"},
				Usage:    "App configuration file(.json,.yaml,.toml)",
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			return internal.Run(ctx,
				internal.SetConfigFile(c.String("conf")),
				internal.SetVersion(VERSION),
				internal.SetAppType(c.Command.Name),
			)
		},
	}
}
