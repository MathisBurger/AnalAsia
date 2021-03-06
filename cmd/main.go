package main

import (
	"fmt"
	"github.com/MathisBurger/AnalAsia/internal/collector"
	"github.com/MathisBurger/AnalAsia/internal/commands"
	"github.com/MathisBurger/AnalAsia/internal/database"
	"github.com/MathisBurger/commander"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// Load environment variables from .env file
	// if bot is in development mode
	// This makes only sense for canary builds and tests
	if os.Getenv("mode") == "dev" {
		err := godotenv.Load()
		if err != nil {
			panic(err.Error())
		}
	}

	// Initialize the database
	database.InitDatabase()

	fmt.Println("-------------------------------------------------------------------------------------")
	fmt.Println("Welcome to AnalAsia")
	fmt.Println("token:", os.Getenv("botToken"))
	fmt.Println("Prefix:", os.Getenv("botPrefix"))
	fmt.Println("-------------------------------------------------------------------------------------")

	fmt.Println("Running MathisBurger/commander version: v0.0.5")

	// Initialize bot session
	session, err := discordgo.New("Bot " + os.Getenv("botToken"))
	if err != nil {
		fmt.Println("Error while creating discord session")
		return
	}

	// Init command handler
	handler := commander.New(os.Getenv("botPrefix"), "")

	// register commands
	handler.Register("info", "Information command", commands.InfoCommand, 100)
	handler.Register("top", "Returns the top 10 words of guild", commands.TopCommand, 100)

	// add command handler and word collector
	session.AddHandler(handler.Execute)
	session.AddHandler(collector.Collector)

	// Configures the bot to run infinitely
	// and stops if you press CTRL+C
	session.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsGuildMessages)
	err = session.Open()
	if err != nil {
		fmt.Println("Cannot connect to discord websocket")
		return
	}

	// sets the activity of the bot
	err = session.UpdateListeningStatus("chat")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("The bot is running now. Terminate by using CTRL-C")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	session.Close()
}
