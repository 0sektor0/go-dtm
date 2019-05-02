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
	ERROR_TEXT_PERMISION 	   		= "permissions denied"
	ERROR_TEXT_UNIMPLEMENTED_API 	= "api not implemented"
	CONTEXT_SESSION_KEY        		= "session"
	POST_LOGIN_PARAM           		= "login"
	POST_PASSWORD_PARAM        		= "password"
	POST_TOKEN_PARAM           		= "token"
	POST_TEXT_PARAM            		= "text"
	POST_TITLE_PARAM           		= "title"
	POST_TASK_START_DATE_PARAM 		= "taskStartDate"
	POST_TASK_END_DATE_PARAM   		= "taskEndDate"
	POST_LIMIT_PARAM           		= "limit"
	POST_OFFSET_PARAM          		= "offset"
	POST_TASK_UPDATE_PARAM     		= "taskUpdate"
	POST_COMMENT_TEXT_PARAM	   		= "commentText"
	POST_ID_PARAM	   				= "id"
)

type IPermisionChecker interface {
	CheckPermision(user *models.User, id int) bool
}

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

func SendErrorResponse(ctx router.IContext, errorText string) {
	err := errors.New(errorText)
	SendResponse(ctx, nil, err)
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
			return;
		}

		ctx.AddCtxParam(CONTEXT_SESSION_KEY, session)
		next(ctx)
	}
}

func (this *NetworkHandler) PermisionMiddleware(next router.HandlerFunc, idKey string, checker IPermisionChecker) router.HandlerFunc {
	return func(ctx router.IContext) {
		id, err := ctx.PostParamInt(idKey)
		if err != nil {
			ctx.Logger().Error(errors.New("cant find idkey"))
			SendErrorResponse(ctx, ERROR_TEXT_PERMISION)
			return
		}

		session := GetSessionFromContext(ctx)
		canEditTask := checker.CheckPermision(session.User, id)
		if !canEditTask {
			ctx.Logger().Error(errors.New("cant check permission"))
			SendErrorResponse(ctx, ERROR_TEXT_PERMISION)
			return
		}

		ctx.AddCtxParam(idKey, id)
		next(ctx)
	}
}

func (this *NetworkHandler) TasksPermisionMiddleware(next router.HandlerFunc) router.HandlerFunc {
	return  this.PermisionMiddleware(next, POST_ID_PARAM, this._apiClient.Tasks)
}

func (this *NetworkHandler) CommentsPermisionMiddleware(next router.HandlerFunc) router.HandlerFunc {
	return  this.PermisionMiddleware(next, POST_ID_PARAM, this._apiClient.Comments)
}

func (this *NetworkHandler) DataTypePermisionMiddleware(next router.HandlerFunc) router.HandlerFunc {
	return  this.PermisionMiddleware(next, POST_ID_PARAM, this._apiClient.TaskTypes)
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

	taskType, err := ctx.PostParamInt(POST_ID_PARAM)
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
	id, err := ctx.PostParamInt(POST_ID_PARAM)
	if err != nil {
		SendResponse(ctx, nil, err)
		return
	}

	task, err := this._apiClient.Tasks.FindById(id)
	SendResponse(ctx, task, err)
}

func (this *NetworkHandler) GetTasks(ctx router.IContext) {
	limit, err := ctx.PostParamInt(POST_LIMIT_PARAM)
	if err != nil {
		SendResponse(ctx, nil, err)
		return
	}

	offset, err := ctx.PostParamInt(POST_OFFSET_PARAM)
	if err != nil {
		SendResponse(ctx, nil, err)
		return
	}

	tasks, err := this._apiClient.Tasks.GetList(offset, limit)
	SendResponse(ctx, tasks, err)
}

func (this *NetworkHandler) GetUser(ctx router.IContext) {
	id, err := ctx.PostParamInt(POST_ID_PARAM)
	if err != nil {
		SendResponse(ctx, nil, err)
		return
	}

	login := ctx.PostParam(POST_LOGIN_PARAM)
	if err != nil {
		SendResponse(ctx, nil, err)
		return
	}

	if login == "" {
		user, err := this._apiClient.Users.FindById(id)
		SendResponse(ctx, user, err)
	} else {
		user, err := this._apiClient.Users.FindByLogin(login)
		SendResponse(ctx, user, err)
	}
}

