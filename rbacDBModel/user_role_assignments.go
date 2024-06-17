package rbacDBModel

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

type UserRoleAssignment struct {
	UserID string `json:"user_id"`
	RoleID string `json:"role_id"`
}

func assignRoleToUser(w http.ResponseWriter, r *http.Request) {
	var userRoleAssignment UserRoleAssignment
	if err := json.NewDecoder(r.Body).Decode(&userRoleAssignment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := userRoleAssignmentCollection.InsertOne(context.TODO(), userRoleAssignment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(userRoleAssignment)
}

func getUserRoleAssignments(w http.ResponseWriter, r *http.Request) {
	cursor, err := userRoleAssignmentCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.TODO())

	var userRoleAssignments []UserRoleAssignment
	if err := cursor.All(context.TODO(), &userRoleAssignments); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(userRoleAssignments)
}

func getUserRoleAssignment(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["user_id"]
	roleID := params["role_id"]

	var userRoleAssignment UserRoleAssignment
	err := userRoleAssignmentCollection.FindOne(context.TODO(), bson.M{"user_id": userID, "role_id": roleID}).Decode(&userRoleAssignment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(userRoleAssignment)
}

func unassignRoleFromUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID := params["user_id"]
	roleID := params["role_id"]

	_, err := userRoleAssignmentCollection.DeleteOne(context.TODO(), bson.M{"user_id": userID, "role_id": roleID})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
