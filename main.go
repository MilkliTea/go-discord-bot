package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"log"
	"os"
	"os/signal"
	"syscall"
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
