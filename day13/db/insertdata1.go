package day13

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type students1 struct {
	Id    int
	Name  string
	Email string
	Age   int
}

func Insertdata1() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/khushi")
	if err != nil {
		log.Fatalf("Error opening DB : %v", err)

	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatalf("Error connecting DB : %s", err)
	}
	fmt.Println("connected to mysql succesfully")
	reader := bufio.NewReader(os.Stdin)
	for {
		var std1 students1

		fmt.Println("Enter id:")
		fmt.Scanln(&std1.Id)

		fmt.Println("Enter name:")
		fmt.Scanln(&std1.Name)

		fmt.Println("Enter email:")
		fmt.Scanln(&std1.Email)

		fmt.Println("Enter age:")
		fmt.Scanln(&std1.Age)

		if err := insertStudents1(db, students(std1)); err != nil {
			log.Fatalf("error inserting student : %v", err)
		}
		fmt.Println("Values are added!")

		fmt.Print("Enter 'exit' to quit or press Enter to continue: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)
		if strings.ToLower(choice) == "exit" {
			fmt.Println("Exiting program...")
			break
		}
	}
}
func insertStudents1(db *sql.DB, std students) error {
	query := `INSERT INTO students(id , name ,email, age) VALUES (? , ? , ? , ?)`
	_, err := db.Exec(query, std.Id, std.Name, std.Email, std.Age)
	return err
}
