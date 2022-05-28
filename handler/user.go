package handler

import (
	"bwastartup/helper"
	"bwastartup/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userservice user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegisterInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatErrorValidation(err)
		error_message := gin.H{"error": errors}
		response := helper.FormatResponse("register failed", http.StatusUnprocessableEntity, "error", error_message)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	new_user, err := h.userservice.RegisterUser(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	format_user := user.FormatUser(new_user, "jwttoken")
	response := helper.FormatResponse("account has been registered", http.StatusOK, "success", format_user)
	c.JSON(http.StatusOK, response)

}

func (h *userHandler) Login(c *gin.Context) {
	var input user.LoginInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatErrorValidation(err)
		error_message := gin.H{"error": errors}
		response := helper.FormatResponse("login failed", http.StatusUnprocessableEntity, "error", error_message)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	loginUser, err := h.userservice.LoginUser(input)
	if err != nil {
		error_message := gin.H{"error": err.Error()}
		response := helper.FormatResponse("login failed", http.StatusUnprocessableEntity, "error", error_message)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	format_user := user.FormatUser(loginUser, "jwttoken")
	response := helper.FormatResponse("login succes", http.StatusOK, "succes", format_user)
	c.JSON(http.StatusOK, response)
}
