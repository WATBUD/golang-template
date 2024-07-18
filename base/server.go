package base

import (
	"goa.design/goa/v3/http"
	"mai.today/api/gen/base"
	"mai.today/api/gen/http/base/server"
)

func NewServer() *server.Server {
	endpoints := base.NewEndpoints(Instance())
	dec := http.RequestDecoder
	enc := http.ResponseEncoder
	svr := server.New(endpoints, http.NewMuxer(), dec, enc, nil, nil)

	return svr
}
