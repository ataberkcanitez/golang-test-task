package api

import (
	"fmt"
	"net/http"

	"twitch_chat_analysis/pkg/rabbit"
	rbt "twitch_chat_analysis/pkg/rabbit"
	"twitch_chat_analysis/pkg/redis"
	"twitch_chat_analysis/pkg/report"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
)

type RequestBody struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Message  string `json:"message"`
}

func SendMessageToRabbit(c *gin.Context, channel *amqp.Channel, rabbit rabbit.Rabbit, body RequestBody) {
	var messageBody rbt.Message
	if err := c.ShouldBindJSON(&messageBody); err != nil {
		fmt.Println("olmadi api")
		c.AbortWithStatus(http.StatusBadRequest)
	}

	err := rabbit.Publish(channel, messageBody)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
}

func Getlist(c *gin.Context, body report.Body) {
	rdsConn := redis.NewRedis("localhost:6379")
	result, err := rdsConn.LRange(body.Sender+":::"+body.Receiver, 0, 1).Result()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.JSON(200, gin.H{"data": result})

}
