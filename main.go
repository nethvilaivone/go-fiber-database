package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	dataname = "mydatabase"
	username = "myuser"
	password = "mypassword"
)


var db *sql.DB

type Product struct {
	ID    int
	Name  string
	Price int
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, username, password, dataname)

	dbs, err := sql.Open("postgres", psqlInfo)
	
	if err != nil {
		log.Fatal(err)
	}
	db = dbs
	defer db.Close()
	defer fmt.Println("finished")
    db.Ping()

	fmt.Println("ping is ok!")

    pro, err := getProducts()
	
	if err != nil {
      log.Fatal(err)
	}

	fmt.Println("get sucessful from database = ",pro )
 
}
