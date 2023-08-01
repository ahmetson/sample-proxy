package main

import (
	"github.com/ahmetson/service-lib/configuration"
	"github.com/ahmetson/service-lib/log"
	"github.com/ahmetson/service-lib/proxy"
	"github.com/ahmetson/web-controller"
)

// a web-proxy example
func main() {
	logger, err := log.New("sample-proxy", false)
	if err != nil {
		log.Fatal("failed to create a log instance", "error", err)
	}

	appConfig, err := configuration.New(logger)
	if err != nil {
		logger.Fatal("configuration.NewAppConfig", "error", err)
	}

	////////////////////////////////////////////////////////////////////////
	//
	// Initialize the proxy
	//
	////////////////////////////////////////////////////////////////////////

	// the proxy creation will validate the config
	webController, err := web.NewWebController(logger)
	if err != nil {
		logger.Fatal("failed to create a webController controller", "error", err)
	}

	service := proxy.New(appConfig, logger)
	service.SetCustomSource(webController)

	if err != nil {
		logger.Fatal("failed to add source controller to the proxy", "error", err)
	}
	service.Controller.RequireDestination(configuration.ReplierType)

	err = service.Prepare()
	if err != nil {
		logger.Fatal("failed to prepare the service", "error", err)
	} else {
		logger.Info("server was prepared", "url", service.Config.Service.Url)
	}

	destinationLogger := logger.Child("temp-destination")
	destinationConfig, _ := service.Config.Service.GetController(configuration.DestinationName)
	destination := newDestination(destinationConfig, service.Config.Service.Url, destinationLogger)

	go service.Run()
	err = destination.Run()
	if err != nil {
		destinationLogger.Fatal("controller.Run", "error", err)
	}
}
