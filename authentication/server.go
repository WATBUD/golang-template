package authentication

import (
	"goa.design/goa/v3/http"
	"mai.today/api/gen/authentication"
	"mai.today/api/gen/http/authentication/server"
)

func NewServer() *server.Server {
	endpoints := authentication.NewEndpoints(Instance())
	dec := http.RequestDecoder
	enc := http.ResponseEncoder
	return server.New(endpoints, http.NewMuxer(), dec, enc, nil, nil)
}
