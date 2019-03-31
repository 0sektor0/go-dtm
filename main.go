package main

import (
	"net/http"
	"os"

	"github.com/0sektor0/go-dtm/api"

	"github.com/0sektor0/go-dtm/router"
	"github.com/op/go-logging"
)

func CreateLogger() router.ILogger {
	format := logging.MustStringFormatter(`%{color}%{time:15:04:05.000} %{shortfunc} > %{level:.4s} %{id:03x}%{color:reset} %{message}`)
	backend := logging.NewLogBackend(os.Stderr, "", 0)
	formatter := logging.NewBackendFormatter(backend, format)

	log := logging.MustGetLogger("logger")
	logging.SetBackend(formatter)

	return log
}

func main() {
	logger := CreateLogger()
	apiRouter := router.NewRouter(logger)

	nh, err := NewNetworkHandler()
	if err != nil {
		logger.Error(err)
		return
	}

	apiRouter.AddHandlerPost("/signup", nh.AddUser)
	apiRouter.AddHandlerPost("/auth", nh.Authorize)
	apiRouter.AddHandlerPost("/logout", nh.LogOut)

	apiRouter.AddHandlerPost("/tasks/get", nh.AuthMiddleware(nh.GetTasks))
	apiRouter.AddHandlerPost("/task/create", nh.AuthMiddleware(nh.AddTask))
	apiRouter.AddHandlerPost("/task/update", nh.AuthMiddleware(nh.TaskPermisionMiddleware(nh.UpdateTask)))
	apiRouter.AddHandlerPost("/task/get", nh.AuthMiddleware(nh.GetTask))

	settings, _ := api.GetSettings()
	logger.Info(settings)

	err = http.ListenAndServe(settings.Port, apiRouter)
	logger.Critical(err)
}
