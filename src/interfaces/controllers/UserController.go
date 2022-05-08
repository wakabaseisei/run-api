package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wakabaseisei/runapp/usecase"
)

// ハンドラー
type UserHandler interface {
	Get(c Context)
	GetAll(c Context)
	Post(c Context)
	Delete(c Context)
	Update(c Context)
}

type UserController struct {
	Interactor usecase.UserUsecase
}

func NewUserController(usecase usecase.UserUsecase) UserHandler {
	return &UserController{Interactor: usecase}
}

func (controller *UserController) Get(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	user, res := controller.Interactor.Get(id)
	if res.Error != nil {
		c.JSON(res.StatusCode, NewH(res.Error.Error(), nil))
		return
	}
	userResponse := ConvertToUserResponse(user)
	c.JSON(res.StatusCode, NewH("success", userResponse))
}

func (controller *UserController) GetAll(c Context) {
	users, res := controller.Interactor.GetAll()
	if res.Error != nil {
		c.JSON(res.StatusCode, NewH(res.Error.Error(), nil))
		return
	}

	usersResponse := []UserResponse{}
	for i := range users {
		usersResponse = append(usersResponse, ConvertToUserResponse(users[i]))
	}

	c.JSON(res.StatusCode, NewH("success", usersResponse))
}

func (controller *UserController) Post(c Context) {
	var jsonRequest UserJsonRequest
	if err := c.ShouldBindJSON(&jsonRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, res := controller.Interactor.Post(jsonRequest.Sex, jsonRequest.Introduction)

	if res.Error != nil {
		c.JSON(res.StatusCode, NewH(res.Error.Error(), nil))
		return
	}

	userResponse := ConvertToUserResponse(user)
	c.JSON(res.StatusCode, NewH("user has been created", userResponse))
}

func (controller *UserController) Delete(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	user, res := controller.Interactor.Delete(id)
	if res.Error != nil {
		c.JSON(res.StatusCode, NewH(res.Error.Error(), nil))
		return
	}

	userResponse := ConvertToUserResponse(user)
	c.JSON(res.StatusCode, NewH("user has been deleted", userResponse))
}

func (controller *UserController) Update(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var jsonRequest UserJsonRequest
	if err := c.ShouldBindJSON(&jsonRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, res := controller.Interactor.Update(id, jsonRequest.Sex, jsonRequest.Introduction)
	if res.Error != nil {
		c.JSON(res.StatusCode, NewH(res.Error.Error(), nil))
		return
	}

	userResponse := ConvertToUserResponse(user)
	c.JSON(res.StatusCode, NewH("user has been updated", userResponse))
}
