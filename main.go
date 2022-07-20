package main

import (
	"github.com/mohammaderm/todoList/app"
	"github.com/mohammaderm/todoList/config"
	account "github.com/mohammaderm/todoList/internal/repository/user"
	"github.com/mohammaderm/todoList/log"
)

func main() {
	config, _ := config.NewConfig("./config/config.yaml")
	logger, err := log.New(&log.Logconfig{
		Path:         config.Logger.Internal_Path,
		Pattern:      config.Logger.Filename_Pattern,
		MaxAge:       config.Logger.Max_Age,
		RotationTime: config.Logger.Rotation_Time,
		RotationSize: config.Logger.Max_Size,
	})
	if err != nil {
		logger.Panic("can not use log pkg.")
	}
	db, err := app.DBconnection(logger, &config.Database)
	if err != nil {
		logger.Panic("can not connect to database.")
	}
	userRepository := account.NewRepository(db, logger)

}
