package mysql

import (
	"database/sql"
	_ "embed"
	"log"
	confg "xm/internal/config"

	l "xm/internal/logger"

	_ "github.com/go-sql-driver/mysql"
)

var (
	//go:embed schema.sql
	schemaSQL string
)

//create a new instance of mysql to be used for data operations
func NewMysql(conf confg.Configuration, lg l.Logger) (*sql.DB, error) {
	lg.Info("starting mysql database...")
	log.Println("starting mysql database...")

	_db, err := sql.Open("mysql", conf.MysqlConfig.ConnStr)
	if err != nil {
		lg.Fatal("db startup", err)
		return nil, err
	}

	if err = _db.Ping(); err != nil {
		lg.Fatal("db ping", err)
		return nil, err
	}
	_, err = _db.Exec(schemaSQL)
	if err != nil {
		return nil, err
	}

	lg.Info("database started successfully")
	log.Println("database started successfully")
	return _db, nil
}
