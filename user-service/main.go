package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var (
	users = map[string]User{}
	mu    sync.RWMutex
)

func main() {
	http.HandleFunc("/users", usersHandler)
	log.Println("User service running at :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		mu.RLock()
		defer mu.RUnlock()
		var userList []User
		for _, u := range users {
			userList = append(userList, u)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(userList)

	case http.MethodPost:
		var u User
		if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
			http.Error(w, "Invalid user data", http.StatusBadRequest)
			return
		}
		mu.Lock()
		users[u.ID] = u
		mu.Unlock()
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(u)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
