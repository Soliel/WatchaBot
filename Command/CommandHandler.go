package Command

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

//The Context of the situation the command was used in.
type Context struct {
	Msg     *discordgo.MessageCreate
	Author  *discordgo.User
	Session *discordgo.Session
	Guild   *discordgo.Guild
	Channel *discordgo.Channel
	Content string
}

//Message is used to bundle command string and message string
type Message struct {
	Command string
	Content string
}

type commandFunc func(Context)

type commandMap map[string]commandFunc

//Handler is the container to hold the locations of all commands
type Handler struct {
	commands commandMap
}

//FailedToRegisterError is thrown when registration fails for any reason.
type FailedToRegisterError struct {
	command string
	desc    string
}

//NotHandledError is fired when a command does not have a suitable handler.
type NotHandledError struct {
	command string
	desc    string
}

func (e FailedToRegisterError) Error() string {
	return fmt.Sprintf("Command: %s failed to register. Reason: %s", e.command, e.desc)
}

func (e NotHandledError) Error() string {
	return fmt.Sprintf("Command: %s was unable to start. Reason: %s", e.command, e.desc)
}

//CreateHandler initializes the command handler object to store our commands.
func CreateHandler() *Handler {
	return &Handler{make(commandMap)}
}

//Register adds a command into the Handler's map
func (handler Handler) Register(name string, command commandFunc) error {
	if len(name) <= 3 {
		err := FailedToRegisterError{name, "Command name is less then 4 letters."}
		return err
	}

	if strings.Contains(name, " ") {
		err := FailedToRegisterError{name, "Command contains space"}
		return err
	}

	lwrName := strings.ToLower(name)

	_, cmdFound := handler.commands[lwrName]
	if cmdFound {
		err := FailedToRegisterError{lwrName, "Command already exists."}
		return err
	}

	handler.commands[lwrName] = command

	_, abrvFound := handler.commands[lwrName[:3]]

	if !abrvFound {
		handler.commands[lwrName[:3]] = command
	}

	return nil
}

//HandleCommand is used to process a message into context and activate the appropriate command
func (handler Handler) HandleCommand(m *discordgo.MessageCreate, s *discordgo.Session, command Message) error {
	cmdFunc, found := handler.commands[command.Command]
	if !found {
		fmt.Println("command not found")
		fmt.Println(handler.commands)
		return nil
	}

	channel, err := s.State.Channel(m.ChannelID)
	if err != nil {
		return NotHandledError{
			command: command.Command,
			desc:    "Unable to get channel information: " + err.Error(),
		}
	}

	guild, err := s.State.Guild(m.GuildID)
	if err != nil {
		return NotHandledError{
			command: command.Command,
			desc:    "Unable to get guild information: " + err.Error(),
		}
	}

	ctx := Context{
		Msg:     m,
		Author:  m.Author,
		Session: s,
		Guild:   guild,
		Channel: channel,
		Content: command.Content,
	}

	go cmdFunc(ctx)
	return nil
}
