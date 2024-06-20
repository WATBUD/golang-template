package http

import (
	"context"
	"encoding/json"
	"folder_API/internal/entities"
	"folder_API/internal/usecases"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type FolderHandler struct {
	repo usecases.FolderRepository
}

func NewFolderHandler(repo usecases.FolderRepository) *FolderHandler {
	return &FolderHandler{repo: repo}
}

func (h *FolderHandler) CreateFolder(w http.ResponseWriter, r *http.Request) {
	var folder entities.Folder
	if err := json.NewDecoder(r.Body).Decode(&folder); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.repo.Create(context.Background(), &folder); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Define a response struct without the ID field
	type FolderResponse struct {
		Name        string `json:"name"`
		Color       string `json:"color"`
		Index       int    `json:"index"`
		ParentIndex int    `json:"parentIndex"`
	}

	response := FolderResponse{
		Name:        folder.Name,
		Color:       folder.Color,
		Index:       folder.Index,
		ParentIndex: folder.ParentIndex,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func (h *FolderHandler) GetFolders(w http.ResponseWriter, r *http.Request) {
	folders, err := h.repo.FindAll(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(folders)
}

func (h *FolderHandler) GetFolder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	folder, err := h.repo.FindByID(r.Context(), id.Hex())
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(folder)
}

func (h *FolderHandler) UpdateFolder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var folder entities.Folder
	if err := json.NewDecoder(r.Body).Decode(&folder); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	folder.ID = id

	if err := h.repo.Update(r.Context(), &folder); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(folder)
}

func (h *FolderHandler) DeleteFolder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.repo.Delete(r.Context(), id.Hex()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *FolderHandler) UpdateFolderIndex(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var requestBody struct {
		Index int `json:"index"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.repo.UpdateIndex(r.Context(), id.Hex(), requestBody.Index); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *FolderHandler) UpdateFolderParent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(params["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var requestBody struct {
		ParentIndex int `json:"parentIndex"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.repo.UpdateParent(r.Context(), id.Hex(), requestBody.ParentIndex); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
