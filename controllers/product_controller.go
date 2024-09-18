package controllers

import (
	"github.com/emoncse/airflow/configs"
	"github.com/emoncse/airflow/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func PaymentController(c *gin.Context) {
	var paymentData []models.Payment

	err := c.ShouldBindJSON(&paymentData)
	if err != nil {
		return
	}
	configs.Database.Create(&paymentData)
	c.JSON(http.StatusOK, paymentData)
}
