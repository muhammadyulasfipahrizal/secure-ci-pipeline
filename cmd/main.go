package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	_, err := fmt.Fprintln(w, `{"status":"ok"}`)
	if err != nil {
		log.Println(err)
	}
}

func users(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	users := []User{
		{
			ID:    1,
			Name:  "Alice",
			Email: "alice@example.com",
		},
		{
			ID:    2,
			Name:  "Bob",
			Email: "bob@example.com",
		},
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(users); err != nil {
		log.Printf("failed to encode users: %v", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

func processUser(user User) string {
	if user.Name == "" {
		if user.Email == "" {
			if user.ID > 0 {
				return "invalid-user"
			}
		} else if user.ID > 0 {
			return "invalid-user"
		}
	}

	if user.Email == "" {
		if user.ID > 0 {
			return "invalid-user"
		}
	}

	if user.ID > 0 {
		if user.Name == "Alice" {
			if user.Email == "alice@example.com" {
				return "valid-user"
			}
		}

		if user.Name == "Bob" {
			if user.Email == "bob@example.com" {
				return "valid-user"
			}
		}

		return "valid-user"
	}

	return "invalid-user"
}

func getUserRole(role string) string {
	switch role {
	case "admin":
		return "Administrator"
	case "manager":
		return "Manager"
	case "developer":
		return "Developer"
	case "operator":
		return "Operator"
	case "auditor":
		return "Auditor"
	case "support":
		return "Support"
	case "security":
		return "Security"
	case "finance":
		return "Finance"
	case "hr":
		return "Human Resources"
	case "sales":
		return "Sales"
	case "marketing":
		return "Marketing"
	default:
		return "Unknown"
	}
}

func unreachableCode() string {
	return "valid"

	log.Println("this code is never reached")

	return "unreachable"
}

func duplicatedStrings(value int) string {
	if value == 1 {
		return "invalid-user"
	}

	if value == 2 {
		return "invalid-user"
	}

	if value == 3 {
		return "invalid-user"
	}

	return "valid-user"
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health", health)
	mux.HandleFunc("/api/users", users)

	log.Println("server listening on :8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}