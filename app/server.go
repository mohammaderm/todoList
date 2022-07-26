package app

import (
	"context"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/mohammaderm/todoList/config"
	handler "github.com/mohammaderm/todoList/internal/presentation/http"
	"github.com/mohammaderm/todoList/log"
)

type RouteProvider struct {
	AccountHandler handler.AuthHandlerContract
	JobHandler     handler.JobHandlerContract
}

func ServerProvider(logger log.Logger, config *config.Server, router *mux.Router) *http.Server {
	srv := &http.Server{
		Addr:    config.Host + ":" + config.Port,
		Handler: router,
	}
	_, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	return srv
}

func RouterProvider(rp *RouteProvider) *mux.Router {
	r := mux.NewRouter()
	userRoute := r.PathPrefix("/api/v1/job").Subrouter()
	route := r.PathPrefix("/api/v1").Subrouter()
	userRoute.Use(handler.Auth)
	route.HandleFunc("/auth/login", rp.AccountHandler.Login).Methods("Post")
	route.HandleFunc("/auth/register", rp.AccountHandler.Register).Methods("Post")

	userRoute.HandleFunc("/create", rp.JobHandler.Create).Methods("Post")
	userRoute.HandleFunc("/getAll/", rp.JobHandler.GetAll).Methods("Get")
	userRoute.HandleFunc("/delete/{jobid}", rp.JobHandler.Delete).Methods("Delete")
	userRoute.HandleFunc("/update/{jobid}", rp.JobHandler.Update).Methods("PUT")

	return r
}
