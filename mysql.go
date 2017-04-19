package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	db, err := sql.Open("mysql", "root:root@/test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	stmpInsert, err := db.Prepare("INSERT INTO user (name, email,passwd) values (?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmpInsert.Close()
	stemQuery, err := db.Prepare("SELECT * FROM user WHERE id <=20")
	if err != nil {
		log.Fatal(err)
	}
	defer stemQuery.Close()
	for i := 0; i < 50; i++ {
		_, err = stmpInsert.Exec("nazul-"+string(i), "nazul-"+string(i), "nazul-"+string(i)+"@gmail.com")
		if err != nil {
			log.Fatal(err)
		}
	}

	type User struct {
		id     int64
		name   string
		email  string
		passwd string
	}

	rows, err := stemQuery.Query()
	if err != nil {
		log.Fatal(err)
	}
	cloumns, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}
	values := make([]sql.RawBytes, len(cloumns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			log.Fatal(err)
		}
		var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			log.Println(cloumns[i], ":", value)
		}
		log.Println("-------------------------")
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	result, err := db.Exec("DELETE FROM user WHERE id >= ?", 15)
	if err != nil {
		log.Fatal(err)
	}
	affectNum, err := result.RowsAffected()
	log.Print("Delete:", affectNum, " records!")
}
