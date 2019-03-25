package main

import (
	"github.com/0sektor0/go-dtm/api"
	"github.com/0sektor0/go-dtm/router"
)

const (
	POST_LOGIN_PARAM = "login"
)

type NetworkHandler struct {
	_apiClient *api.ApiClient
}

func NewNetworkHandler() (*NetworkHandler, error) {
	NetworkHandler := &NetworkHandler{}

	apiClient, err := api.NewApiClient()
	if err != nil {
		return nil, err
	}

	NetworkHandler._apiClient = apiClient
	return NetworkHandler, nil
}

// TODO
func (this *NetworkHandler) AuthMiddleware(next router.HandlerFunc) router.HandlerFunc {
	return func(ctx router.IContext) {
		next(ctx)
	}
}

// POST: /auth login=...&password=... => 200 [session json]
func (this *NetworkHandler) Authorize(ctx router.IContext) {

}

// POST: /task/create title=...&text=...&asignee... => 200 [task json]
func (this *NetworkHandler) AddTask(ctx router.IContext) {

}

// POST: /comment/create text=...&task_id=... => 200 [comment json]
func (this *NetworkHandler) AddComment(ctx router.IContext) {

}

// POST: /attachment/create 
func (this *NetworkHandler) AddAttachment(ctx router.IContext) {

}

// POST /status/create name=... => 200 [taskStatus json]
func (this *NetworkHandler) AddTaskStaus(ctx router.IContext) {

}

// POST /type/create name=... => 200 [taskType json]
func (this *NetworkHandler) AddTaskType(ctx router.IContext) {

}

// POST /signin/create login=...&password=... => 200 [user json]
func (this *NetworkHandler) AddUser(ctx router.IContext) {
	login := ctx.PostParam(POST_LOGIN_PARAM)
	data := []byte(login)
	
	ctx.Write(data)
}

// GET /task/{id} => 200 [task json]
// GET /task/{offset}/{limit}
func (this *NetworkHandler) GetTask(ctx router.IContext) {

}

// GET /user/{id}
func (this *NetworkHandler) GetUser(ctx router.IContext) {

}

func (this *NetworkHandler) GetTaskStatus(ctx router.IContext) {

}

func (this *NetworkHandler) GetTaskType(ctx router.IContext) {

}

func (this *NetworkHandler) EditTask(ctx router.IContext) {

}

func (this *NetworkHandler) EditComment(ctx router.IContext) {

}

func (this *NetworkHandler) EditUser(ctx router.IContext) {

}

func (this *NetworkHandler) EditTaskStaus(ctx router.IContext) {

}

func (this *NetworkHandler) EditTasType(ctx router.IContext) {

}