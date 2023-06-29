package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the authenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	if !strings.HasPrefix(m.Content, Conf.Botconfig.Prefix) {
		fmt.Print("Prefix is: ", Conf.Botconfig.Prefix)
		return
	}

	m.Content = strings.TrimPrefix(m.Content, Conf.Botconfig.Prefix)
	// Remove whitespace from beginning and end of string
	m.Content = strings.TrimSpace(m.Content)

	// Take our message and pass it into parser
	output, err := Parse(m.Content, 32, "liber-al")
	if err != nil {
		fmt.Println(err)
		return
	}

	_, _ = s.ChannelMessageSend(m.ChannelID, output)

}
