package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"rbac"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var rolesCollection *mongo.Collection
var usersCollection *mongo.Collection
var permissionsCollection *mongo.Collection
var userRoleAssignmentCollection *mongo.Collection

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoClient, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer mongoClient.Disconnect(ctx)

	rolesCollection = mongoClient.Database("rbac_db").Collection("roles")
	usersCollection = mongoClient.Database("rbac_db").Collection("users")
	permissionsCollection = mongoClient.Database("rbac_db").Collection("permissions")
	userRoleAssignmentCollection = mongoClient.Database("rbac_db").Collection("user_role_assignments")

	router := mux.NewRouter()

	// Role endpoints
	router.HandleFunc("/roles", createRole).Methods("POST")
	router.HandleFunc("/roles", getRoles).Methods("GET")
	router.HandleFunc("/roles/{name}", getRole).Methods("GET")
	router.HandleFunc("/roles/{name}", updateRole).Methods("PUT")
	router.HandleFunc("/roles/{name}", deleteRole).Methods("DELETE")

	// User endpoints
	router.HandleFunc("/users", createUser).Methods("POST")
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users/{id}", getUser).Methods("GET")
	router.HandleFunc("/users/{id}", updateUser).Methods("PUT")
	router.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

	// Permission endpoints
	router.HandleFunc("/permissions", createPermission).Methods("POST")
	router.HandleFunc("/permissions", getPermissions).Methods("GET")
	router.HandleFunc("/permissions/{name}", getPermission).Methods("GET")
	router.HandleFunc("/permissions/{name}", updatePermission).Methods("PUT")
	router.HandleFunc("/permissions/{name}", deletePermission).Methods("DELETE")

	// UserRoleAssignment endpoints
	router.HandleFunc("/assign-role", assignRoleToUser).Methods("POST")
	router.HandleFunc("/user-role-assignments", getUserRoleAssignments).Methods("GET")
	router.HandleFunc("/user-role-assignments/{user_id}/{role_id}", getUserRoleAssignment).Methods("GET")
	router.HandleFunc("/user-role-assignments/{user_id}/{role_id}", unassignRoleFromUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}

// Create Role
func createRole(w http.ResponseWriter, r *http.Request) {
	var role rbac.Role
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

// Get All Roles
func getRoles(w http.ResponseWriter, r *http.Request) {
	cursor, err := rolesCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.TODO())

	var roles []rbac.Role
	if err := cursor.All(context.TODO(), &roles); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(roles)
}

// Get Role
func getRole(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]

	var role rbac.Role
	err := rolesCollection.FindOne(context.TODO(), bson.M{"name": name}).Decode(&role)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(role)
}

// Update Role
func updateRole(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	name := params["name"]

	var updatedRole rbac.Role
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

// Delete Role
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

// Create User
func createUser(w http.ResponseWriter, r *http.Request) {
	var user rbac.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err := usersCollection.InsertOne(context.TODO(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// Get All Users
func getUsers(w http.ResponseWriter, r *http.Request) {
	cursor, err := usersCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(context.TODO())

	var users []rbac.User
	if err := cursor.All(context.TODO(), &users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(users)
}

// Get User
func getUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var user rbac.User
	err := usersCollection.FindOne(context.TODO(), bson.M{"id": id}).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

// Update User
func updateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var updatedUser rbac.User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	filter := bson.M{"id": id}
	update := bson.M{"$set": updatedUser}
	_, err := usersCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(updatedUser)
}

// Delete User
func deleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	_, err := usersCollection.DeleteOne(context.TODO(), bson.M{"id": id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Create Permission
func createPermission(w http.ResponseWriter, r *http.Request) {
	var permission rbac.Permission
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
