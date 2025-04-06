package controllers

import (
	"net/http"
	"tripat3k2/url_shortner/config"
	"tripat3k2/url_shortner/models"
	"tripat3k2/url_shortner/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginUserDTO struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

func RegisterUser(c *gin.Context) {
	obj := &RegisterUserRequest{}
	err := c.BindJSON(obj)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(obj.Password), bcrypt.DefaultCost)
	if hashErr != nil {
		c.JSON(http.StatusInternalServerError, hashErr)
		return
	}

	newUser := &models.User{
		Name:     obj.Name,
		Email:    obj.Email,
		Password: string(hashedPassword),
	}

	result := config.DB.Create(newUser)
	if result.Error != nil {
		result.Rollback()
		c.JSON(http.StatusInternalServerError, result.Error.Error())
		return
	}

	result.Commit()
	c.JSON(http.StatusAccepted, gin.H{"message": "User created"})
}

func LoginUser(c *gin.Context) {
	obj := &LoginUserRequest{}
	err := c.BindJSON(obj)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	user := &LoginUserDTO{}
	result := config.DB.Select("email", "name", "password", "id").Where("email = ?", obj.Email).Limit(1).Model(&models.User{}).Find(user)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, result.Error.Error())
		return
	}

	hashErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(obj.Password))
	if hashErr != nil {
		c.JSON(http.StatusBadRequest, hashErr)
		return
	}

	token, tokenErr := utils.CreateToken(user.ID, 30)
	if tokenErr != nil {
		c.JSON(http.StatusInternalServerError, tokenErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successfull", "token": token})
}
