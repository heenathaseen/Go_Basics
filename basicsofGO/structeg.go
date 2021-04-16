package main

import "fmt"


type Address struct {
	Name string
	city string
	Pincode int
}

func main() {

	var a Address
	fmt.Println(a)

	a1 := Address{"heena", "vellore", 363572}

	fmt.Println("Address1: ", a1)

	a2 := Address{Name: "thaseen", city: "ranipet",
								Pincode: 277001}

	fmt.Println("Address2: ", a2)


	a3 := Address{Name: "chennai"}
	fmt.Println("Address3: ", a3)
}
