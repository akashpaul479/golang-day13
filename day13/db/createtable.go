package day13

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Createtable() {
	dsn := "root:root@tcp(127.0.0.1:3306)/khushi"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening DB %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to DB :%v", err)
	}
	fmt.Println("connected to mysql succesfully!")

	Createtable := `
	CREATE TABLE IF NOT EXISTS students (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		email VARCHAR(100) UNIQUE NOT NULL,
		age INT NOT NULL
	);`
	_, err = db.Exec(Createtable)
	if err != nil {
		log.Fatalf("Error creating table : %v", err)
	}
	fmt.Println("studednts table created succesfully!")
}
