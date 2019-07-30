package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"

	// "github.com/davecgh/go-spew/spew"
	_ "github.com/mattn/go-sqlite3"
	"github.com/mb-14/gomarkov"
)

var db *sql.DB
var chain *gomarkov.Chain

// ChatMessage - Struct that olds the message data, for use in saving to the databaes
type ChatMessage struct {
	Timestamp time.Time `json:"timestamp"`
	ServerID  string    `json:"server_id,omitempty"`
	Server    string    `json:"server,omitempty"`
	ChannelID string    `json:"channel_id,omitempty"`
	Channel   string    `json:"channel,omitempty"`
	UserID    string    `json:"user_id"`
	User      string    `json:"user"`
	Message   string    `json:"message"`
}

func init() {
	var err error
	db, err = sql.Open("sqlite3", "data/bebot.sqlite3")
	if err != nil {
		log.Panic(err)
	}

	if err := db.Ping(); err != nil {
		log.Panic(err)
	}

	chain, err = loadModel("data/model.json")
	if err != nil {
		//model is either empty or has bad data
		chain = gomarkov.NewChain(1)
		saveModel()
	}
}

////
// Bebot stuff
////

// Save - Save the `ChatMessage` to the database
func (c ChatMessage) Save() int64 {
	insert, err := db.Prepare("INSERT INTO logs (timestamp, server_id, server, channel_id, channel, user_id, user, message) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)")
	if err != nil {
		fmt.Println(err)
	}

	res, err := insert.Exec(c.Timestamp, c.ServerID, c.Server, c.ChannelID, c.Channel, c.UserID, c.User, c.Message)
	if err != nil {
		fmt.Println(err)
	}

	rowID, _ := res.LastInsertId()

	chain.Add(strings.Split(c.Message, " "))

	go saveModel()

	return rowID
}

////
// Markov stuff
////

// Loads the model into the chains
func loadModel(modelFile string) (*gomarkov.Chain, error) {
	data, err := ioutil.ReadFile(modelFile)
	if err != nil {
		return chain, err
	}
	err = json.Unmarshal(data, &chain)
	if err != nil {
		return chain, err
	}

	return chain, nil
}

// Dumps the chains into a json file
func saveModel() {
	jsonObj, _ := json.Marshal(chain)
	err := ioutil.WriteFile("data/model.json", jsonObj, 0600)
	if err != nil {
		fmt.Println(err)
	}
}

// Babble - Return a string of markov generated text
func Babble() string {
	tokens := []string{gomarkov.StartToken}

	for tokens[len(tokens)-1] != gomarkov.EndToken {
		next, _ := chain.Generate(tokens[(len(tokens) - 1):])
		tokens = append(tokens, next)
	}

	return strings.Join(tokens[1:len(tokens)-1], " ")
}
