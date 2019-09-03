//go:generate goagen bootstrap -d sebs.software\Malangelus\Clipboard\design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"sebs.software/Malangelus/Clipboard/app"
)

func main() {
	cache := make(map[string]string)

	// Create service
	service := goa.New("Clipboard")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "actions" controller
	c := NewActionsController(service, cache)
	app.MountActionsController(service, c)

	// Start service
	if err := service.ListenAndServe(":8081"); err != nil {
		service.LogError("startup", "err", err)
	}

}
