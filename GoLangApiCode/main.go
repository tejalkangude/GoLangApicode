// package main

// import (
// 	"fmt"
// 	"strings"
// )

// type deck []string

// func main() {
// 	var name string = "Tejal"
// 	card := newcards()
// 	card2 := []string{newcard2(), name}
// 	// fmt.Println(name)
// 	fmt.Println(card)
// 	card2 = append(card2, "appends new string")
// 	// fmt.Println(card2)

// 	// for index, card2 :=range card2{
// 	// 	fmt.Println(index, card2)
// 	// }
// 	fruit := strings.Split(string(fruits()), ",")
// 	fmt.Println(fruit)
// }

// func fruits() string {
// 	return "apple, Mango, Banana, watermailn"
// }
// func newcards() int {
// 	return 12
// }
// func newcard2() string {
// 	return "Hello"
// }

// // func (d deck) print(){
// // 	for i,card2 := range d{
// // 		fmt.Println(i, card2)
// // 	}
// // }

package main

import (
	"log"
	"net/http"

	"github.com/tejalkangude/GoLangApicode/GoLangApiCode/user"

	"github.com/gorilla/mux"
)

func initializeRouter() {
	r := mux.NewRouter()

	r.HandleFunc("/users", GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", getUsers).Methods("GET")
	r.HandleFunc("/users/create", CreateUser).Methods("POST")
	r.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
	r.HandleFunc("/api/users/{id}", DeleteUser).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func main() {

	users = []User{
		{ID: 1, Username: "tejal", Email: "user1@example.com"},
		{ID: 2, Username: "user2", Email: "user2@example.com"},
	}
	http.Handle("/", http.FileServer(http.Dir("user.go")))
	// InitialMigration()
	initializeRouter()
}

// package main

// import (
// 	"encoding/json"
// 	"log"
// 	"net/http"
// )

// type User struct {
// 	ID       int    `json:"id"`
// 	Username string `json:"username"`
// 	Email    string `json:"email"`
// }

// var users []User

// func getUsers(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(users)
// }

// func createUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var newUser User
// 	err := json.NewDecoder(r.Body).Decode(&newUser)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	users = append(users, newUser)
// 	json.NewEncoder(w).Encode(newUser)
// }

// func main() {
// 	// Initialize the users slice
// 	users = []User{
// 		{ID: 1, Username: "user1", Email: "user1@example.com"},
// 		{ID: 2, Username: "user2", Email: "user2@example.com"},
// 	}

// 	// Define routes
// 	http.HandleFunc("/users", getUsers)
// 	http.HandleFunc("/users/create", createUser)

// 	// Start the HTTP server
// 	log.Fatal(http.ListenAndServe(":8080", nil))
// }
