package main

import (
	"bufio"
    "database/sql"
    "fmt"
    "log"
	"os"
    _ "github.com/go-sql-driver/mysql"
	//"strings"
)
type Student struct {
	regno int
	name  string
	dept  string
}
func main(){
	db, err := sql.Open("mysql", "heena:heena@tcp(45.120.136.152:3306)/heena?parseTime=true")
	if err != nil {
        log.Fatal(err)
    }
    if err := db.Ping(); err != nil {
        log.Fatal(err)
    }
	var s Student


	var choose int

	fmt.Println("----------------------------")
	fmt.Println("******STUDENT DETAILS******")
	fmt.Println("----------------------------")
	fmt.Println("1   create user")
	fmt.Println("2   read")
	fmt.Println("3   oneview")
	fmt.Println("4   delete")
	fmt.Println("Enter your choice")
    fmt.Scanf("%d",&choose)

	switch choose{
	case 1:
		fmt.Println("Enter the Student details")
		addstudent()
		fmt.Println("Successfully inserted")


	case 2:
		fmt.Println("Student Details")
		view()
	case 3:
		fmt.Println("Student detail")
		fmt.Println("Enter the Regno.")
		
			stdin := bufio.NewReader(os.Stdin)
			var userI string

			for {

			_, err := fmt.Fscan(stdin, &userI)

			if err == nil {

			break

			}

			stdin.ReadString('\n')

			}
				
		
				query := "SELECT regno,name, dept FROM studentdetails WHERE regno= ?"
				if err := db.QueryRow(query, userI).Scan(&s.regno, &s.name, &s.dept); err != nil {
					log.Fatal(err)
				}
		
				fmt.Println(s.regno, s.name, s.dept)
			
	case 4:
		fmt.Println("Enter the Registration number you want to delete")
        {
			stdin := bufio.NewReader(os.Stdin)
			var userI string

			for {

			_, err := fmt.Fscan(stdin, &userI)

			if err == nil {

			break

			}

			stdin.ReadString('\n')

			}


			_, err := db.Exec(`DELETE FROM studentdetails  WHERE regno = ?`,userI)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

}



func view(){
	{ db, err := sql.Open("mysql", "heena:heena@tcp(45.120.136.152:3306)/heena?parseTime=true")
	if err != nil {
        log.Fatal(err)
    }
    if err := db.Ping(); err != nil {
        log.Fatal(err)
    }
		rows, err := db.Query(`SELECT * FROM studentdetails`)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		var students []Student
		for rows.Next() {
			var s Student

			err := rows.Scan(&s.regno, &s.name, &s.dept)
			if err != nil {
				log.Fatal(err)
			}
			students = append(students, s)
		}
		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}

		fmt.Println(students)
		}
}
func addstudent() {
	

	fmt.Println("Enter student Regno: ")
	a:= inputFromUserint()
	fmt.Println("Enter student Name: ")
	b := inputFromUser()
	fmt.Println("Enter student Department: ")
	c:= inputFromUser()

	
	db, err := sql.Open("mysql", "heena:heena@tcp(45.120.136.152:3306)/heena?parseTime=true")
	if err != nil {
        log.Fatal(err)
    }
    if err := db.Ping(); err != nil {
        log.Fatal(err)
    }
	var s Student
	defer db.Close()

	s.regno = a
	s.name = b
	s.dept= c
	
	fmt.Println(s.regno, s.name, s.dept)
	_, err = db.Exec(`INSERT INTO studentdetails (regno, name, dept) VALUES (?, ?, ?)`, s.regno, s.name, s.dept)
}

func inputFromUser() string {
	stdin := bufio.NewReader(os.Stdin)
	var a string
	for {

		_, err := fmt.Fscan(stdin, &a)
		if err == nil {
			break
		}
		stdin.ReadString('\n')
	}
	return a
}
func inputFromUserint() int{
	stdin := bufio.NewReader(os.Stdin)
	var a int
	for {
		_, err := fmt.Fscan(stdin, &a)
		if err == nil {
			break
		}
		stdin.ReadString('\n')
	}
	return a
}


