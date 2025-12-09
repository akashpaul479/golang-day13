package project

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type User1 struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func Getuser(db *sql.DB) ([]User1, error) {
	rows, err := db.Query("SELECT id,name ,email, age FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User1
	for rows.Next() {
		var u User1
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Age); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}
func Createuser(db *sql.DB, user *User1) error {
	res, err := db.Exec("INSERT INTO users (name , email , age)VALUES(? , ? , ?)", user.Name, user.Email, user.Age)
	if err != nil {
		return err
	}
	id, _ := res.LastInsertId()
	user.ID = int(id)
	return nil
}
func Updateuser(db *sql.DB, user *User1) error {
	_, err := db.Exec("UPDATE users SET name=?,email=?,age=? WHERE id=?", user.Name, user.Email, user.Age, user.ID)
	return err
}
func Deleteuser(db *sql.DB, id int) error {
	_, err := db.Exec("DELETE FROM users WHERE id=?", id)
	return err
}

func GetUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := Getuser(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(users)

	}
}
func CreateUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user User1
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid input:"+err.Error(), http.StatusBadRequest)
			return
		}
		if err := Createuser(db, &user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(user)
	}
}
func UpdateUserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idstr := vars["id"]
		id, err := strconv.Atoi(idstr)
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}
		var user User1
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid input"+err.Error(), http.StatusBadRequest)
			return
		}
		user.ID = id
		if err := Updateuser(db, &user); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}
func DeleteuserHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		idstr := vars["id"]
		id, err := strconv.Atoi(idstr)
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusBadRequest)
			return
		}
		if err := Deleteuser(db, id); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}
func Restapihandler() {
	dsn := "root:root@tcp(127.0.0.1:3306)/khushi"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS users(
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100)  UNIQUE NOT NULL,
		age INT NOT NULL

	)`)
	if err != nil {
		log.Fatal(err)
	}
	r := mux.NewRouter()
	r.HandleFunc("/users", GetUserHandler(db)).Methods("GET")
	r.HandleFunc("/users", CreateUserHandler(db)).Methods("POST")
	r.HandleFunc("/user{id}", UpdateUserHandler(db)).Methods("PUT")
	r.HandleFunc("/users{id}", DeleteuserHandler(db)).Methods("DELETE")

	fmt.Println("server runnimg on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
