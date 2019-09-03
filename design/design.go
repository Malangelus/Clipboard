package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("Clipboard", func() {
	Title("Clipboard API")
	Description("HTTP Clipboard API")
	Host("localhost:8080")
	Scheme("http")
})

var StoreResponse = MediaType("application/vnd.clipboard.store", func() {
	Attributes(func() {
		Attribute("guid")
	})
	View("default", func() {
		Attribute("guid")
	})
})
var RetrieveResponse = MediaType("application/vnd.clipboard.retrieve", func() {
	Attributes(func() {
		Attribute("value")
	})
	View("default", func() {
		Attribute("value")
	})
})

var _ = Resource("actions", func() {
	Action("store", func() {
		Routing(GET("store/:value"))
		Routing(POST("store"))
		Description("Stores a string value in memory")
		Params(func() {
			Param("value", String, "Value to be stored")
		})
		Response(OK, StoreResponse)
	})
	Action("retrieve", func() {
		Routing(GET("retrieve/:guid"))
		Routing(POST("retrieve"))
		Description("Stores a string value in memory")
		Params(func() {
			Param("guid", UUID, "guid corresponding to a stored value")
		})
		Response(OK, RetrieveResponse)
	})
})
