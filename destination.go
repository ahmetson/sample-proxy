package main

//
// A destination should be another service.
// Here, we define it for example only. So we can run this proxy without relying on another service.
//

import (
	"github.com/ahmetson/service-lib/configuration"
	"github.com/ahmetson/service-lib/controller"
	"github.com/ahmetson/service-lib/log"
)

// creates a sample destination controller
func newDestination(destinationConfig configuration.Controller, serviceUrl string, logger *log.Logger) *controller.Controller {
	destination, _ := controller.SyncReplier(logger)
	destination.AddConfig(&destinationConfig, serviceUrl)

	err := controller.AnyRoute(destination)
	if err != nil {
		logger.Fatal("command.AnyRoute: %w", err)
	}

	return destination
}
