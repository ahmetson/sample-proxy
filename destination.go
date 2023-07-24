package main

import (
	"github.com/ahmetson/service-lib/configuration"
	"github.com/ahmetson/service-lib/controller"
	"github.com/ahmetson/service-lib/log"
)

// creates a sample destination controller
func newDestination(destinationConfig configuration.Controller) controller.Interface {

	logger, _ := log.New("destination", false)
	destination, _ := controller.NewReplier(logger)
	destination.AddConfig(&destinationConfig)

	err := controller.AnyRoute(destination)
	if err != nil {
		logger.Fatal("command.AnyRoute: %w", err)
	}

	return destination
}
