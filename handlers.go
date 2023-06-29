package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

var (
	commandHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){

		"eqb": func(s *discordgo.Session, i *discordgo.InteractionCreate) {
			// Access options in the order provided by the user.
			options := i.ApplicationCommandData().Options

			// Or convert the slice into a map
			optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
			for _, opt := range options {
				optionMap[opt.Name] = opt
			}

			book := "liber-al"
			words := ""
			sum := false
			var count int64 = 10 // default to 10

			// Get the value from the option map.
			// When the option exists, ok = true
			if option, ok := optionMap["words"]; ok {
				words = option.StringValue()
			}

			if option, ok := optionMap["book"]; ok {
				book = option.StringValue()
			}

			if option, ok := optionMap["count"]; ok {
				count = option.IntValue()
			}

			if option, ok := optionMap["sum"]; ok {
				sum = option.BoolValue()
			}

			if count > 32 {
				count = 32
			}

			// Take our message and pass it into parser
			output, err := Parse(words, int(count), book, sum)
			if err != nil {
				fmt.Println(err)
				return
			}

			s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				// Ignore type for now, they will be discussed in "responses"
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &discordgo.InteractionResponseData{
					Content: output,
				},
			})
		},
	}
)
