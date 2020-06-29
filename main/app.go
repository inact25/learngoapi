package main

import (
	"log"
	"net/http"

	"cobadulu/services"

	_ "github.com/go-sql-driver/mysql"
)

var db = services.InitDB()
var serv = services.NewService(db)

// var cd = services.GetStudent(db)

func main() {
	log.Println("success connected to the api")
	http.HandleFunc("/students", serv.GetStudent())
	http.HandleFunc("/subjects", serv.GetSubject())
	http.HandleFunc("/teachers", serv.GetTeacher())
	err := http.ListenAndServe("localhost:2020", nil)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
