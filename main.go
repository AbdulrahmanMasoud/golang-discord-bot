package main

import (
	"fmt"
	conf "github.com/AbdulrahmanMasoud/golang-discord-bot/config"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	c := conf.Config{
		Token:     os.Getenv("TOKEN"),
		BotPrefix: os.Getenv("BOT_PREFIX"),
	}
	fmt.Print(c.GetToke() + " " + c.GetBotPrefix())
}
