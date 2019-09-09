package main

import (
	"github.com/bwmarrin/discordgo"

	"strings"

	"github.com/soliel/WatchaBot/command"
)

func filterMessages(s *discordgo.Session, m *discordgo.MessageCreate) command.Message {
	var CommandMsg command.Message

	if m.Author.ID == s.State.User.ID {
		return CommandMsg
	}

	if len(m.Content) < len(conf.BotPrefix) {
		return CommandMsg
	}

	if m.Content[:len(conf.BotPrefix)] != conf.BotPrefix {
		return CommandMsg
	}

	content := m.Content[len(conf.BotPrefix):]
	if len(content) < 1 {
		return CommandMsg
	}

	LastCommandIndex := strings.Index(content, " ")
	if LastCommandIndex < 0 {
		LastCommandIndex = len(content)
	}
	CommandName := content[:LastCommandIndex]
	CommandName = strings.ToLower(CommandName)

	if len(CommandName) == len(content) {
		content = ""
	} else {
		content = content[len(CommandName)+1:]
	}

	CommandMsg = command.Message{Command: CommandName, Content: content}

	return CommandMsg
}
