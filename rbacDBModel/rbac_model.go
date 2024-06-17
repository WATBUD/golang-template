package main

import "time"

// Role represents a user role.
type Role struct {
	Name        string   `json:"name" bson:"name"`
	Permissions []string `json:"permissions" bson:"permissions"`
}

// User represents a user in the system.
type User struct {
	ID       string   `json:"id" bson:"id"`
	Username string   `json:"username" bson:"username"`
	Roles    []string `json:"roles" bson:"roles"`
}

// Permission represents a permission in the system.
type Permission struct {
	Name string `json:"name" bson:"name"`
}

// UserRoleAssignment represents the assignment of roles to users.
type UserRoleAssignment struct {
	UserID     string    `json:"user_id" bson:"user_id"`
	RoleID     string    `json:"role_id" bson:"role_id"`
	AssignedAt time.Time `json:"assigned_at" bson:"assigned_at"`
}
