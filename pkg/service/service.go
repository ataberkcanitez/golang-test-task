package service

import (
	"fmt"
	"net/http"
	"twitch_chat_analysis/pkg/api"
	"twitch_chat_analysis/pkg/rabbit"
	"twitch_chat_analysis/pkg/report"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

func NewService(rabbit rabbit.Rabbit, channel *amqp.Channel) *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "OK")
	})

	r.POST("/message", func(c *gin.Context) {
		var requestBody api.RequestBody
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			fmt.Println("olmadi service")
			c.AbortWithStatus(http.StatusBadRequest)
		}

		api.SendMessageToRabbit(c, channel, rabbit, requestBody)
	})

	r.GET("/message/list", func(c *gin.Context) {
		var body report.Body
		if err := c.ShouldBindJSON(&body); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
		}

		api.Getlist(c, body)

	})

	return r
}
