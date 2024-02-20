package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"
	"log"
	"payhere/config"
	"time"
)

var SqlxModule = fx.Module("sqlx", fx.Provide(NewSqlx))

func NewSqlx(cfg *config.Config) (*sqlx.DB, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.Mysql.User, cfg.Mysql.Password, cfg.Mysql.Host, cfg.Mysql.Port, cfg.Mysql.DbName)

	db, err := sqlx.Connect("mysql", dataSourceName)
	if err != nil {
		log.Fatalln(err)
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, err
}
