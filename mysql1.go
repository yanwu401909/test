package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
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
	stemQuery, err := db.Prepare("SELECT * FROM user WHERE ID <=15")
	if err != nil {
		log.Fatal(err)
	}
	defer stemQuery.Close()
	for i := 0; i < 50; i++ {
		_, err = stmpInsert.Exec("nazul-"+strconv.Itoa(i), "nazul-"+strconv.Itoa(i), "nazul-"+strconv.Itoa(i)+"@gmail.com")
		if err != nil {
			log.Fatal(err)
		}
	}
	rows, err := stemQuery.Query()
	if err != nil {
		log.Fatal(err)
	}
	var id, name, email, passwd, last_modified []byte
	for rows.Next() {
		err = rows.Scan(&id, &name, &email, &passwd, &last_modified)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%s %s %s %s %s", string(id), string(name), string(email), string(passwd), string(last_modified))
	}

}
