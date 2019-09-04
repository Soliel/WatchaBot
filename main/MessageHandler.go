package main

import (
	"github.com/bwmarrin/discordgo"

	"strings"

	"github.com/soliel/WatchaBot/Command"
)

func filterMessages(s *discordgo.Session, m *discordgo.MessageCreate) Command.Message {
	var CommandMsg Command.Message

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

	CommandName := content[:strings.Index(content, " ")]
	CommandName = strings.ToLower(CommandName)

	content = content[len(CommandName)+1:]

	CommandMsg = Command.Message{Command: CommandName, Content: content}

	return CommandMsg
}
