package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"folder_API/infrastructure/db"
	folderHttp "folder_API/internal/interfaces/http"
	"folder_API/internal/interfaces/repository"

	"github.com/gorilla/mux"
)

func main() {
	// Create a context with a timeout of 10 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Initialize MongoDB client
	mongoClient, err := db.NewMongoClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer mongoClient.Disconnect(ctx)

	// Initialize repository and handler
	folderRepo := repository.NewMongoFolderRepository(mongoClient)
	folderHandler := folderHttp.NewFolderHandler(folderRepo)

	// Set up the router
	router := mux.NewRouter()

	// Define folder endpoints
	router.HandleFunc("/folders", folderHandler.CreateFolder).Methods("POST")
	router.HandleFunc("/folders", folderHandler.GetFolders).Methods("GET")
	router.HandleFunc("/folders/{id}", folderHandler.GetFolder).Methods("GET")
	router.HandleFunc("/folders/{id}", folderHandler.UpdateFolder).Methods("PUT")
	router.HandleFunc("/folders/{id}", folderHandler.DeleteFolder).Methods("DELETE")
	router.HandleFunc("/folders/{id}/parent", folderHandler.UpdateFolderParent).Methods("PUT")

	// Start the HTTP server
	log.Fatal(http.ListenAndServe(":8080", router))
}
