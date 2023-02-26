package controllers

import (
	"github.com/gin-gonic/gin"
	"mail/db"
	"mail/models"
	"strconv"
	"time"
)

type StatusesController struct{}

func (controller StatusesController) Add(c *gin.Context) {
	idStr := c.Param("id")
	id, errParse := strconv.ParseUint(idStr, 10, 64)

	if errParse != nil {
		c.String(404, "Cannot parse id")
	}

	var statusJson models.StatusRes

	if err := c.BindJSON(&statusJson); err != nil {
		c.String(400, "Cannot parse json")
		return
	}

	status := models.Status{
		ParcelID:  id,
		MessageID: statusJson.Message,
		CityID:    statusJson.City,
		Timestamp: time.Now(),
	}

	session := db.GetDB()

	if err := session.Create(&status).Error; err != nil {
		c.String(500, "Error while creating status")
		return
	}

	c.String(200, "Status created")
}
