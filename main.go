package main

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// Variables used for command line parameters
var (
	ConfPath string
	Conf     Config
)

func init() {
	// Read our command line options
	flag.StringVar(&ConfPath, "c", "eqb.conf", "Path to Config File")
	flag.Parse()

	_, err := os.Stat(ConfPath)
	if err != nil {
		log.Fatal("Config file is missing: ", ConfPath)
	}

	// Verify we can actually read our config file
	err = ReadConfig(ConfPath)
	if err != nil {
		log.Fatal("error reading config file at: ", ConfPath)
		return
	}

}

func main() {
	fmt.Println("\n\n|| Starting EQB - English Qabalah Bot ||")
	log.SetOutput(ioutil.Discard)

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Conf.DiscordConfig.Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}
	defer dg.Close()

	// Register the messageCreate func as a callback for MessageCreate events.
	dg.AddHandler(messageCreate)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}