func (this *NetworkHandler) GetUsers(ctx router.IContext) {
	limit, err := ctx.PostParamInt(POST_LIMIT_PARAM)
	if err != nil {
		SendResponse(ctx, nil, err)
		return
	}

	offset, err := ctx.PostParamInt(POST_OFFSET_PARAM)
	if err != nil {
		SendResponse(ctx, nil, err)
		return
	}

	users, err := this._apiClient.Users.GetList(offset, limit)
	SendResponse(ctx, users, err)
}

func (this *NetworkHandler) DeleteTask(ctx router.IContext) {
	value, _ := ctx.CtxParam(POST_ID_PARAM)
	id := value.(int)

	err := this._apiClient.Tasks.Delete(id)
	SendResponse(ctx, nil, err)
}

func (this *NetworkHandler) UpdateTask(ctx router.IContext) {
	value, _ := ctx.CtxParam(POST_ID_PARAM)
	id := value.(int)

	taskBody := []byte(ctx.PostParam(POST_TASK_UPDATE_PARAM))
	task := &models.Task{}

	err := json.Unmarshal(taskBody, task)
	if err != nil {
		SendResponse(ctx, nil, err)
		return
	}

	err = this._apiClient.Tasks.Change(id, task)
	SendResponse(ctx, nil, err)
}

func (this *NetworkHandler) AddComment(ctx router.IContext) {
	session := GetSessionFromContext(ctx)
	
	taskId, err := ctx.PostParamInt(POST_ID_PARAM)
	if err != nil {
		SendResponse(ctx, nil, err)
		return
	}
	
	text := ctx.PostParam(POST_COMMENT_TEXT_PARAM)
	if text == "" {
		SendErrorResponse(ctx, "empty comment")
		return
	}

	err = this._apiClient.Comments.Add(taskId, text, session.User.Id)
	SendResponse(ctx, nil, err)
}

func (this *NetworkHandler) UpdateComment(ctx router.IContext) {
	value, _ := ctx.CtxParam(POST_ID_PARAM)
	id := value.(int)

	text := ctx.PostParam(POST_COMMENT_TEXT_PARAM)
	if text == "" {
		SendErrorResponse(ctx, "empty content")
		return
	}

	err := this._apiClient.Comments.Edit(id, text)
	SendResponse(ctx, nil, err)
}

func (this *NetworkHandler) DeleteComment(ctx router.IContext) {
	value, _ := ctx.CtxParam(POST_ID_PARAM)
	id := value.(int)

	err := this._apiClient.Comments.Delete(id)
	SendResponse(ctx, nil, err)
}

func (this *NetworkHandler) AddAttachment(ctx router.IContext) {
	SendErrorResponse(ctx, ERROR_TEXT_UNIMPLEMENTED_API)
}

func (this *NetworkHandler) DeleteAttachment(ctx router.IContext) {
	SendErrorResponse(ctx, ERROR_TEXT_UNIMPLEMENTED_API)
}

func (this *NetworkHandler) GetTasksStatuses(ctx router.IContext) {
	statuses := this._apiClient.TaskStatuses.GetTypes()
	SendResponse(ctx, statuses, nil)
}

func (this *NetworkHandler) AddTasksStatuses(ctx router.IContext) {
	name := ctx.PostParam(POST_TEXT_PARAM)

	this._apiClient.TaskStatuses.Create(name)
	SendResponse(ctx, nil, nil)
}

func (this *NetworkHandler) DeleteTasksStatuses(ctx router.IContext) {
	value, _ := ctx.CtxParam(POST_ID_PARAM)
	id := value.(int)

	this._apiClient.TaskStatuses.Delete(id)
	SendResponse(ctx, nil, nil)
}

func (this *NetworkHandler) UpdateTaskStatuses(ctx router.IContext) {
	value, _ := ctx.CtxParam(POST_ID_PARAM)
	id := value.(int)
	name := ctx.PostParam(POST_TEXT_PARAM)

	this._apiClient.TaskStatuses.Update(id, name)
	SendResponse(ctx, nil, nil)
}
