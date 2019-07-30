package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func initDB() {
	var err error
	db, err = sql.Open("sqlite3", "bebot.sqlite3")
	if err != nil {
		log.Panic(err)
	}

	if err := db.Ping(); err != nil {
		log.Panic(err)
	}
}

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

// Save - Save the `ChatMessage` to the database
func (c ChatMessage) Save() int64 {
	// if the database isn't yet initialized, do so
	if db == nil {
		initDB()
	}

	// dump, _ := string(json.Marshal(c))

	insert, err := db.Prepare("INSERT INTO logs (timestamp, server_id, server, channel_id, channel, user_id, user, message) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)")
	if err != nil {
		fmt.Println(err)
	}

	res, err := insert.Exec(c.Timestamp, c.ServerID, c.Server, c.ChannelID, c.Channel, c.UserID, c.User, c.Message)
	if err != nil {
		fmt.Println(err)
	}

	rowID, _ := res.LastInsertId()
	return rowID
}
