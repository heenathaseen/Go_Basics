package main

import "fmt"

type dept struct {
	name	 string
	branch string
	year int
	regno int
}

func (a dept) show() {

	fmt.Println(" Name: ", a.name)
	fmt.Println("Branch Name: ", a.branch)
	fmt.Println("Year: ", a.year)
	fmt.Println("Register no.: ", a.regno)
}

func main() {


	b := dept{
		name:	 "heena",
		branch: "CSE",
		year: 2020,
		regno: 340001,
	}

	
	b.show()
}
