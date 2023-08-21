package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golangdorset/typing-loudly/database"
	"github.com/golangdorset/typing-loudly/ids"
	"github.com/jmoiron/sqlx"
)

func main() {
	db := sqlx.MustOpen("mysql", "root:password@tcp(127.0.0.1:3306)/golangdorset?parseTime=true&loc=Local")

	router := chi.NewRouter()
	router.Post("/user", createUser(db))
	router.Post("/post", createPost(db))

	http.ListenAndServe(":8080", router)
}

func createUser(db *sqlx.DB) http.HandlerFunc {
	type request struct {
		Name string `json:"name" db:"name"`
		Age  int    `json:"age" db:"age"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var input request
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		id, err := database.CreateUser(db, database.User{
			Name: input.Name,
			Age:  input.Age,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		user, err := database.GetUser(db, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_ = json.NewEncoder(w).Encode(user)
	}
}

func createPost(db *sqlx.DB) http.HandlerFunc {
	type request struct {
		Title  string     `json:"title" db:"title"`
		Body   string     `json:"body" db:"body"`
		UserID ids.UserID `json:"user_id" db:"user_id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {
		var input request
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		id, err := database.CreatePost(db, database.Post{
			UserID: input.UserID,
			Title:  input.Title,
			Body:   input.Body,
		})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		post, err := database.GetPost(db, id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_ = json.NewEncoder(w).Encode(post)
	}
}
