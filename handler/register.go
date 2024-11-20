package handler

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"mygo/db"
	"mygo/model"
	"net/http"
)

func UserRegister(c *gin.Context) {
	user := &model.User{}
	err := c.ShouldBindJSON(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	user.Password = string(hashedPassword)

	res := db.Db.Create(user)
	if res.Error != nil {
		c.JSON(http.StatusInternalServerError, "database error")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "注册成功"})
}
