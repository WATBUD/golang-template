package rbacDBModel

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
)

type Role struct {
	ID   string `json:"id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name"`
}

func createRole(w http.ResponseWriter, r *http.Request) {
	var role Role
	if err := json.NewDecoder(r.Body).Decode(&role); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := rolesCollection.InsertOne(context.TODO(), role)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(role)
}

func getRoles(w http.ResponseWriter, r *http.Request) {
	cursor, err := rolesCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.TODO())

	var roles []Role
	if err := cursor.All(context.TODO(), &roles); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(roles)
}

func getRole(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]

	var role Role
	err := rolesCollection.FindOne(context.TODO(), bson.M{"name": name}).Decode(&role)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(role)
}

func updateRole(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]

	var updatedRole Role
	if err := json.NewDecoder(r.Body).Decode(&updatedRole); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	filter := bson.M{"name": name}
	update := bson.M{"$set": updatedRole}
	_, err := rolesCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updatedRole)
}

func deleteRole(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]

	_, err := rolesCollection.DeleteOne(context.TODO(), bson.M{"name": name})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
