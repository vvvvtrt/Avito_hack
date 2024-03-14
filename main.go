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

	rows, err := db.Query("SELECT price, where_price FROM microcategory WHERE id = ?", data.Microcategory_id)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	new_data := WriteAPI{Price: 0, Location_id: 2, Microcategory_id: 2, Matrix_id: 2, User_segment_id: 213}

	var price int
	var where_price int
	for rows.Next() {
		err := rows.Scan(&price, &where_price)
		if err != nil {
			panic(err.Error())
		}
	}

	if where_price != data.Microcategory_id {
		rows3, err := db.Query("SELECT price, where_price FROM microcategory WHERE id = ?", where_price)
		if err != nil {
			panic(err.Error())
		}
		defer rows3.Close()

		for rows3.Next() {
			err := rows3.Scan(&price, &where_price)
			if err != nil {
				panic(err.Error())
			}
		}
	}

	new_data.Microcategory_id = where_price
	new_data.Price = price

	rows2, err := db.Query("SELECT price, where_price FROM location WHERE id = ?", data.Location_id)
	if err != nil {
		panic(err.Error())
	}
	defer rows2.Close()

	for rows2.Next() {
		err := rows2.Scan(&price, &where_price)
		if err != nil {
			panic(err.Error())
		}
	}

	if where_price != data.Microcategory_id {
		rows4, err := db.Query("SELECT price, where_price FROM location WHERE id = ?", where_price)
		if err != nil {
			panic(err.Error())
		}
		defer rows4.Close()

		for rows4.Next() {
			err := rows4.Scan(&price, &where_price)
			if err != nil {
				panic(err.Error())
			}
		}
	}

	new_data.Price = new_data.Price + price
	new_data.Location_id = where_price

	return new_data
}

type Country struct {
	Country_ string `json:"country"`
	ID       string `json:"id"`
}

func country(w http.ResponseWriter, r *http.Request) {
	country_ := Country{Country_: "[Россия]", ID: "[1]"}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(country_)
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
	http.HandleFunc("/country/", country)
	http.ListenAndServe(":8080", nil)
}
