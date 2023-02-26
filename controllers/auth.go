package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"mail/config"
	"mail/models"
	"net/http"
)

type AuthController struct{}

func (controller AuthController) Token(c *gin.Context) {
	cfg := config.GetConfig()

	jsonBytes, errRaw := io.ReadAll(c.Request.Body)
	if errRaw != nil {
		c.String(500, "Error while reading bytes")
		return
	}

	res, errRes := http.Post(
		cfg.GetString("wp.url")+"jwt-auth/v1/token",
		"application/json",
		bytes.NewBuffer(jsonBytes))
	if errRes != nil {
		c.String(500, "Error while request to api")
		return
	}

	if res.StatusCode != 200 {
		c.String(404, "Username or password incorrect")
		return
	}

	var jwt models.JWT
	errDecode := json.NewDecoder(res.Body).Decode(&jwt)
	if errDecode != nil {
		c.String(500, "Error while decoding json")
		return
	}

	c.JSON(200, jwt)
}
