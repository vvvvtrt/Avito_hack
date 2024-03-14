package main

import (
	"fmt"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"encoding/json"
	//"io/ioutil"
	"log"
	//"net/http"
	"time"
)

var user string = "root"
var password string = "root"
var ip string = "127.0.0.1:3306"
var table_name string = "127.0.0.1:3306"

type Req_data struct {
	Time   int64  `json:"time"`
	Table  string `json:"table"`
	Action string `json:"action"`
	Id     int    `json:"id"`
	Price  int    `json:"price"`
}

func check_update() {
	for true {
		time.Sleep(1 * time.Second)

		formattedString := fmt.Sprintf("%s:%s@tcp(%s)/avito_operation", user, password, ip)

		db, err := sql.Open("mysql", formattedString)
		if err != nil {
			fmt.Println("[ERROR]: ", err.Error())
			continue
		}
		defer db.Close()

		read_operation, err := db.Query("SELECT * FROM operation WHERE time = (SELECT MIN(time) FROM operation) LIMIT 1")
		if err != nil {
			fmt.Println("[ERROR]: ", err.Error())
			continue
		}
		defer read_operation.Close()

		for read_operation.Next() {
			data := Req_data{Time: 0, Table: "", Action: "", Id: 0, Price: 0}
			err = read_operation.Scan(&data.Time, &data.Table, &data.Action, &data.Id, &data.Price)
			if err != nil {
				fmt.Println("Error scanning row:", err)
				return
			}

			query := "SELECT * FROM relevant WHERE id = (SELECT MAX(id) FROM relevant)"
			rows, err := db.Query(query)
			if err != nil {
				panic(err.Error())
			}
			defer rows.Close()

			var id int
			var first_table, secont_table string
			visited := make(map[int]bool)

			for rows.Next() {
				err := rows.Scan(&id, &first_table, &secont_table)
				if err != nil {
					panic(err.Error())
				}

				if data.Action == "add" {
					DFS(id, visited, id, secont_table, data.Table)
				}
			}
		}
	}
}

func DFS(startVertex int, visited map[int]bool, id_update int, table_name string, category string) {
	formattedString := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, ip, table_name)

	db, err := sql.Open("mysql", formattedString)
	if err != nil {
		fmt.Println("[ERROR]: ", err.Error())
		return
	}
	defer db.Close()

	formattedString = fmt.Sprintf("SELECT price, node FROM %s WHERE id = ?", category)
	rows, err := db.Query(formattedString, startVertex)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// new_data := WriteAPI{Price: 0, Location_id: 2, Microcategory_id: 2, Matrix_id: 2, User_segment_id: 213}

	var price int
	var node string
	for rows.Next() {
		err := rows.Scan(&price, &node)
		if err != nil {
			fmt.Println("[ERROR]: ", err.Error())
			return
		}
	}

	visited[startVertex] = true
	fmt.Printf("%d -> ", startVertex)

	var numbers []int
	err = json.Unmarshal([]byte(node), &numbers)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	for _, v := range numbers {
		if !visited[v] {
			//DFS(v, visited)
		}
	}
}

// доработать
func reload_data() {
	db1, err1 := sql.Open("mysql", "пользователь1:пароль1@/avito1")
	if err1 != nil {
		log.Fatal(err1)
	}
	defer db1.Close()

	db2, err2 := sql.Open("mysql", "пользователь2:пароль2@/avito2")
	if err2 != nil {
		log.Fatal(err2)
	}
	defer db2.Close()

	// Выбор данных из таблицы a1 в avito1
	rows, err := db1.Query("SELECT * FROM a1")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Подготовка запроса на вставку данных в таблицу a2 в avito2
	stmt, err := db2.Prepare("INSERT INTO a2 (column1, column2) VALUES (?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// Итерация по результирующему набору и вставка данных в таблицу a2 в avito2
	for rows.Next() {
		var value1, value2 string
		if err := rows.Scan(&value1, &value2); err != nil {
			log.Fatal(err)
		}

		// Вставка данных в таблицу a2 в avito2
		_, err = stmt.Exec(value1, value2)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("Данные успешно скопированы из avito1 в avito2")
}

func main() {
	//go check_update()

	// select {}
}
