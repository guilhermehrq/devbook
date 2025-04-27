package controllers

import "net/http"

// CreateUser is a function that creates a new user
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User created successfully"))
}

// GetUsers is a function that gets all users
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Users fetched successfully"))
}

// GetUser is a function that gets a user by id
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User fetched successfully"))
}

// UpdateUser is a function that updates a user by id
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User updated successfully"))
}

// DeleteUser is a function that deletes a user by id
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User deleted successfully"))
}
