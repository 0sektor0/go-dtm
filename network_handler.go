package main

import (
	"net/http"
	"encoding/json"
	"github.com/0sektor0/go-dtm/api"
	"github.com/0sektor0/go-dtm/router"
)

const (
	POST_LOGIN_PARAM = "login"
	POST_PASSWORD_PARAM = "password"
	POST_TOKEN_PARAM = "token"
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

func SendResponse(ctx router.IContext, data interface{}, err error) {
	if err != nil {
		ctx.Logger().Error(err)
		ctx.StatusCode(http.StatusBadRequest)
		ctx.Write([]byte(err.Error()))
		return
	}

	bytes, _ := json.Marshal(data)
	ctx.Write(bytes)
}

// TODO
func (this *NetworkHandler) AuthMiddleware(next router.HandlerFunc) router.HandlerFunc {
	return func(ctx router.IContext) {
		next(ctx)
	}
}

// POST: /auth login=...&password=... => 200 [session json]
func (this *NetworkHandler) Authorize(ctx router.IContext) {
	login := ctx.PostParam(POST_LOGIN_PARAM)
	password := ctx.PostParam(POST_PASSWORD_PARAM)
	
	session, err := this._apiClient.Sessions.LogIn(login, password)
	SendResponse(ctx, session, err)
}

func (this *NetworkHandler) LogOut(ctx router.IContext) {
	token := ctx.PostParam(POST_TOKEN_PARAM)
	err := this._apiClient.Sessions.LogOut(token)
	SendResponse(ctx, "ok", err)
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

// POST /signup login=...&password=... => 200 [user json]
func (this *NetworkHandler) AddUser(ctx router.IContext) {
	login := ctx.PostParam(POST_LOGIN_PARAM)
	password := ctx.PostParam(POST_PASSWORD_PARAM)
	
	user, err := this._apiClient.Users.Create(login, password)
	SendResponse(ctx, user, err)
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