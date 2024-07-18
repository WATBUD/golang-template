package main

import (
	"fmt"
	"folder_mod"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	"mai.today/authentication"
	"mai.today/base"
)

func main() {
	mux := goahttp.NewMuxer()

	goahttp.Servers{
		authentication.NewServer(),
		base.NewServer(),
		folder_mod.NewFolderServer(),
	}.Mount(mux)

	httpsvr := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	fmt.Println("Server is listening on port 8080")
	if err := httpsvr.ListenAndServe(); err != nil {
		panic(err)
	}
}
