package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	godotenv.Load()

	bot, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	bot.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {
		if m.Author.ID == s.State.User.ID {
			return
		}

		if m.Content == "sa" {
			s.ChannelMessageSend(m.ChannelID, "as ben bot")
		}

		if m.Content == "owoh" || m.Content == "owo h" && m.Author.ID == os.Getenv("AUTHOR_ID") && m.ChannelID == os.Getenv("CHANNEL_ID") {
			time.Sleep(15 * time.Second)
			sendFarmMessage()
		}
	})

	bot.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentsGuildMessages

	err = bot.Open()

	if err != nil {
		log.Fatal(err)
	}

	defer bot.Close()

	fmt.Printf("the bot is online")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc
}

func sendFarmMessage() {
	url := os.Getenv("CHANNEL_URL")
	wcfMessage := map[string]interface{}{
		"content": "owo h",
		"nonce":   time.Now().Format("20230708015126"), // her seferinde farklı olmalı bir nevi message ID'si gibi bir şey
		"tts":     false,
	}

	jsonValue, _ := json.Marshal(wcfMessage)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "text/plain")
	req.Header.Set("Authorization", os.Getenv("BEARER_TOKEN"))

	client := &http.Client{}
	resp, _ := client.Do(req)

	defer resp.Body.Close()

	defer ioutil.ReadAll(resp.Body)
}
