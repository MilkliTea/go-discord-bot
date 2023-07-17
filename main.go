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
	"strings"
	"syscall"
	"time"
	"math/rand"
	"strconv"
)

func main() {

	godotenv.Load()

	bot, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}
	var verifyCaptcha = true
	bot.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {

		if strings.Contains(m.Content, "captcha") {
			verifyCaptcha = false
			s.ChannelMessageSend(m.ChannelID, "<@"+os.Getenv("AUTHOR_ID")+"> hocam bi buraya bak hele yine geldi")
		}

		if m.Content == "sa" {
			verifyCaptcha = true
			s.ChannelMessageSend(m.ChannelID, "as ben bot")
		}

		if m.Content == "dur" {
			verifyCaptcha = false
			s.MessageReactionAdd(m.ChannelID, m.ID, "\U0001F44D")
		}

		if m.Content == "owoh" || m.Content == "owo h" {
		    for {
			if !verifyCaptcha {
			     break
			}

		 	time.Sleep(30 * time.Second)
			sendFarmMessage()
		    }
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

	//log.Print(ioutil.ReadAll(resp.Body))
}
