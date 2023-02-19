package controllers

import (
	"github.com/gin-gonic/gin"
	"mail/db"
	"mail/models"
	"strconv"
	"time"
)

type ItemsController struct{}

func (controller ItemsController) GetItemsByParcelID(c *gin.Context) {
	id := c.Param("id")

	var items []models.Item
	session := db.GetDB()
	if err := session.Where("parcel_id = ?", id).Find(&items).Error; err != nil {
		c.String(500, "Error: "+err.Error())
		return
	}

	if len(items) == 0 {
		c.String(404, "Items not found")
		return
	}

	c.JSON(200, items)
}

func (controller ItemsController) Create(c *gin.Context) {
	parcelIdStr := c.Param("id")
	parcelId, errParse := strconv.ParseUint(parcelIdStr, 10, 64)
	if errParse != nil {
		c.String(400, "Cannot parse id to int")
		return
	}

	item := models.Item{
		ParcelID:  parcelId,
		Name:      "",
		Unit:      "шт.",
		Quantity:  1,
		Timestamp: time.Now(),
	}

	session := db.GetDB()
	createdItem := session.Create(&item)
	if createdItem.Error != nil {
		c.String(500, "Error while creating item")
		return
	}

	c.JSON(200, item)
}

func (controller ItemsController) Update(c *gin.Context) {
	var itemJson models.ItemUpdate
	if err := c.BindJSON(&itemJson); err != nil {
		c.String(400, "JSON cannot be parsed")
		return
	}

	id := c.Param("id")

	var item models.Item

	session := db.GetDB()
	if err := session.First(&item, id).Error; err != nil {
		c.String(404, "Item not found")
		return
	}

	item.Name = itemJson.Name
	item.Unit = itemJson.Unit
	item.Quantity = itemJson.Quantity

	if err := session.Save(&item).Error; err != nil {
		c.String(500, "Item cannot be updated")
		return
	}

	c.String(200, "Item updated")
}
