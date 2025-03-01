package storage

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func MustInitDB(dbms string, cs string) *sqlx.DB {
	db, err := sqlx.Connect(dbms, cs)
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}
	return db
}
