package app

import (
	"context"
	"exchange-service/internal/configuration"
	"exchange-service/internal/discovery"
	"exchange-service/internal/service/exchange"
	"fmt"
	"log/slog"

	"go.uber.org/fx"
)

type AppSetupManager interface {
	Setup() error
	Shutdown() error
}
type appSetupManagerImpl struct {
	exchange  exchange.ExchangeSetupService
	discovery discovery.ServiceDiscovery
	app       RESTApp
}

func NewAppSetupManager(exchange exchange.ExchangeSetupService, discovery discovery.ServiceDiscovery, app RESTApp) AppSetupManager {
	return appSetupManagerImpl{exchange, discovery, app}
}

func (a appSetupManagerImpl) Setup() error {
	err := a.exchange.Setup()
	if err != nil {
		return err
	}

	err = a.app.setup()
	if err != nil {
		return err
	}

	err = a.discovery.Register()
	if err != nil {
		return err
	}

	return nil
}

func (a appSetupManagerImpl) Shutdown() error {
	return nil
}

func ManageLifeCycle(lc fx.Lifecycle, config configuration.Config, log *slog.Logger, app RESTApp, manager AppSetupManager) {
	err := manager.Setup()
	if err != nil {
		panic(err)
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			log.Info("Starting the server")
			go app.server().Start(fmt.Sprintf(":%v", config.Server.Port))
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info("Shuting down the server")
			err := manager.Shutdown()
			if err != nil {
				return err
			}
			return app.server().Shutdown(ctx)
		},
	})
}
