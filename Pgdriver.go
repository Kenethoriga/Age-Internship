package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	connStr := "user=your_username password=your_password dbname=your_database sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM public.user_table")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}

	data := []map[string]interface{}{}
	for rows.Next() {
		values := make([]interface{}, len(columns))
		columnPointers := make([]interface{}, len(columns))
		for i := range columns {
			columnPointers[i] = &values[i]
		}
		if err := rows.Scan(columnPointers...); err != nil {
			log.Fatal(err)
		}

		entry := make(map[string]interface{})
		for i, colName := range columns {
			val := values[i]
			b, ok := val.([]byte)
			if ok {
				entry[colName] = string(b)
			} else {
				entry[colName] = val
			}
		}
		data = append(data, entry)
	}

	result := map[string]interface{}{
		"driver": map[string]interface{}{
			"status_code": 200,
			"data":        data,
		},
	}

	jsonResult, err := json.Marshal(result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonResult))
}
