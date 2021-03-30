package main

import (
	"fmt"
	"github.com/MathisBurger/AnalAsia/internal/commands"
	"github.com/MathisBurger/commander"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	session, err := discordgo.New("Bot " + os.Getenv("botToken"))
	if err != nil {
		fmt.Println("Error while creating discord session")
		return
	}

	handler := commander.New(";;", "826378911444500501")

	handler.Register("info", "Information command", commands.InfoCommand, 100)

	session.AddHandler(handler.Execute)

	session.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)
	err = session.Open()
	if err != nil {
		fmt.Println("Cannot connect to discord websocket")
		return
	}

	fmt.Println("The bot is running now. Terminate by using CTRL-C")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	session.Close()
}
