package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/soliel/WatchaBot/command"
	"github.com/soliel/WatchaBot/configuration"
)

var (
	conf    *configuration.BotConfig
	dbConf  *configuration.DatabaseConfig
	handler *command.Handler
)

func main() {
	loadBotConfBytes, err := ioutil.ReadFile("../ConfigurationFiles/WatchaConf.json")
	loadDbConfBytes, err := ioutil.ReadFile("../ConfigurationFiles/DbConf.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	conf = new(configuration.BotConfig)
	dbConf = new(configuration.DatabaseConfig)
	err = conf.LoadConfig(loadBotConfBytes)
	err = dbConf.LoadConfig(loadDbConfBytes)
	if err != nil {
		fmt.Println("Error getting bot congifuration: ", err)
	}

	dg, err := discordgo.New("Bot " + conf.BotToken)
	if err != nil {
		fmt.Println("Error starting discord session: ", err)
		return
	}

	dg.AddHandler(onMessageReceived)

	err = dg.Open()
	if err != nil {
		fmt.Println("Error opening communication with discord: ", err)
		return
	}
	defer dg.Close()

	db, err := gorm.Open("postgres", dbConf.CreateDatabaseString())
	if err != nil {
		fmt.Println("Error opening database connection: ", err)
		return
	}
	defer db.Close()

	handler = command.CreateHandler()
	registerCommands()

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
}

func onMessageReceived(s *discordgo.Session, m *discordgo.MessageCreate) {
	command := filterMessages(s, m)

	if command.Command == "" {
		return
	}

	handler.HandleCommand(m, s, command)
}

func registerCommands() {
	handler.Register("ping", ping)
}

func ping(context command.Context) {
	context.Session.ChannelMessageSend(context.Channel.ID, "PONG")
}
