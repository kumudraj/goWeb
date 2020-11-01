package storage

import (
	"log"

	"github.com/jmoiron/sqlx"
	config "github.com/kumudraj/goWeb/config"
	_ "github.com/lib/pq"
)

var DB *sqlx.DB

func NewDB(params ...string) *sqlx.DB {
	var err error
	conString := config.GetPostgresConnectionString()

	log.Print(conString)

	DB, err = sqlx.Connect(config.GetDBType(), conString)

	if err != nil {
		log.Panic(err)
	}

	return DB
}

func GetDBInstance() *sqlx.DB {
	return DB
}
