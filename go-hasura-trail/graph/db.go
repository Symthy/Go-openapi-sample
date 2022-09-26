package graph

import (
	"fmt"
	"os"

	"github.com/go-pg/pg"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgrespassword"
	dbname   = "default"
)

func Connect() *pg.DB {
	connStr := os.Getenv("DB_URL")
	opt, err := pg.ParseURL(connStr)
	if err != nil {
		panic(err)
	}
	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	host, port, user, password, dbname)
	// db, err := sql.Open("postgres", psqlInfo)
	db := pg.Connect(opt)
	if _, DBStatus := db.Exec("SELECT 1"); DBStatus != nil {
		fmt.Println(DBStatus)
		panic("PostgreSQL is down")
	}
	return db
}
