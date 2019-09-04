package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"

	"github.com/soliel/WatchaBot/Command"
	"github.com/soliel/WatchaBot/Configuration"
)

var (
	conf    *Configuration.BotConfig
	handler *Command.Handler
)

func main() {
	loadConfBytes, err := ioutil.ReadFile("../ConfigurationFiles/WatchaConf.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	conf = new(Configuration.BotConfig)
	err = conf.LoadConfig(loadConfBytes)
	if err != nil {
		fmt.Println("Error getting bot congifuration: ", err)
	}

	fmt.Println(conf.BotToken)

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

	handler = Command.CreateHandler()
	registerCommands()

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
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

func ping(context Command.Context) {
	context.Session.ChannelMessageSend(context.Channel.ID, "PONG")
}
