package main

import (
	"log"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "postgres"
)

var (
  role_id int
  role_name string
  )

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
  }
  
	result, err := db.Query("select * from account")
	if err != nil {
		panic(err)
  }
  for result.Next(){
    fmt.Println("Reading from db out")
    err := result.Scan(&role_id,&role_name)
    if err != nil {
    log.Fatal(err)
    }
    fmt.Println(role_id, role_name)
    }
	fmt.Println("Successfully connected!")
}
