package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/uuid"
	"sebs.software/Malangelus/Clipboard/app"
)

// ActionsController implements the actions resource.
type ActionsController struct {
	*goa.Controller
	Cache map[string]string
}

// NewActionsController creates a actions controller.
func NewActionsController(service *goa.Service, cache map[string]string) *ActionsController {
	return &ActionsController{Controller: service.NewController("ActionsController"), Cache: cache}
}

// Retrieve runs the retrieve action.
func (c *ActionsController) Retrieve(ctx *app.RetrieveActionsContext) error {
	// ActionsController_Retrieve: start_implement

	// Put your logic here

	res := &app.ClipboardRetrieve{}
	if value := c.Cache[ctx.GUID.String()]; value != "" {
		res.Value = &value
		c.Cache[ctx.GUID.String()] = ""
	}
	return ctx.OK(res)
	// ActionsController_Retrieve: end_implement
}

// Store runs the store action.
func (c *ActionsController) Store(ctx *app.StoreActionsContext) error {
	// ActionsController_Store: start_implement

	// Put your logic here
	res := &app.ClipboardStore{}
	guid := uuid.NewV4().String()
	c.Cache[guid] = ctx.Value
	res.GUID = &guid
	return ctx.OK(res)
	// ActionsController_Store: end_implement
}
