package User

import (
	"encoding/json"
	"strconv"

	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	// gorm.Model        // for gorm type model struct
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var users []User

// InitializeUsers initializes the users slice with sample data
func InitializeUsers() {
	users = []User{
		{ID: 1, Username: "tejal", Email: "user1@example.com"},
		{ID: 2, Username: "user2", Email: "user2@example.com"},
	}
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	// var users []User
	// DB.Find(&users)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// DeleteUser deletes a user by ID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get URL parameters

	// Convert the ID parameter to an integer
	userID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Find the index of the user with the given ID in the slice
	index := -1
	for i, user := range users {
		if user.ID == userID {
			index = i
			break
		}
	}

	// If the user was not found, return a 404 error
	if index == -1 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "User not found"})
		return
	}

	// Remove the user from the slice using append and slicing
	users = append(users[:index], users[index+1:]...)
	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully"})
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") //content type
	var user User                                      // creating variable user for type User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// users.create(&user)
	users = append(users, user)
	json.NewEncoder(w).Encode(user) // passing data back to browser
}
func GetUsers2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	userID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Find user by ID
	var user User
	for _, u := range users {
		if u.ID == userID {
			user = u
			break
		}
	}
	if user.ID == 0 {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	// Extract user ID from URL parameters
	userID := mux.Vars(r)["id"]
	// Decode the request body to get updated user data
	var updatedUser User
	if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Find the user with the corresponding ID (assuming users is a slice)
	for i, user := range users {
		if strconv.Itoa(user.ID) == userID {
			// Update the user with new data
			users[i] = updatedUser
			// Return the updated user
			json.NewEncoder(w).Encode(updatedUser)
			return
		}
	}
	// If user with given ID was not found, return 404 Not Found
	http.Error(w, "User not found", http.StatusNotFound)
}
