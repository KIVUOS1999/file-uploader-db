package store

import (
	"database/sql"
	"fmt"

	"github.com/KIVUOS1999/easyLogs/pkg/log"
	"github.com/KIVUOS1999/file-uploader-db/models"

	_ "github.com/lib/pq"
)

type storeStruct struct {
	db *sql.DB
}

func New(cred models.DBCredentials) (IStore, error) {
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable", cred.User, cred.Database, cred.Pass, cred.Host, cred.Port)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Error("Error in opening connection", err.Error())

		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Error("Error ping db", err.Error())

		return nil, err
	}

	log.Info("DB connection succes")

	return &storeStruct{db: db}, nil
}
