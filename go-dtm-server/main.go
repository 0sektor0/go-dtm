package main

import (
	"net/http"
	"os"

	"github.com/0sektor0/go-dtm/go-dtm-server/api"

	"github.com/0sektor0/go-dtm/go-dtm-server/router"
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

	apiRouter.AddHandlerPost("/task/get", nh.AuthMiddleware(nh.GetTask))
	apiRouter.AddHandlerPost("/tasks/get", nh.AuthMiddleware(nh.GetTasks))
	apiRouter.AddHandlerPost("/tasks/create", nh.AuthMiddleware(nh.AddTask))
	apiRouter.AddHandlerPost("/tasks/update", nh.AuthMiddleware(nh.TasksPermisionMiddleware(nh.UpdateTask)))
	apiRouter.AddHandlerPost("/tasks/delete", nh.AuthMiddleware(nh.TasksPermisionMiddleware(nh.DeleteTask)))

	apiRouter.AddHandlerPost("/user/get", nh.AuthMiddleware(nh.GetUser))
	apiRouter.AddHandlerPost("/users/get", nh.AuthMiddleware(nh.GetUsers))
	
	apiRouter.AddHandlerPost("/comments/add", nh.AuthMiddleware(nh.AddComment))
	apiRouter.AddHandlerPost("/comments/update", nh.AuthMiddleware(nh.CommentsPermisionMiddleware(nh.UpdateComment)))
	apiRouter.AddHandlerPost("/comments/delete", nh.AuthMiddleware(nh.CommentsPermisionMiddleware(nh.DeleteComment)))

	apiRouter.AddHandlerPost("/attachments/update", nh.AuthMiddleware(nh.TasksPermisionMiddleware(nh.AddAttachment)))
	apiRouter.AddHandlerPost("/attachments/delete", nh.AuthMiddleware(nh.TasksPermisionMiddleware(nh.DeleteAttachment)))
	
	apiRouter.AddHandlerPost("/statuses/get", nh.AuthMiddleware(nh.GetTasksStatuses))
	apiRouter.AddHandlerPost("/statuses/delete", nh.AuthMiddleware(nh.DeleteTasksStatuses))
	apiRouter.AddHandlerPost("/statuses/update", nh.AuthMiddleware(nh.UpdateTaskStatuses))

	settings, _ := api.GetSettings()
	logger.Info(settings)

	err = http.ListenAndServe(settings.Port, apiRouter)
	logger.Critical(err)
}
