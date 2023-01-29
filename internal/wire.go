//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package internal

import (
	"context"
	"github.com/stc-community/fvm-dpki/internal/service"
	"github.com/google/wire"

	"github.com/stc-community/fvm-dpki/internal/api"
	"github.com/stc-community/fvm-dpki/internal/dao"
	"github.com/stc-community/fvm-dpki/internal/initx"
	"github.com/stc-community/fvm-dpki/internal/router"
)

func BuildInjector(ctx context.Context) (*Injector, func(), error) {
	wire.Build(
		initx.InitCfssl,
		initx.InitStorage,
		initx.InitEthClient,
		//initx.InitCrdtKv,
		initx.InitOcspCache,
		dao.RepoSet,
		service.ServiceSet,
		InitGinEngine,
		InitOcspEngine,
		api.APISet,
		router.RouterSet,
		InjectorSet,
	)
	return new(Injector), nil, nil
}
