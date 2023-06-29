package main

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

// ReadConfig function
func ReadConfig(path string) (err error) {

	var conf Config
	if _, err := toml.DecodeFile(path, &conf); err != nil {
		fmt.Println(err)
		return err
	}

	Conf = conf
	return nil
}

// Config struct
type Config struct {
	DiscordConfig discordConfig `toml:"discord"`
	Botconfig     botConfig     `toml:"bot"`
}

// discordConfig struct
type discordConfig struct {
	Token string `toml:"bot_token"`
}

// botConfig struct
type botConfig struct {
	Prefix string `toml:"command_prefix"`
}
