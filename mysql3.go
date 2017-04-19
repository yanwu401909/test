package main

import (
	"database/sql"
	"fmt"
	_ "github.com/ziutek/mymysql/godrv"
	"time"
)

const (
	DB_NAME = "test"
	DB_USER = "root"
	DB_PASS = "root"
)

type User struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Passwd       string `json:"passwd"`
	LastModified int64  `json:"last_modified"`
}

func OpenDB() *sql.DB {
	db, err := sql.Open("mymysql", fmt.Sprintf("%s/%s/%s", DB_NAME, DB_USER, DB_PASS))
	if err != nil {
		panic(err)
	}
	return db
}

func UserById(id int64) *User {
	db := OpenDB()
	defer db.Close()
	row := db.QueryRow("SELECT * FROM user WHERE id = ?", id)
	user := new(User)
	row.Scan(&user.Id, &user.Name, &user.Email, &user.Passwd, &user.LastModified)
	return user
}

func FindByName(name string) []User {
	db := OpenDB()
	defer db.Close()
	stmp, err := db.Prepare("SELECT * FROM user WHERE name like ?")
	if err != nil {
		panic(err)
	}
	rows, err := stmp.Query("%" + name + "%")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	result := make([]User, 0)
	for rows.Next() {
		user := new(User)
		rows.Scan(&user.Id, &user.Name, &user.Passwd, &user.Email, &user.LastModified)
		result = append(result, *user)
	}
	return result
}

func DeleteById(id int64) (r int64, err error) {
	db := OpenDB()
	defer db.Close()
	stem, err := db.Prepare("DELETE FROM user WHERE id = ?")
	if err != nil {
		return 0, err
	}
	defer stem.Close()
	result, err := stem.Exec(id)
	if err != nil {
		return 0, err
	}
	r, err = result.RowsAffected()
	return r, err
}

func InsertUser(user *User) (r int64, err error) {
	db := OpenDB()
	defer db.Close()
	stem, err := db.Prepare("INSERT INTO user (name, email, passwd, last_modified) values (?,?,?,?)")
	if err != nil {
		return 0, err
	}
	defer stem.Close()
	result, err := stem.Exec(&user.Name, &user.Email, &user.Passwd, &user.LastModified)
	if err != nil {
		return 0, err
	}
	r, err = result.RowsAffected()
	return r, err
}

func SaveOrUpdateUser(user *User) (r int64, err error) {
	db := OpenDB()
	defer db.Close()
	var stem *sql.Stmt
	var result sql.Result
	if &user.Id != nil && user.Id > 0 {
		stem, err = db.Prepare("UPDATE user SET name = ?, email = ?, passwd = ?, last_modified = ? WHERE id = ?")
	} else {
		stem, err = db.Prepare("INSERT INTO user (name, email, passwd, last_modified) values (?,?,?,?)")
	}
	defer stem.Close()
	if &user.Id != nil && user.Id > 0 {
		result, err = stem.Exec(&user.Name, &user.Email, &user.Passwd, &user.LastModified, &user.Id)
	} else {
		result, err = stem.Exec(&user.Name, &user.Email, &user.Passwd, &user.LastModified)
	}
	if err != nil {
		return 0, err
	}
	r, err = result.RowsAffected()
	return r, err
}
func main() {
	//result := FindByName("nazul")
	//for _, user := range result {
	//	fmt.Print(user.Id, "<=>", user.Email, "\n")
	//}
	//result, err := DeleteById(1813)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(result)
	user := User{Id: 1867, Name: "HhH", Email: "sdfsdfsdf@ccc.com", Passwd: "*****", LastModified: time.Now().Unix()}
	fmt.Println(user)
	result, err := SaveOrUpdateUser(&user)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
