package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/davecgh/go-spew/spew"
	"github.com/getsentry/sentry-go"
	"github.com/kuzmik/goelo2"
)

var (
	//Debug - dump objects to console for debugging purposes
	Debug bool

	//Token - the authentication token that was provided when the bot was created in discord. Stored in 1password
	Token string

	//TokenFile - file that contains the token, to be read on program start. Keeps the token out of the processlist for shared hosts
	TokenFile string
)

type DiscordConfig struct {
	Discord struct {
		ApiKey string `json:"api_key"`
	} `json:"discord"`
}

func init() {
	flag.BoolVar(&Debug, "d", false, "Debug mode prints extra data to the console")
	flag.StringVar(&Token, "t", "", "Discord bot token")
	flag.StringVar(&TokenFile, "f", "config/secrets.json", "File that contains the bot token")
	flag.Parse()

	// If a token file is specified, read that.
	if TokenFile != "" {
		jsonData, err := ioutil.ReadFile(TokenFile)
		if err != nil {
			fmt.Println("Error reading JSON data:", err)
			return
		}

		var cfg DiscordConfig
		json.Unmarshal(jsonData, &cfg)

		Token = cfg.Discord.ApiKey
	}

	// If no there is no token specified, either via commandline or via file, bail out
	if Token == "" {
		fmt.Println("You need to specify a token. Use --help for help")
		return
	}

	// Set up the sentry reportig
	sentry.Init(sentry.ClientOptions{
		Dsn: "https://3779a47dff1f4fb08d8c16e2f73f90f9@sentry.io/1509313",
	})
}

func main() {
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		handleError("Error creating bot", err)
		return
	}

	dg.AddHandler(botReady)
	dg.AddHandler(messageCreate)
	dg.AddHandler(messageUpdate)

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		handleError("Error during connecting:", err)
		return
	}

	// start up the twitter monitor
	//go startStream()

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

// Handles all errors, which includes sending to sentry
func handleError(message string, err error) {
	fmt.Println(message, err)
	sentry.CaptureException(err)
	sentry.Flush(time.Second * 5)
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

	timestamp, _ := m.Message.Timestamp.Parse()

	// All that work to print this to the console.
	fmt.Printf("[%v] [%s] [%s] [%s] %s\n", timestamp, server, channel, m.Author, m.Message.Content)

	author := fmt.Sprintf("%s#%s", m.Author.Username, m.Author.Discriminator)

	msg := ChatMessage{
		Timestamp: timestamp,
		ServerID:  m.Message.GuildID,
		Server:    server,
		ChannelID: m.ChannelID,
		Channel:   channel,
		UserID:    m.Author.ID,
		User:      author,
		Message:   m.Message.Content,
	}

	msg.Save()

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
