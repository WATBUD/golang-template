package design

import (
	. "goa.design/goa/v3/dsl"
	"mai.today/api/design/authentication"
)

var _ = API("api", func() {
	Title("Mai.Today Server")
	Description("Mai.Today Server")
	Version("1.0.0")
	Server("Mai.Today", func() {
		Host("localhost", func() {
			Description("Local hosts.")
			URI("http://localhost:8080")
		})

		Services(
			authentication.ServiceName,
		)
	})
})
