package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	rolesCollection              *mongo.Collection
	usersCollection              *mongo.Collection
	permissionsCollection        *mongo.Collection
	userRoleAssignmentCollection *mongo.Collection
)

func main() {
	log.Println("Application is starting...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb+srv://williamchuang:hkdlOv3cRPlHUTR1@maitoday.kfzdeqh.mongodb.net/?retryWrites=true&w=majority&appName=MaiToday")

	mongoClient, err := mongo.Connect(ctx, clientOptions)
	// mongoClient, err := database.NewMongoClient()
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
	fmt.Println("Server started successfully on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
