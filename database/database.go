package database

import (
	"github.com/golangdorset/typing-loudly/ids"
	"github.com/jmoiron/sqlx"
)

type User struct {
	Name string     `db:"name" json:"name"`
	Age  int        `db:"age" json:"age"`
	ID   ids.UserID `db:"id" json:"id"`
}

func CreateUser(db *sqlx.DB, input User) (ids.UserID, error) {
	id := ids.MakeULID[ids.UserID]()

	if _, err := db.Exec(
		"INSERT INTO users (id, name, age) VALUES (?, ?, ?)",
		id, input.Name, input.Age,
	); err != nil {
		return ids.UserID{}, err
	}
	return id, nil
}

func GetUser(db *sqlx.DB, id ids.UserID) (User, error) {
	var user User
	if err := db.Get(&user, "SELECT id, name, age FROM users WHERE id = ?", id); err != nil {
		return User{}, err
	}
	return user, nil
}

type Post struct {
	Title  string     `db:"title" json:"title"`
	Body   string     `db:"body" json:"body"`
	ID     ids.PostID `db:"id" json:"id"`
	UserID ids.UserID `db:"user_id" json:"user_id"`
}

func CreatePost(db *sqlx.DB, input Post) (ids.PostID, error) {
	id := ids.MakeULID[ids.PostID]()

	if _, err := db.Exec(
		"INSERT INTO posts (id, user_id, title, body) VALUES (?, ?, ?, ?)",
		id, input.UserID, input.Title, input.Body,
	); err != nil {
		return ids.PostID{}, err
	}
	return id, nil
}

func GetPost(db *sqlx.DB, id ids.PostID) (Post, error) {
	var post Post
	if err := db.Get(&post, "SELECT id, user_id, title, body FROM posts WHERE id = ?", id); err != nil {
		return Post{}, err
	}
	return post, nil
}
