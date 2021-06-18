package main

import (
	"database/sql"
	"fmt"
)
import _ "github.com/go-sql-driver/mysql"

// Customer type
type Customer struct {
	Id int 
	Name string 
}

func main() {

	// inserting a rows
	insert(Customer{101, "heena"})
	insert(Customer{102, "thaseen"})

	// updating the customer by id
	updateById(Customer{101, "Heena Asad"})

	// select all customers
	results := selectAll()

	// iterating a results
	for results.Next() {
		var customer Customer
		results.Scan(&customer.Id, &customer.Name)
		fmt.Println(customer.Id, customer.Name)
	}

	// select customer by id
	result := selectById(101)
	var customer Customer
	result.Scan(&customer.Id, &customer.Name)
	fmt.Println(customer.Id, customer.Name)

	// delete a customer by id
	delete(102)
}

// function to get a database connection
func connect() *sql.DB {
	db, err := sql.Open("mysql", "heena:heena@tcp(45.120.136.152:3306)/heena")
	if err != nil {
		fmt.Println("Error! Getting connection...")
	}
	return db;
}

// function to insert a row in customer table
func insert(customer Customer) {
	db := connect()
	insert, err := db.Query("INSERT INTO customer VALUES (?, ?)", customer.Id, customer.Name)
	if err != nil {
		fmt.Println("Error! Inserting records...")
	}
	defer insert.Close()
	defer db.Close()
}

// function to select all records from customer table
func selectAll() *sql.Rows {
	db := connect()
	results, err := db.Query("SELECT * FROM customer")
	if err != nil {
		fmt.Println("Error! Getting records...")
	}
	defer db.Close()
	return results
}

// function to select a customer record from table by customer id 
func selectById(id int) *sql.Row {
	db := connect()
	result := db.QueryRow("SELECT * FROM customer WHERE id=?", id)
	defer db.Close()
	return result
}

// function to update a customer record by customer id
func updateById(customer Customer) {
	db := connect()
	db.QueryRow("UPDATE customer SET name=? WHERE id=?", customer.Name, customer.Id)
}

// function to delete a customer by customer id
func delete(id int) {
	db := connect()
	db.QueryRow("DELETE FROM customer WHERE id=?", id)
}