package main

import (
	"github.com/ahmetson/service-lib/configuration"
	"github.com/ahmetson/service-lib/controller"
	"github.com/ahmetson/service-lib/log"
)

// creates a sample destination controller
func newDestination(destinationConfig configuration.Controller, serviceUrl string) *controller.Controller {

	logger, _ := log.New("destination", false)
	destination, _ := controller.SyncReplier(logger)
	destination.AddConfig(&destinationConfig, serviceUrl)

	err := controller.AnyRoute(destination)
	if err != nil {
		logger.Fatal("command.AnyRoute: %w", err)
	}

	return destination
}
