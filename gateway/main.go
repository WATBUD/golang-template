package main

import (
	"fmt"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"mai.today/database"

	goahttp "goa.design/goa/v3/http"
	auth "mai.today/api/gen/authentication"
	authHttpSever "mai.today/api/gen/http/authentication/server"
	"mai.today/authentication"
)

func logRequestMiddleware() func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Received request on path:", r.URL.Path)
			h.ServeHTTP(w, r)
		})
	}
}

func initAuthService(mux goahttp.Muxer, mongo *mongo.Client) *authHttpSever.Server {
	authClient := authentication.InitFirebase("../maitoday-168-dev-fireBase.json")
	authService := authentication.NewAuthentication(authClient, mongo)

	authEndpoints := auth.NewEndpoints(authService)

	dec := goahttp.RequestDecoder
	enc := goahttp.ResponseEncoder
	authServer := authHttpSever.New(authEndpoints, mux, dec, enc, nil, nil)

	return authServer
}

func main() {
	// Connect to MongoDB
	mongoClient, err := database.NewMongoClient()
	if err != nil {
		panic(err)
	}

	// Create HTTP Muxer
	mux := goahttp.NewMuxer()
	authServer := initAuthService(mux, mongoClient)
	authServer.Use(logRequestMiddleware())

	authHttpSever.Mount(mux, authServer)

	httpsvr := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	fmt.Println("Server is listening on port 8080")
	if err := httpsvr.ListenAndServe(); err != nil {
		panic(err)
	}
}
