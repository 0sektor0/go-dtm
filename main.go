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

	networkHandler, err := NewNetworkHandler()
	if err != nil {
		logger.Error(err)
		return
	}

	apiRouter.AddHandlerPost("/signup", networkHandler.AddUser)
	apiRouter.AddHandlerPost("/auth", networkHandler.Authorize)
	apiRouter.AddHandlerPost("/logout", networkHandler.LogOut)

	apiRouter.AddHandlerPost("/task/create", networkHandler.AuthMiddleware(networkHandler.AddTask))
	apiRouter.AddHandlerPost("/task/get", networkHandler.AuthMiddleware(networkHandler.GetTask))
	apiRouter.AddHandlerPost("/tasks/get", networkHandler.AuthMiddleware(networkHandler.GetTasks))

	settings, _ := api.GetSettings()
	logger.Info(settings)

	err = http.ListenAndServe(settings.Port, apiRouter)
	logger.Critical(err)
}
