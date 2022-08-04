package main

import (
	"fmt"

	"github.com/rekar-accounting/backend/cmd/database"
)

func main() {
	fmt.Println("Start Service General ............")
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err)

	} else {
		fmt.Println("Connected")

		postdb, _ := db.DB()
		defer postdb.Close()
	}
}
