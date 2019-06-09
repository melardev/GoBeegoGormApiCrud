package controllers

import (
	"encoding/json"
	"github.com/melardev/GoBeegoGormApiCrud/dtos"
	"github.com/melardev/GoBeegoGormApiCrud/models"
	"github.com/melardev/GoBeegoGormApiCrud/services"
	"net/http"
	"strconv"
)

type TodosController struct {
	BaseController
}

func (this *TodosController) GetAllTodos() {
	todos := services.FetchTodos()
	this.SendJson(dtos.GetTodoListDto(todos))
}

func (this *TodosController) GetAllPendingTodos() {
	todos := services.FetchPendingTodos()
	this.SendJson(dtos.GetTodoListDto(todos))
}

func (this *TodosController) GetAllCompletedTodos() {
	todos := services.FetchCompletedTodos()
	this.SendJson(dtos.GetTodoListDto(todos))
}

func (this *TodosController) GetTodoById() {
	idStr := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		this.SendJson(dtos.CreateErrorDtoWithMessage(err.Error()))
	}

	todo, err := services.FetchById(uint(id))
	if err != nil {
		this.SendJson(dtos.CreateErrorDtoWithMessage(err.Error()))
	}

	this.SendJson(dtos.GetTodoDetaislDto(&todo))
}

func (this *TodosController) CreateTodo() {
	todoInput := &models.Todo{}

	if err := json.Unmarshal(this.Ctx.Input.RequestBody, todoInput); err != nil {
		// this.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		this.SendJson(dtos.CreateErrorDtoWithMessage(err.Error()))
		return
	}

	todo, err := services.CreateTodo(todoInput.Title, todoInput.Description, todoInput.Completed)
	if err != nil {
		this.Ctx.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		this.SendJson(dtos.CreateErrorDtoWithMessage(err.Error()))
		return
	}

	this.SendJson(dtos.GetTodoDetaislDto(&todo))
}

func (this *TodosController) UpdateTodo() {
	idStr := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		this.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		this.SendJson(dtos.CreateErrorDtoWithMessage("You must set an ID"))
		return
	}

	var todoInput models.Todo
	if err := json.Unmarshal(this.Ctx.Input.RequestBody, &todoInput); err != nil {
		this.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		this.SendJson(dtos.CreateErrorDtoWithMessage(err.Error()))
		return
	}

	todo, err := services.UpdateTodo(uint(id), todoInput.Title, todoInput.Description, todoInput.Completed)
	if err != nil {
		this.SendJson(dtos.CreateErrorDtoWithMessage(err.Error()))
		return
	}

	this.SendJson(dtos.GetTodoDetaislDto(&todo))
}

func (this *TodosController) DeleteTodo() {
	idStr := this.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		this.Ctx.ResponseWriter.WriteHeader(http.StatusBadRequest)
		this.SendJson(dtos.CreateErrorDtoWithMessage("You must set an ID"))
		this.ServeJSON()
		return
	}

	todo, err := services.FetchById(uint(id))

	if err != nil {
		this.Ctx.ResponseWriter.WriteHeader(http.StatusNotFound)
		this.SendJson(dtos.CreateErrorDtoWithMessage("todo not found"))
		return
	}

	err = services.DeleteTodo(&todo)

	if err != nil {
		this.Ctx.ResponseWriter.WriteHeader(http.StatusNotFound)
		this.SendJson(dtos.CreateErrorDtoWithMessage("Could not delete Todo"))
		return
	}

	this.Ctx.ResponseWriter.WriteHeader(http.StatusNoContent)
	this.ServeJSON()
}

func (this *TodosController) DeleteAllTodos() {
	services.DeleteAllTodos()
	this.Ctx.ResponseWriter.WriteHeader(http.StatusNoContent)
	this.ServeJSON()
}
