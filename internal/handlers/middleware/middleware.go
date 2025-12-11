package middleware

import (
	"auth/internal/model"
	"auth/internal/repository"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

// middleware func for http-request (create user)
func CreateUserHandler(repo *repository.UsersRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			http.Error(w, "🚫 | This method is not supported!", http.StatusMethodNotAllowed)
			return
		}

		var newUser model.User
		if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
			http.Error(w, "🚨 | Invalid user format!", http.StatusBadRequest)
			return
		}

		if err := repo.CreateUser(r.Context(), &newUser); err != nil {
			log.Printf("🚨 | Error creating user: %v", err)
			http.Error(w, "🚨 | Server error while saving to database", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		log.Println("✅ | User created successfully!")
		json.NewEncoder(w).Encode(newUser)
	}
}

func ReadUser(repo *repository.UsersRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			fmt.Println("🚨| Incorrect request method, need [GET] - ", r.Method)
			http.Error(w, "❌ | This method is not supported!", http.StatusMethodNotAllowed)
			return
		}
		getUserID := chi.URLParam(r, "id")
		userID, err := strconv.ParseInt(getUserID, 10, 64)

		if err != nil {
			log.Println("🚨 | strconv error : ", err)
			http.Error(w, "🚨 | Invalid user ID!", http.StatusBadRequest)
			return
		}
		var user model.User
		if err := repo.GetUser(r.Context(), userID, &user); err != nil {
			log.Println("🚨 | ERROR: GETUSER - ", err)
			http.Error(w, "🚨 | User not found!", http.StatusNotFound)
			return
		}
		if err = json.NewEncoder(w).Encode(user); err != nil {
			log.Println("🚨 | ERROR: Encode - ", err)
			http.Error(w, "🚨 | Error encoding response!", http.StatusInternalServerError)
			return
		}
		log.Println("✅ | User data retrieved successfully!")
	}
}

func ListMiddleWare(u repository.UsersRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			log.Println("🚫 | Incorrect request method")
			http.Error(w, "🚨 | Method not allowed - NEED [GET]", http.StatusMethodNotAllowed)
			return
		}
		users, err := u.ListUsers(context.Background())
		log.Println(users)
		if err != nil {
			log.Println("🚨 | Function error")
			http.Error(w, "🚨 | Internal server error", http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&users)
		log.Println("✅ | Users list sent successfully!")
	}
}

func DeleteUser(u *repository.UsersRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			http.Error(w, "🚫 | Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		getUserID := chi.URLParam(r, "id")
		userID, err := strconv.ParseInt(getUserID, 10, 64)
		if err != nil {
			log.Println("🚨 | Invalid user ID: ", err)
			http.Error(w, "🚨 | Invalid user ID", http.StatusBadRequest)
			return
		}
		if err := u.DeleteUser(context.Background(), userID); err != nil {
			http.Error(w, "🚨 | Internal server error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "✅ | User deleted successfully!"})
		log.Println("✅ | User deleted successfully!")
	}
}

func UpdateUser(u repository.UsersRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PATCH" {
			fmt.Println("🚨| Incorrect request method, need [PATCH] - ", r.Method)
			http.Error(w, "🚨| Incorrect request method, need [PATCH]", http.StatusMethodNotAllowed)
			return
		}

		var updatedUser model.User
		if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
			fmt.Println("🚨| Decoding error: ", err)
			http.Error(w, "🚨| Decoding error!", http.StatusBadRequest)
			return
		}

		getUserID := chi.URLParam(r, "id")
		userID, err := strconv.ParseInt(getUserID, 10, 64)
		updatedUser.ID = userID

		if err != nil {
			log.Println("🚨 | User not found")
			http.Error(w, "🚨 | User not found", http.StatusBadRequest)
			return
		}
		if err := u.UpdateUser(r.Context(), &updatedUser); err != nil {
			http.Error(w, "🚨 | Internal server error!", http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "✅ | User information has been updated!"})
		log.Println("✅ | User information updated!")
	}
}
