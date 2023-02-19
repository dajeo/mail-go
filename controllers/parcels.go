package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"mail/config"
	"mail/db"
	"mail/models"
	"net/http"
	"strconv"
)

type ParcelsController struct{}

func (controller ParcelsController) Load(c *gin.Context) {
	session := db.GetDB()

	var cities []models.City
	session.Find(&cities)

	var battalions []models.Battalion
	session.Find(&battalions)

	var specialUnits []models.SpecialUnit
	session.Find(&specialUnits)

	c.JSON(200, gin.H{
		"cities":       cities,
		"battalions":   battalions,
		"specialUnits": specialUnits,
	})
}

func (controller ParcelsController) Get(c *gin.Context) {
	id := c.Param("id")

	var parcel models.Parcel
	session := db.GetDB()
	session.First(&parcel, id)

	if parcel == (models.Parcel{}) {
		c.String(404, "Parcel not found")
		return
	}
	// @todo: check user?
	c.JSON(200, parcel)
}

func (controller ParcelsController) Create(c *gin.Context) {
	var parcelJson models.ParcelRes

	errJson := c.BindJSON(&parcelJson)
	if errJson != nil {
		c.String(400, "JSON not be validated")
		return
	}

	cfg := config.GetConfig()

	res, errFetch := http.Get(cfg.GetString("wp.url") + "wp/v2/users/?search=" + parcelJson.SenderUsername)
	if errFetch != nil {
		c.String(500, "Error while sending request to api")
		return
	}

	var users []models.UserRes
	errDecode := json.NewDecoder(res.Body).Decode(&users)
	if errDecode != nil {
		c.String(500, "Error while decoding json")
		return
	}

	if len(users) <= 0 {
		c.String(404, "User not found")
		return
	}

	senderPhone, errParse := strconv.ParseUint(parcelJson.SenderPhone, 10, 64)
	if errParse != nil {
		c.String(400, "Sender phone cannot be parse to number")
		return
	}

	parcel := models.Parcel{
		UserID:       users[0].ID,
		CityID:       parcelJson.SenderCity,
		Recipient:    parcelJson.RecipientName,
		BattalionID:  parcelJson.Battalion,
		Company:      parcelJson.Company,
		Platoon:      parcelJson.Platoon,
		Department:   parcelJson.Department,
		Sender:       parcelJson.SenderName,
		SenderNumber: senderPhone,
	}

	session := db.GetDB()
	createdParcel := session.Create(&parcel)
	if createdParcel.Error != nil {
		c.String(500, "Error while creating parcel")
		return
	}

	c.JSON(200, gin.H{
		"status": "Parcel created",
		"id":     parcel.ID,
	})
}
