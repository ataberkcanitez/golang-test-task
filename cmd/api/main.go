package main

import (
 "twitch_chat_analysis/pkg/service"
 "twitch_chat_analysis/pkg/rabbit"

)

func main() {

	rabbit, err := rabbit.NewRabbit()
	if err != nil {
		panic(err)
	}
	defer rabbit.Conn.Close()

	channel, err := rabbit.Conn.Channel()
	if err != nil {
		panic(err)
	}


	service := service.NewService(*rabbit, channel)
	service.Run()
}
