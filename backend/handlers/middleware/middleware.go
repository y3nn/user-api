package middleware

import (
	"auth/backend/model"
	"auth/backend/repository"
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

func CreateItemsMiddleware(repo *repository.ItemsRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			fmt.Println("🚨| Не корретный запрос нужен [POST] - ", r.Method)
			http.Error(w, "🚨| Не корретный запрос нужен [POST]", http.StatusMethodNotAllowed)
			return
		}

		var newItem model.Item
		fmt.Println("newItem: ", newItem)
		if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
			fmt.Println("🚨| Ошибка декодирования: ", err)
			http.Error(w, "🚨| Ошибка декодирования!", http.StatusBadRequest)
			return
		}

		if err := repo.CreateItem(r.Context(), &newItem); err != nil {
			log.Println("🚨| Ошибка записи в БД: ", err)
			http.Error(w, "Ошибка записи в БД", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newItem)
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
