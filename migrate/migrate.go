package main

import (
	"fmt"
	"go-sample/db"
	"go-sample/model"
	"log"
)

func main() {
	dbConn := db.NewDB()
	defer fmt.Println("Success")
	defer db.CloseDB(dbConn)
	err := dbConn.AutoMigrate(&model.User{}, &model.Task{})
	if err != nil {
		log.Fatalln(err)
	}
}
