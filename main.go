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
	"regexp"
	"strings"
	"syscall"
	"time"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		return
	}

	bot, err := discordgo.New("Bot " + os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}
	var verifyCaptcha = true
	bot.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate) {

		//captcha geldiğinde botu durdurur.
		if strings.Contains(m.Content, "captcha") {
			verifyCaptcha = false
			s.ChannelMessageSend(m.ChannelID, "<@"+os.Getenv("AUTHOR_ID")+"> hocam bi buraya bak hele yine geldi")
		}

		// gem bittiği zaman çalışır.
		if strings.Contains(m.Content, "spent") && strings.Contains(m.Content, "caught") {
			sendFarmMessage("owo inv")
		}

		if strings.Contains(m.Content, "Inventory") {
			s.ChannelMessageSend(m.ChannelID, "gem bitmiş takviye yapılıyor")

			var text string = "owo use "
			for i := 1; i < 5; i++ {
				if i == 2 {
					continue
				}

				//inventorydaki gemleri listeler ve en yüksek değerli olanı kullanmak için parse eder
				regexpString := fmt.Sprintf("(\\d+)`<:(?:c|u|l|r|e|m|f)?gem%d:\\d+>", i)
				re := regexp.MustCompile(regexpString)
				matches := re.FindAllStringSubmatch(m.Content, -1)

				var result string
				for _, match := range matches {
					result += strings.Replace(match[1], "0", "", -1) + " "
				}

				nums := strings.Split(result, " ")

				if len(nums) == 0 {
					continue
				}

				text += " " + nums[len(nums)-2]
			}
			sendFarmMessage(text)
		}

		if m.Content == "sa" {
			verifyCaptcha = true
			s.ChannelMessageSend(m.ChannelID, "as ben bot")
		}

		if m.Content == "dur" {
			verifyCaptcha = false
			s.MessageReactionAdd(m.ChannelID, m.ID, "\U0001F44D")
		}

		//farmı başlatır
		if m.Content == "owoh" {
			for {
				if !verifyCaptcha {
					break
				}

				time.Sleep(30 * time.Second)
				sendFarmMessage("owo h")
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

func sendFarmMessage(content string) {
	url := os.Getenv("CHANNEL_URL")
	wcfMessage := map[string]interface{}{
		"content": content,
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
