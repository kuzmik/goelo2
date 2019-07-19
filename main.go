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
	//Debug - dump objects to console for debugging purposes
	Debug bool

	//Token - the authentication token that was provided when the bot was created in discord. Stored in 1password
	Token string

	//TokenFile - file that contains the token, to be read on program start. Keeps the token out of the processlist for shared hosts
	TokenFile string
)

func init() {
	flag.BoolVar(&Debug, "d", false, "Debug")
	flag.StringVar(&Token, "t", "", "Bot token")
	flag.StringVar(&TokenFile, "f", "env", "File that contains the bot token")
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

	dg.AddHandler(botReady)
	dg.AddHandler(messageCreate)
	dg.AddHandler(messageUpdate)

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
// `Message` is created on any `Channel` that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	// Ignore all messages created by the bot itself; not required, but a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	// Dump the `MessageCreate` struct to console
	if Debug == true {
		spew.Dump(m)
	}

	// Get the `Channel` name because APPARENTLY it isnt included in `m`
	channel := ""
	_channel, err := s.State.Channel(m.ChannelID)
	if err != nil {
		fmt.Println("Failure getting channel:", err)
	}

	// Don't accept DMs or any of the other channel types for now
	if _channel.Type != discordgo.ChannelTypeGuildText {
		return
	}

	if _channel.Name == "" {
		channel = "PRIVMSG"
	} else {
		channel = _channel.Name
	}

	// get the `Guild` which is the stupid name for a server
	server := ""
	if m.GuildID == "" {
		server = "PRIVMSG"
	} else {
		_guild, err := s.State.Guild(m.GuildID)

		if err != nil {
			fmt.Println("Failure getting guild:", err)
		}
		server = _guild.Name
	}

	// All that work to print this to the console.
	fmt.Printf("[%v] [%s] [%s] [%s] %s\n", m.Message.Timestamp, server, channel, m.Author, m.Message.Content)

	// If the message is "ping" reply with "Pong!"
	if m.Message.Content == "ping" {
		_, err := s.ChannelMessageSend(m.ChannelID, "pong")
		if err != nil {
			fmt.Println("Failure sending message:", err)
		}
	}
}

// This function will be called (due to AddHandler above) every time a
// `Message` is changed on any `Channel` that the autenticated bot has access to.
func messageUpdate(s *discordgo.Session, m *discordgo.MessageUpdate) {
	if Debug == true {
		spew.Dump(m)
	}
}

// This function will be called (due to AddHandler above) when the bot receives
// the "ready" event from Discord.
func botReady(s *discordgo.Session, event *discordgo.Ready) {
	// Set the playing status... for fun?
	s.UpdateStatus(0, "!honk")
}
