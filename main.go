// main.go
package main

import (
	"fmt"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"encoding/json"
	"net/http"
)

type person struct {
	Name string
	Age  int8
}

// работа с mysql
func openSQL() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/avito")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	insert, err := db.Query("INSERT INTO `microcategory` (`name`, `node`, `price`) VALUES('Телефоны', '2', 0)")

	if err != nil {
		panic(err)
	}

	defer insert.Close()

	fmt.Println("Nice")
}

type Message struct {
	Message string `json:"text"`
}

func getMessageHandler(w http.ResponseWriter, r *http.Request) {
	message := Message{Message: "Hello from Golang backend!"}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	json.NewEncoder(w).Encode(message)
}

func main() {
	// http.HandleFunc("/api/message/", getMessageHandler)

	// http.ListenAndServe(":8000", nil)

	openSQL()
}
