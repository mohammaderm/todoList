package app

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/mohammaderm/todoList/config"
	"github.com/mohammaderm/todoList/log"
)

func DBconnection(logger log.Logger, config *config.Database) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Postgresql.Host,
		config.Postgresql.Port,
		config.Postgresql.Username,
		config.Postgresql.Password,
		config.Postgresql.Database,
	)
	con, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		logger.Panic("can not connect to database.", map[string]interface{}{
			"error": err.Error(),
		})
		return nil, err
	}
	return con, nil

}
