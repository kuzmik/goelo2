package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/davecgh/go-spew/spew"
)

var (
	//Token - the authentication token that was provided when the bot was created in discord. Stored in 1password
	Token string

	//TokenFile - file that contains the token, to be read on program start. Keeps the token out of the processlist for shared hosts
	TokenFile string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot token")
	flag.StringVar(&TokenFile, "f", "", "File that contains the bot token")
	flag.Parse()
}

func main() {
	// If a token file is specified, read that.
	if TokenFile != "" {
		dat, err := ioutil.ReadFile(TokenFile)
		if err != nil {
			panic(err)
		}

		Token = strings.TrimSpace(string(dat))
	}

	if Token == "" {
		fmt.Println("You need to specify a token with -t")
		return
	}

	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(messageCreate)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself; not required, but a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	// get the `Guild` which is the stupid name for a server
	guild, err := s.State.Guild(m.GuildID)
	if err != nil {
		fmt.Println("Failure getting guild:", err)
	}

	// get the `Channel` name because APPARENTLY it isnt included in `m`
	channel, err := s.State.Channel(m.ChannelID)
	if err != nil {
		fmt.Println("Failure getting channel:", err)
	}

	// All that work to print this to the console.
	fmt.Printf("[%v] [%s] [%s] [%s] %s\n", m.Message.Timestamp, guild.Name, channel.Name, m.Author, m.Message.Content)

	// Some debug output
	if m.Message.Content == ".debugMessage" {
		spew.Dump(m)
	}

	if m.Message.Content == ".debugState" {
		spew.Dump(m)
	}

	// If the message is "ping" reply with "Pong!"
	if m.Message.Content == "ping" {
		_, err := s.ChannelMessageSend(m.ChannelID, "pong")
		if err != nil {
			fmt.Println("Failure sending message:", err)
		}
	}
}
