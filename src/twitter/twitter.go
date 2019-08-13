package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"os"
// 	"os/signal"
// 	"syscall"

// 	// "github.com/davecgh/go-spew/spew"
// 	"github.com/dghubble/go-twitter/twitter"
// 	"github.com/dghubble/oauth1"
// )

// type Config struct {
// 	Twitter struct {
// 		ScreenName     string `json:"screen_name"`
// 		ConsumerKey    string `json:"consumer_key"`
// 		ConsumerSecret string `json:"consumer_secret"`
// 		AccessToken    string `json:"access_token"`
// 		AccessSecret   string `json:"access_secret"`
// 	} `json:"twitter"`
// }

// func init() {
// 	jsonData, err := ioutil.ReadFile("config/secrets.json")
// 	if err != nil {
// 		fmt.Println("Error reading JSON data:", err)
// 		return
// 	}

// 	var cfg Config
// 	json.Unmarshal(jsonData, &cfg)

// 	if cfg.Twitter.ConsumerKey == "" || cfg.Twitter.ConsumerSecret == "" || cfg.Twitter.AccessToken == "" || cfg.Twitter.AccessSecret == "" {
// 		log.Fatal("Consumer key/secret and Access token/secret required")
// 	}

// 	config := oauth1.NewConfig(cfg.Twitter.ConsumerKey, cfg.Twitter.ConsumerSecret)
// 	token := oauth1.NewToken(cfg.Twitter.AccessToken, cfg.Twitter.AccessSecret)
// 	// OAuth1 http.Client will automatically authorize Requests
// 	httpClient := config.Client(oauth1.NoContext, token)

// 	// Twitter Client
// 	client := twitter.NewClient(httpClient)

// 	// Convenience Demux demultiplexed stream messages
// 	demux := twitter.NewSwitchDemux()

// 	demux.Tweet = func(tweet *twitter.Tweet) {
// 		fmt.Printf("[%d] (@%s) %s\n", tweet.ID, tweet.User.ScreenName, tweet.Text)
// 	}

// 	fmt.Println("Starting Stream...")

// 	// handle tweets that mention the bot username
// 	params := &twitter.StreamFilterParams{
// 		Track:         []string{cfg.Twitter.ScreenName},
// 		StallWarnings: twitter.Bool(true),
// 	}
// 	stream, err := client.Streams.Filter(params)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	go demux.HandleChan(stream.Messages)

// 	// Wait for SIGINT and SIGTERM (HIT CTRL-C)
// 	ch := make(chan os.Signal)
// 	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
// 	log.Println(<-ch)

// 	fmt.Println("Stopping Stream...")
// 	stream.Stop()
// }
