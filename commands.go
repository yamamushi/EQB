package main

import (
	"github.com/bwmarrin/discordgo"
)

var (
	Commands = []*discordgo.ApplicationCommand{
		{
			Name:        "eqb",
			Description: "Use English Qabalah Bot to search for words in a given book",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "input",
					Description: "Input word or phrase to generate results for",
					Required:    true,
				},
				// Required options must be listed first since optional parameters
				// always come after when they're used.
				// The same concept applies to Discord's Slash-commands API
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "book",
					Description: "Defaults to liber-al",
					Required:    false,
				},
				{
					Type:        discordgo.ApplicationCommandOptionBoolean,
					Name:        "sum",
					Description: "Only return sum value (true or false)",
					Required:    false,
				},
				{
					Type:        discordgo.ApplicationCommandOptionNumber,
					Name:        "count",
					Description: "Number of results to return, max 32",
					Required:    false,
				},
			},
		},
	}
)
