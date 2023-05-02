package main

import (
	"database/sql"
	f "fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
    // Open a connection to the database
    db, err := sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/learning")
    if err != nil {

        log.Fatal(err)
    }
    defer db.Close()
	f.Println("Success!")

	err=db.Ping()
	if err!=nil{
		f.Println("error verifying connection with db.ping")
		panic(err.Error())
	}

	insert, err := db.Query("INSERT INTO godata VALUES(4,'shiva');")
    if err !=nil {
        panic(err.Error())
    }
    defer insert.Close()
    f.Println("Yay, values added!")
}



