package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/wakabaseisei/runapp/domain"
	"github.com/wakabaseisei/runapp/usecase"
)

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
	c.JSON(res.StatusCode, NewH("success", user))
}

func (controller *UserController) GetAll(c Context) {

	users, res := controller.Interactor.GetAll()
	if res.Error != nil {
		c.JSON(res.StatusCode, NewH(res.Error.Error(), nil))
		return
	}
	c.JSON(res.StatusCode, NewH("success", users))
}

func (controller *UserController) Post(c Context) {
	var userPost domain.UserPost
	if err := c.ShouldBindJSON(&userPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, res := controller.Interactor.Post(userPost.Sex, userPost.Introduction)

	if res.Error != nil {
		c.JSON(res.StatusCode, NewH(res.Error.Error(), nil))
		return
	}
	c.JSON(res.StatusCode, NewH("user has been created", user))
}

func (controller *UserController) Delete(c Context) {

	id, _ := strconv.Atoi(c.Param("id"))

	user, res := controller.Interactor.Delete(id)
	if res.Error != nil {
		c.JSON(res.StatusCode, NewH(res.Error.Error(), nil))
		return
	}
	c.JSON(res.StatusCode, NewH("user has been deleted", user))
}

func (controller *UserController) Update(c Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var userPost domain.UserPost
	if err := c.ShouldBindJSON(&userPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, res := controller.Interactor.Update(id, userPost.Sex, userPost.Introduction)
	if res.Error != nil {
		c.JSON(res.StatusCode, NewH(res.Error.Error(), nil))
		return
	}
	c.JSON(res.StatusCode, NewH("user has been updated", user))
}
