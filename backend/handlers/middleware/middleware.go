package middleware

import (
	"auth/backend/model"
	"auth/backend/repository"
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
			http.Error(w, "❌ | Данный метод не поддерживается! - ", http.StatusMethodNotAllowed)
		}

		var newUser model.User
		if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
			http.Error(w, "🚨 | Неверный формат пользователя! - ", http.StatusBadRequest)
			return
		}

		if err := repo.CreateUser(r.Context(), &newUser); err != nil {
			log.Printf("🚨 | Ошибка при создании пользователя: %v", err)
			http.Error(w, " 🚨 | Ошибка сервера при  записи в БД", http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusCreated)
		log.Println("✅ | Пользователь Создан!")
		json.NewEncoder(w).Encode(newUser)
	}
}

func ReadUser(repo *repository.UsersRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			fmt.Println("🚨| Не корретный запрос нужен [GET] - ", r.Method)
			http.Error(w, "Данный метод не поддерживается ", http.StatusMethodNotAllowed)
		}
		getUserID := chi.URLParam(r, "id")
		userID, err := strconv.ParseInt(getUserID, 10, 64)

		if err != nil {
			log.Println("🚨 | strconv error :  ", err)
			return
		}
		var user model.User
		if err := repo.GetUser(r.Context(), userID, &user); err != nil {
			log.Println("🚨 | ОШИБКА: GETUSER - ", err)
			return
		}
		if err = json.NewEncoder(w).Encode(user); err != nil {
			log.Println("🚨 | ОШИБКА: Encode - ", err)
			return
		}
		log.Println("✅  Данные пользователя получены!")
	}
}

func ListMiddleWare(u repository.UsersRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			log.Println("🚫 Uncorrectly request method")
			http.Error(w, "🚨 |  Method now allowed - NEED [GET]", http.StatusMethodNotAllowed)
			return
		}
		users, err := u.ListUsers(context.Background())
		log.Println(users)
		if err != nil {
			log.Println("🚨 | Function error ")
			http.Error(w, "Internal Error ", http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&users)
		log.Println("✅  Список пользователей отправлен!")
	}
}

func DeleteUser(u *repository.UsersRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "DELETE" {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		getUserID := chi.URLParam(r, "id")
		userID, err := strconv.ParseInt(getUserID, 10, 64)
		if err != nil {
			log.Println("🚨 | Invalid user ID: ", err)
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}
		if err := u.DeleteUser(context.Background(), userID); err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully"})
		log.Println("✅ User deleted successfully")
	}
}

func UpdateUser(u repository.UsersRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "PATCH" {
			fmt.Println("🚨| Не корретный запрос нужен [POST] - ", r.Method)
			http.Error(w, "🚨| Не корретный запрос нужен [POST]", http.StatusMethodNotAllowed)
			return
		}

		var updatedUser model.User
		if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
			fmt.Println("🚨| Ошибка декодирования: ", err)
			http.Error(w, "🚨| Ошибка декодирования!", http.StatusBadRequest)
			return
		}

		getUserID := chi.URLParam(r, "id")
		userID, err := strconv.ParseInt(getUserID, 10, 64)
		updatedUser.ID = userID

		if err != nil {
			log.Println("🚨 | User not found")
			http.Error(w, "User not found ", http.StatusBadRequest)
			return
		}
		if err := u.UpdateUser(r.Context(), &updatedUser); err != nil {
			http.Error(w, "Internal Server Error!", http.StatusInternalServerError)
			return
		}
		w.Header().Add("Content-type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "information has beed updated!"})
		log.Println("✅| Пользовательские данные были!")
	}
}
