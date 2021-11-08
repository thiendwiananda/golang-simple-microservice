package mysql

import (
	"database/sql"
	"time"

	"github.com/thiendwiananda/golang-simple-microservice/domain"
)

type MysqlLogRepository struct {
	DB *sql.DB
}

func NewMysqlLogRepository(db *sql.DB) domain.LogRepository {
	return &MysqlLogRepository{
		DB: db,
	}
}

func (_lr *MysqlLogRepository) Store(log *domain.Log) (err error) {
	// store to DB implementation
	time.Sleep(2 * time.Second)
	return
}
