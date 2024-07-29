package http

import (
	"chat_room_mod/core/application"
	"chat_room_mod/core/domain/chatroom"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type ChatroomHandler struct {
	ChatroomUsecase application.ChatroomUsecase
}

func NewChatroomHandler(chatroomUsecase application.ChatroomUsecase) *ChatroomHandler {
	return &ChatroomHandler{ChatroomUsecase: chatroomUsecase}
}

func (h *ChatroomHandler) CreateChatroom(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title  string `json:"title"`
		BaseID string `json:"base_id"`
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Check if both Title and BaseID are provided and not empty
	if input.Title == "" || input.BaseID == "" {
		http.Error(w, "title and base_id are required and cannot be empty", http.StatusBadRequest)
		return
	}

	chatroom, err := h.ChatroomUsecase.CreateChatroom(input.Title, input.BaseID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(chatroom)
}

func (h *ChatroomHandler) SendMessage(w http.ResponseWriter, r *http.Request) {
	var input struct {
		ChatroomID string `json:"chatroom_id"`
		SenderID   string `json:"sender_id"`
		Type       string `json:"type"`
		Data       string `json:"data"`
	}
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	ctime := time.Now()
	utime := ctime
	message, err := h.ChatroomUsecase.SendMessage(input.ChatroomID, input.SenderID, input.Type, input.Data, ctime, utime)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(message)
}

func (h *ChatroomHandler) CheckChatroomId(w http.ResponseWriter, r *http.Request) {
	//id := r.URL.Query().Get("id")
	vars := mux.Vars(r)
	chatroomID := vars["chartroom_id"]
	chatroom, err := h.ChatroomUsecase.CheckChatroomId(chatroomID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(chatroom)
}

func (h *ChatroomHandler) GetChatrooms(w http.ResponseWriter, r *http.Request) {
	chatrooms, err := h.ChatroomUsecase.GetChatrooms()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if chatrooms == nil {
		chatrooms = []*chatroom.Entity_Chatroom{}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(chatrooms)
}

func (h *ChatroomHandler) UpdateChatroom(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	chatroomID := vars["chartroom_id"]

	var input struct {
		Title  string `json:"title"`
		Avatar string `json:"avatar"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Fetch existing chatroom details
	existingChatroom, err := h.ChatroomUsecase.CheckChatroomId(chatroomID)
	if err != nil {
		http.Error(w, "Chatroom not found", http.StatusNotFound)
		return
	}
	existingChatroom.Avatar = input.Avatar
	existingChatroom.Title = input.Title
	// Save updated chatroom
	if err := h.ChatroomUsecase.UpdateChatroom(existingChatroom); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *ChatroomHandler) GetMessages(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	chatroomID := vars["chartroom_id"]
	messages, err := h.ChatroomUsecase.GetMessages(chatroomID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(messages)
}

func (h *ChatroomHandler) LeaveChatroom(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	chatroomID := vars["chatroom_id"]
	userID := r.URL.Query().Get("user_id")
	if chatroomID == "" || userID == "" {
		http.Error(w, "Chatroom ID and User ID are required", http.StatusBadRequest)
		return
	}
	err := h.ChatroomUsecase.RemoveUserFromChatroom(chatroomID, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User left the chatroom successfully"))
}
