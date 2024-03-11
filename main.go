package main

import (
	"fmt"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ReadAPI struct {
	Location_id      int `json:"location_id"`
	Microcategory_id int `json:"microcategory_id"`
	User_id          int `json:"user_id"`
}

type WriteAPI struct {
	Price            int `json:"price"`
	Location_id      int `json:"location_id"`
	Microcategory_id int `json:"microcategory_id"`
	Matrix_id        int `json:"matrix_id"`
	User_segment_id  int `json:"user_segment_id"`
}

func load_data(data ReadAPI) WriteAPI {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/avito_hack")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	rows, err := db.Query("SELECT price, where_price FROM microcategory WHERE id = ?", data.Location_id)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	new_data := WriteAPI{Price: 10000, Location_id: 1, Microcategory_id: 1, Matrix_id: 2, User_segment_id: 213}
	for rows.Next() {
		err := rows.Scan(&new_data.Price, &new_data.Location_id)
		if err != nil {
			panic(err.Error())
		}
	}

	rows, err := db.Query("SELECT price, where_price FROM location WHERE id = ?", data.Location_id)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	return new_data
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}

		var data ReadAPI
		err = json.Unmarshal(body, &data)
		if err != nil {
			http.Error(w, "Error parsing JSON", http.StatusBadRequest)
			return
		}

		fmt.Printf("Received data with ID: %d, Name: %d\n", data.Location_id, data.Microcategory_id)

		//message := WriteAPI{Price: 10000, Location_id: 1, Microcategory_id: 1, Matrix_id: 2, User_segment_id: 213}

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(load_data(data))
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/api/", handler)
	http.ListenAndServe(":8080", nil)
}
