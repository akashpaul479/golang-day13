package day13

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type students2 struct {
	Id    int
	Name  string
	Email string
	Age   int
}

func Insertdata2() {
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

		fmt.Println("\n__students management database__")
		fmt.Println("1.Insert students")
		fmt.Println("2.List students")
		fmt.Println("3.Delete students")
		fmt.Println("4.Exit ")
		fmt.Println("Enter your choice:")

		choicestr, _ := reader.ReadString('\n')
		choicestr = strings.TrimSpace(choicestr)
		choice, _ := strconv.Atoi(choicestr)
		switch choice {
		case 1:
			insertFlow(db, reader)
			pauseforEnter(reader, "press enter to return menu...")
		case 2:
			ListStudents(db)
			pauseforEnter(reader, "press enter to return menu...")
		case 3:
			DeleteFlow(db, reader)
			pauseforEnter(reader, "press enter to return menu...")
		case 4:
			fmt.Println("Exiting program...")
			return
		default:
			fmt.Println("Invalid choice please try again")

		}
	}
}
func pauseforEnter(reader *bufio.Reader, msg string) {
	fmt.Println()
	fmt.Println(msg)
	_, _ = reader.ReadString('\n')
}

func insertFlow(db *sql.DB, reader *bufio.Reader) {
	var std1 students1

	fmt.Println("Enter name:")
	std1.Name, _ = reader.ReadString('\n')
	std1.Name = strings.TrimSpace(std1.Name)

	fmt.Println("Enter email:")
	std1.Email, _ = reader.ReadString('\n')
	std1.Email = strings.TrimSpace(std1.Email)

	fmt.Print("Enter age:")
	fmt.Scanln(&std1.Age)

	query := `INSERT INTO students(id , name ,email, age) VALUES (? , ? , ? , ?)`
	_, err := db.Exec(query, std1.Id, std1.Name, std1.Email, std1.Age)
	if err != nil {
		log.Printf("Error inserting students : %v\n", err)
	} else {
		fmt.Print("student added succesfuly!")
	}
}
func ListStudents(db *sql.DB) {
	rows, err := db.Query("SELECT id, name, email, age FROM students")
	if err != nil {
		log.Printf("Error fetching students: %v\n", err)
		return
	}
	defer rows.Close()

	fmt.Println("\n--- Student List ---")
	fmt.Printf("%-5s %-15s %-30s %-2s\n", "ID", "Name", "Email", "Age")
	fmt.Println(strings.Repeat("-", 65))
	for rows.Next() {
		var std students1
		if err := rows.Scan(&std.Id, &std.Name, &std.Email, &std.Age); err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		fmt.Printf("%-5d %-15s %-30s %-2d\n", std.Id, std.Name, std.Email, std.Age)
	}
}

func DeleteFlow(db *sql.DB, reader *bufio.Reader) {
	fmt.Print("Enter Student ID to delete: ")
	var id int
	fmt.Scanln(&id)

	query := `DELETE FROM students WHERE id = ?`
	res, err := db.Exec(query, id)
	if err != nil {
		log.Printf("Error deleting student: %v\n", err)
		return
	}

	count, _ := res.RowsAffected()
	if count == 0 {
		fmt.Println("No student found with that ID.")
	} else {
		fmt.Println("Student deleted successfully!")
	}
}
