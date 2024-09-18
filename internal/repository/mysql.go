package repository

import (
	"fmt"
	"github.com/fanfaronDo/portfolio_v/internal/config"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func NewMySql(cfg config.Mysql) (*sql.DB, error) {
	conn, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfg.User,
		cfg.Password,
		cfg.Address,
		cfg.Port,
		cfg.Database))

	if err != nil {
		return nil, err
	}

	if err := conn.Ping(); err != nil {
		return nil, err
	}

	return conn, nil
}
