package controllers

import (
	"fmt"
	"github.com/emoncse/airflow/configs"
	"github.com/emoncse/airflow/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetUsers(c *gin.Context) {
	var user []models.UserData
	configs.Database.Table("users").Find(&user)
	fmt.Print(user)
	c.JSON(http.StatusOK, user)
}
