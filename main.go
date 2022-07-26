package main

import (
	"github.com/mohammaderm/todoList/app"
	"github.com/mohammaderm/todoList/config"
	handler "github.com/mohammaderm/todoList/internal/presentation/http"
	jobrepo "github.com/mohammaderm/todoList/internal/repository/job"
	account "github.com/mohammaderm/todoList/internal/repository/user"
	jobservice "github.com/mohammaderm/todoList/internal/service/job"
	"github.com/mohammaderm/todoList/internal/service/user"
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
	redisConn := app.ConnectRedis(logger, &config.Redis)
	if err != nil {
		logger.Panic("can not connect to database.")
	}
	userRepository := account.NewRepository(db, logger)
	jobRepository := jobrepo.NewRepository(db, logger)

	jobCache := jobservice.NewJObCache(logger, redisConn)
	userService := user.NewService(logger, userRepository)
	jobService := jobservice.NewService(logger, jobRepository, jobCache)

	userHandler := handler.NewAuthHanlder(logger, userService)
	jobHandler := handler.NewJobHandller(logger, jobService)

	router := app.RouterProvider(&app.RouteProvider{
		AccountHandler: userHandler,
		JobHandler:     jobHandler,
	})
	server := app.ServerProvider(logger, &config.Server, router)
	server.ListenAndServe()
}
