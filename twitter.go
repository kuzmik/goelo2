package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"os"
// 	"os/signal"
// 	"syscall"

// 	"github.com/davecgh/go-spew/spew"
// 	"github.com/dghubble/go-twitter/twitter"
// 	"github.com/dghubble/oauth1"
// )

// type TwitterConfig struct {
// 	Twitter struct {
// 		ScreenName     string `json:"screen_name"`
// 		ConsumerKey    string `json:"consumer_key"`
// 		ConsumerSecret string `json:"consumer_secret"`
// 		AccessKey      string `json:"access_token"`
// 		AccessSecret   string `json:"access_secret"`
// 	} `json:"twitter"`
// }

// var cfg TwitterConfig

// func init() {
// 	jsonData, err := ioutil.ReadFile("config/secrets.json")
// 	if err != nil {
// 		fmt.Println("Error reading JSON data:", err)
// 		return
// 	}

// 	json.Unmarshal(jsonData, &cfg)

// 	if cfg.Twitter.ConsumerKey == "" || cfg.Twitter.ConsumerSecret == "" || cfg.Twitter.AccessKey == "" || cfg.Twitter.AccessSecret == "" {
// 		log.Fatal("Configuration missing")
// 	}

// 	startStream()
// }

// func startStream() {
// 	oauth := oauth1.NewConfig(cfg.Twitter.ConsumerKey, cfg.Twitter.ConsumerSecret)
// 	token := oauth1.NewToken(cfg.Twitter.AccessKey, cfg.Twitter.AccessSecret)
// 	httpClient := oauth.Client(oauth1.NoContext, token)

// 	// main Twitter client
// 	client := twitter.NewClient(httpClient)

// 	// Convenience Demux demultiplexed stream messages
// 	demux := twitter.NewSwitchDemux()

// 	demux.Tweet = func(event *twitter.Tweet) {
// 		spew.Dump(event)
// 	}

// 	demux.DM = func(dm *twitter.DirectMessage) {
// 		fmt.Println(dm.SenderID)
// 	}

// 	demux.Event = func(event *twitter.Event) {
// 		fmt.Printf("%#v\n", event)
// 	}

// 	fmt.Println("Starting Stream...")
// 	stream, err := client.Streams.User(&twitter.StreamUserParams{
// 		StallWarnings: twitter.Bool(true),
// 		Language:      []string{"en"},
// 	})

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Receive messages until stopped or stream quits
// 	go demux.HandleChan(stream.Messages)

// 	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
// 	ch := make(chan os.Signal)
// 	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
// 	log.Println(<-ch)

// 	fmt.Println("Stopping Stream...")

// 	stream.Stop()
// }
