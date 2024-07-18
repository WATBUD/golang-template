package design

import (
	"net/http"

	. "goa.design/goa/v3/dsl"
	cors "goa.design/plugins/v3/cors/dsl"
	"mai.today/api/design/authentication"
	"mai.today/api/design/base"
	"mai.today/api/design/folder"
	"mai.today/api/design/sync"
)

var jwt = JWTSecurity("JWT")

var _ = API("api", func() {
	Title("Mai.Today Server")
	Description("Mai.Today Server")
	Version("1.0.0")
	Security(jwt)
	cors.Origin("*", func() {
		cors.Headers("Authorization", "Content-Type")
		cors.Methods(
			http.MethodGet,
			http.MethodHead,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodConnect,
			http.MethodOptions,
			http.MethodTrace,
		)
	})

	Server("Mai.Today", func() {
		Host("develop", func() {
			Description("Development environment API endpoint.")
			URI("https://dev.api.mai.today/")
		})
		Host("localhost", func() {
			Description("Local hosts.")
			URI("http://localhost:8080")
		})

		Services(
			authentication.ServiceName,
			base.ServiceName,
			sync.ServiceName,
			folder.ServiceName,
		)
	})
})
