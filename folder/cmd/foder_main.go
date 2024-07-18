package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"folder_mod/infrastructure/db"
	folderHttp "folder_mod/internal/interfaces/http"
	"folder_mod/internal/interfaces/repository"

	"github.com/gorilla/mux"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mongoClient, err := db.NewMongoClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer mongoClient.Disconnect(ctx)
	folderRepo := repository.NewMongoFolderRepository(mongoClient)
	folderHandler := folderHttp.NewFolderHandler(folderRepo)
	router := mux.NewRouter()
	router.HandleFunc("/folders", folderHandler.CreateFolder).Methods("POST")
	router.HandleFunc("/folders", folderHandler.GetFolders).Methods("GET")
	router.HandleFunc("/folders/{id}", folderHandler.DeleteFolder).Methods("DELETE")
	router.HandleFunc("/folders/{id}", folderHandler.UpdateFolderData).Methods("PUT")
	// router.HandleFunc("/folders/{id}/parent", folderHandler.UpdateFolderParentAndChildIDs).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8080", router))
}
