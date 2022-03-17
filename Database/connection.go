package Database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/magiconair/properties"
)

var DbConnection *sql.DB

func ConnectDb() *sql.DB {
	p := properties.MustLoadFile("./conf.properties", properties.UTF8).Map()
	var connectionString string = fmt.Sprintf("host=%s user=%s password=%s dbname=%s", p["host"], p["user"], p["password"], p["dbname"])
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")
	DbConnection = db
	return db
}
