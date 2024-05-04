package main

import (
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"

	"exchange-service/internal/api"
	"exchange-service/internal/app"
	"exchange-service/internal/configuration"
	"exchange-service/internal/database"
	"exchange-service/internal/discovery"
	"exchange-service/internal/message"
	"exchange-service/internal/persistence"
	"exchange-service/internal/sdk"
	"exchange-service/internal/service"
	"exchange-service/internal/service/currency"
	"exchange-service/internal/service/exchange"
	"exchange-service/internal/util"
)

func main() {
	fx.New(
		fx.Provide(configuration.NewConfig),
		fx.Provide(app.NewLogger),
		fx.Provide(app.NewFxLogger),
		fx.Provide(app.ProvideEcho),
		fx.Provide(app.NewAppSetupManager),
		fx.Provide(discovery.NewServiceDiscovery),
		fx.Provide(util.NewValidator),
		fx.Provide(util.NewHttpClient),
		fx.Provide(api.NewHTTPErrorHandler),
		fx.Provide(message.NewMessageQueueClient),
		fx.Provide(message.NewMessageHandler),
		fx.Provide(database.NewDatabaseConnection),
		fx.Provide(persistence.NewCurrencyDAO),
		fx.Provide(persistence.NewPairDAO),
		fx.Provide(persistence.NewExchangeDAO),
		fx.Provide(sdk.NewWallexClient),
		fx.Provide(sdk.NewExchangeAPIClientFactory),
		fx.Provide(service.NewOrderService),
		fx.Provide(currency.NewCurrencyService),
		fx.Provide(exchange.NewPairService),
		fx.Provide(exchange.NewPriceService),
		fx.Provide(exchange.NewExchangeService),
		fx.Provide(exchange.NewOrderService),
		fx.Provide(exchange.NewExchangeSetupService),
		asHandler(api.NewHealthCheck),
		asHandler(api.NewExchangeHandler),
		asHandler(api.NewPairHandler),
		asHandler(api.NewCurrencyHandler),
		asHandler(api.NewPriceHandler),
		asHandler(api.NewOrderHandler),
		fx.Provide(fx.Annotate(
			app.NewRestApp,
			fx.ParamTags(`group:"handlers"`),
		)),
		fx.WithLogger(func(log app.FxLogger) fxevent.Logger {
			return &log
		}),
		fx.Invoke(app.ManageLifeCycle),
	).Run()
}

func asHandler(handler interface{}) fx.Option {
	return fx.Provide(fx.Annotate(
		handler,
		fx.As(new(api.Handler)),
		fx.ResultTags(`group:"handlers"`),
	))
}
