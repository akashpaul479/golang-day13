package day13

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type students struct {
	Id    int
	Name  string
	Email string
	Age   int
}

func Insertdata() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/khushi")
	if err != nil {
		log.Fatalf("Error opening DB : %v", err)

	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatalf("Error connecting DB : %s", err)
	}
	fmt.Println("connected to mysql succesfully")

	std := students{
		Id:    1,
		Name:  "Akash",
		Email: "akashpaul479@gmail.com",
		Age:   20,
	}
	if err := insertStudents(db, std); err != nil {
		log.Fatalf("error inserting student : %v", err)
	}
	fmt.Println("Values are added!")
}
func insertStudents(db *sql.DB, std students) error {
	query := `INSERT INTO students(id , name ,email, age) VALUES (? , ? , ? , ?)`
	_, err := db.Exec(query, std.Id, std.Name, std.Email, std.Age)
	return err
}
