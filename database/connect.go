package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = ""
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "cuonggoja"
)

//sau khi kết nối trả về  db (db, err := sql.Open("postgres", psqlInfo)) để xử lý
func ConnUser() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	//defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	//fmt.Println("Successfully connected!")
	return db
}
