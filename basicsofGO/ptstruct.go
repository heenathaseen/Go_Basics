package main

import "fmt"

type Employee struct {
	firstName, lastName string
	age, salary int
}

func main() {


	emp := &Employee{"heena", "thaseen", 22, 16000}

	
	fmt.Println("First Name:", (*emp).firstName)
	fmt.Println("Last Name:",(*emp).lastName)
	fmt.Println("Age:", (*emp).age)
	fmt.Println("employee_details:",(*emp))
}
