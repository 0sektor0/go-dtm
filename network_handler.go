package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/0sektor0/go-dtm/api"
	"github.com/0sektor0/go-dtm/models"
	"github.com/0sektor0/go-dtm/router"
)

const (
	CONTEXT_SESSION_KEY        = "session"
	POST_LOGIN_PARAM           = "login"
	POST_PASSWORD_PARAM        = "password"
	POST_TOKEN_PARAM           = "token"
	POST_TEXT_PARAM            = "text"
	POST_TITLE_PARAM           = "title"
	POST_TASK_TYPE_ID_PARAM    = "taskTypeId"
	POST_TASK_STATUS_ID_PARAM  = "taskStatusId"
	POST_TASK_ASIGNEE_ID_PARAM = "taskAsigneeId"
	POST_TASK_START_DATE_PARAM = "taskStartDate"
	POST_TASK_END_DATE_PARAM   = "taskEndDate"
	POST_TASK_ID_PARAM         = "taskId"
	POST_LIMIT_PARAM           = "limit"
	POST_OFFSET_PARAM          = "offset"
	POST_TASK_UPDATE_PARAM     = "taskUpdate"
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

func (this *NetworkHandler) AuthMiddleware(next router.HandlerFunc) router.HandlerFunc {
	return func(ctx router.IContext) {
		token := ctx.PostParam(POST_TOKEN_PARAM)
		session, err := this._apiClient.Sessions.Authentificate(token)

		if err != nil {
			SendResponse(ctx, nil, err)
		}

		ctx.AddCtxParam(CONTEXT_SESSION_KEY, session)
		next(ctx)
	}
}

func (this *NetworkHandler) TaskPermisionMiddleware(next router.HandlerFunc) router.HandlerFunc {
	return func(ctx router.IContext) {
		taskId, err := ctx.PostParamInt(POST_TASK_ID_PARAM)
		if err != nil {
			err = errors.New("premisions denied")
			SendResponse(ctx, nil, err)
		}

		session := GetSessionFromContext(ctx)
		canEditTask := this._apiClient.Tasks.CanUserEditTask(session.User, taskId)
		if !canEditTask {
			err = errors.New("premisions denied")
			SendResponse(ctx, nil, err)
		}

		ctx.AddCtxParam(POST_TASK_ID_PARAM, taskId)
		next(ctx)
	}
}

func (this *NetworkHandler) Authorize(ctx router.IContext) {
	login := ctx.PostParam(POST_LOGIN_PARAM)
	password := ctx.PostParam(POST_PASSWORD_PARAM)

	session, err := this._apiClient.Sessions.LogIn(login, password)
	SendResponse(ctx, session, err)
}

func GetSessionFromContext(ctx router.IContext) *models.Session {
	data, ok := ctx.CtxParam(CONTEXT_SESSION_KEY)
	if !ok {
		return nil
	}

	session := data.(*models.Session)
	return session
}

func (this *NetworkHandler) AddUser(ctx router.IContext) {
	login := ctx.PostParam(POST_LOGIN_PARAM)
	password := ctx.PostParam(POST_PASSWORD_PARAM)

	user, err := this._apiClient.Users.Create(login, password)
	SendResponse(ctx, user, err)
}

func (this *NetworkHandler) LogOut(ctx router.IContext) {
	token := ctx.PostParam(POST_TOKEN_PARAM)
	err := this._apiClient.Sessions.LogOut(token)
	SendResponse(ctx, "ok", err)
}

func (this *NetworkHandler) AddTask(ctx router.IContext) {
	session := GetSessionFromContext(ctx)

	taskType, err := ctx.PostParamInt(POST_TASK_TYPE_ID_PARAM)
	if err != nil {
		taskType = api.DEFAULT_TASK_TYPE
	}

	creatorId := session.User.Id
	title := ctx.PostParam(POST_TITLE_PARAM)
	text := ctx.PostParam(POST_TEXT_PARAM)

	task, err := this._apiClient.Tasks.Create(creatorId, taskType, title, text)
	SendResponse(ctx, task, err)
}

func (this *NetworkHandler) GetTask(ctx router.IContext) {
	id, err := ctx.PostParamInt(POST_TASK_ID_PARAM)
	if err != nil {
		SendResponse(ctx, nil, err)
	}

	task, err := this._apiClient.Tasks.FindById(id)
	SendResponse(ctx, task, err)
}

func (this *NetworkHandler) GetTasks(ctx router.IContext) {
	limit, err := ctx.PostParamInt(POST_LIMIT_PARAM)
	if err != nil {
		SendResponse(ctx, nil, err)
	}

	offset, err := ctx.PostParamInt(POST_OFFSET_PARAM)
	if err != nil {
		SendResponse(ctx, nil, err)
	}

	tasks, err := this._apiClient.Tasks.GetList(offset, limit)
	SendResponse(ctx, tasks, err)
}

func (this *NetworkHandler) DeleteTask(ctx router.IContext) {
	value, _ := ctx.CtxParam(POST_TASK_ID_PARAM)
	id := value.(int)

	err := this._apiClient.Tasks.Delete(id)
	SendResponse(ctx, nil, err)
}

func (this *NetworkHandler) UpdateTask(ctx router.IContext) {
	value, _ := ctx.CtxParam(POST_TASK_ID_PARAM)
	id := value.(int)

	taskBody := []byte(ctx.PostParam(POST_TASK_UPDATE_PARAM))
	task := &models.Task{}

	err := json.Unmarshal(taskBody, task)
	if err != nil {
		SendResponse(ctx, nil, err)
	}

	err = this._apiClient.Tasks.ChangeTask(id, task)
	SendResponse(ctx, nil, err)
}
