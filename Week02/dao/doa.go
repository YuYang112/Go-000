package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

var (
	db  *sql.DB
	err error
)

func init() {
	db, err = sql.Open("mysql", "root:123456@tcp(localhost:3306)/mysql?charset=utf8")
}

type Result struct {
	Host     string
	Username string
}

func Dao() (int, error) {
	var info Result
	scanErr := db.QueryRow("select host,username from mysql.servers limit 1").Scan(&info.Host, &info.Username)
	if scanErr != nil {
		return -1, errors.Wrap(scanErr, "Scan error")
	}

	defer db.Close()
	return 1, nil
}
