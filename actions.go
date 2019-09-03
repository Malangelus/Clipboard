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
	// c.View(func(txn *badger.Txn) error {
	// 	record, err := txn.Get([]byte(ctx.GUID.String()))
	// 	if err != nil {
	// 		return err
	// 	}
	// 	if value, err := record.Value(); err != nil {
	// 		return err
	// 	} else {
	// 		fmt.Println(value)
	// 		convertedValue := string(value)
	// 		fmt.Println(convertedValue)
	// 		fmt.Println(&convertedValue)
	// 		res.Value = &convertedValue
	// 	}
	// 	return nil
	// })
	if value := c.Cache[ctx.GUID.String()]; value != "" {
		res.Value = &value
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
	// c.Update(func(txn *badger.Txn) error {
	// 	// Your code here…
	// 	txn.Set([]byte(guid), []byte(ctx.Value))
	// 	return nil
	// })
	c.Cache[guid] = ctx.Value
	res.GUID = &guid
	return ctx.OK(res)
	// ActionsController_Store: end_implement
}
