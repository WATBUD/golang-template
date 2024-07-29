package http

import (
	"encoding/json"
	"folder_mod/internal/entities"
	"folder_mod/internal/usecases"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type FolderHandler struct {
	repo usecases.FolderRepository
}

func NewFolderHandler(repo usecases.FolderRepository) *FolderHandler {
	return &FolderHandler{repo: repo}
}

func (h *FolderHandler) CreateFolder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var folder entities.Folder
	err := json.NewDecoder(r.Body).Decode(&folder)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = folder.CheackDefaultValues("create")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.repo.PositionExists(ctx, folder.BaseID, folder.ParentID, *folder.Position)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Create the folder to get its ID
	insertResult, err := h.repo.CreateFolder(ctx, &folder)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(insertResult)
}

func (h *FolderHandler) DeleteFolder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}

	// Recursively delete folder and its children
	err := h.repo.DeleteFolderByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Delete successful"})
}

func (h *FolderHandler) GetFolders(w http.ResponseWriter, r *http.Request) {
	folders, err := h.repo.GetFolders(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(folders)
}

func (h *FolderHandler) UpdateFolderData(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	var folder entities.Folder
	if err := json.NewDecoder(r.Body).Decode(&folder); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Ensure BaseID is provided
	if folder.BaseID == "" {
		http.Error(w, "BaseID is required", http.StatusBadRequest)
		return
	}
	// Set the ID from the URL parameter
	folder.ID = id
	now := time.Now()
	folder.UpdatedAt = now
	if err := h.repo.UpdateFolderData(r.Context(), &folder); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Update successful"})
	//json.NewEncoder(w).Encode(folder)
}
