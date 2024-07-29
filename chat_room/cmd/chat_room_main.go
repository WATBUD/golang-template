package main

import (
	adapters_http "chat_room_mod/adapters/http"
	"chat_room_mod/adapters/repository"
	"chat_room_mod/core/application"
	mongo_database "chat_room_mod/infrastructure"
	"context"
	"net/http"
)

func main() {
	ctx := context.Background()
	client, err := mongo_database.NewMongoClient(ctx)
	if err != nil {
		panic(err)
	}

	chatroomRepo := repository.NewMongoChatroomRepository(client)
	chatroomService := application.NewChatroomService(chatroomRepo)
	chatroomUsecase := application.NewChatroomUsecase(chatroomService)
	chatroomHandler := adapters_http.NewChatroomHandler(chatroomUsecase)

	router := adapters_http.NewRouter()
	router.HandleFunc("/chatrooms", chatroomHandler.CreateChatroom).Methods("POST")
	router.HandleFunc("/chatroom/{chartroom_id}", chatroomHandler.CheckChatroomId).Methods("GET")
	router.HandleFunc("/chatrooms", chatroomHandler.GetChatrooms).Methods("GET")
	router.HandleFunc("/chatrooms/{chartroom_id}", chatroomHandler.UpdateChatroom).Methods("PUT")

	// router.HandleFunc("/chatroom/{chartroom_id}", chatroomHandler.avatar).Methods("POST")
	// router.HandleFunc("/chatroom/{chartroom_id}", chatroomHandler.removeMember).Methods("POST")
	// router.HandleFunc("/chatroom/{chartroom_id}", chatroomHandler.addMember).Methods("POST")

	router.HandleFunc("/chatroom/{chartroom_id}/messages", chatroomHandler.SendMessage).Methods("POST")
	router.HandleFunc("/chatroom/{chartroom_id}/messages", chatroomHandler.GetMessages).Methods("GET")
	router.HandleFunc("/chatroom/{chatroom_id}/leaveChatroom", chatroomHandler.LeaveChatroom).Methods("POST")

	http.ListenAndServe(":8080", router)
}
