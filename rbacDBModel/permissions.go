package rbacDBModel

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

type Permission struct {
	ID   string `json:"id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name"`
}

func createPermission(w http.ResponseWriter, r *http.Request) {
	var permission Permission
	if err := json.NewDecoder(r.Body).Decode(&permission); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := permissionsCollection.InsertOne(context.TODO(), permission)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(permission)
}

func getPermissions(w http.ResponseWriter, r *http.Request) {
	cursor, err := permissionsCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.TODO())

	var permissions []Permission
	if err := cursor.All(context.TODO(), &permissions); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(permissions)
}

func getPermission(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]

	var permission Permission
	err := permissionsCollection.FindOne(context.TODO(), bson.M{"name": name}).Decode(&permission)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(permission)
}

func updatePermission(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]

	var updatedPermission Permission
	if err := json.NewDecoder(r.Body).Decode(&updatedPermission); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	filter := bson.M{"name": name}
	update := bson.M{"$set": updatedPermission}
	_, err := permissionsCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updatedPermission)
}

func deletePermission(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]

	_, err := permissionsCollection.DeleteOne(context.TODO(), bson.M{"name": name})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
