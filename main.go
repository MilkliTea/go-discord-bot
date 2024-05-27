package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"go-discord-bot/commands"
	"go-discord-bot/helpers"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var (
	verifyCaptcha = true
	battleFriends = false
	fastMode      = false
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

	bot.AddHandler(handleMessage)

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

func handleMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return // Ignore messages sent by the bot itself
	}

	if helpers.ContainsCaptcha(m.Content) {
		stopBot(s, m.ChannelID)
	}

	if helpers.ContainsSpentCaught(m.Content) {
		commands.SendFarmMessage("owo inv")
	}

	if helpers.ContainsInventory(m.Content) {
		s.ChannelMessageSend(m.ChannelID, "gem bitmiş takviye yapılıyor")
		helpers.Sleep(2, false)
		commands.UpdateGems(m.Content)
	}

	switch strings.ToLower(m.Content) {
	case "sa":
		verifyCaptcha = true
		s.ChannelMessageSend(m.ChannelID, "as ben bot")
	case "owob fr":
		battleFriends = true
		s.ChannelMessageSend(m.ChannelID, "battle with friends aktif edildi")
		commands.SendFarmMessage("owoh")
	case "owo fast":
		fastMode = true
		s.ChannelMessageSend(m.ChannelID, "fast mode açıldı")
		commands.SendFarmMessage("owoh")
	case "dur":
		verifyCaptcha = false
		battleFriends = false
		s.MessageReactionAdd(m.ChannelID, m.ID, "\U0001F44D")
	case "owoh":
		s.ChannelMessageSend(m.ChannelID, "başlıyorum")
		startFarm(s, m.ChannelID)
	case "sell ww":
		commands.SellWeapons()
		s.ChannelMessageSend(m.ChannelID, "weaponlar satıldı")
	case "ping":
		checkToken(s, m, "ping")
	}
}

func stopBot(s *discordgo.Session, channelID string) {
	verifyCaptcha = false
	fastMode = false
	s.ChannelMessageSend(channelID, "<@"+os.Getenv("AUTHOR_ID")+"> hocam bi buraya bak hele yine geldi")
	s.ChannelMessageSend(channelID, "durdum.")
}

func startFarm(s *discordgo.Session, channelID string) {
	if !verifyCaptcha {
		return
	}

	for i := 0; verifyCaptcha; i++ {
		sleepTime := helpers.GenerateRandomNumber(30, 120)
		helpers.GenerateRandomText(sleepTime, 10)
		helpers.Sleep(sleepTime, fastMode)
		commands.SendFarmMessage("owo h")
		helpers.Sleep(1, false)
		commands.SendBattleFarmText(battleFriends)

		if (i+1)%10 == 0 {
			helpers.Sleep(2, false)
			commands.SendFarmMessage("owo pray")
			text := fmt.Sprintf("%d kere çalıştım azcık mola veriyorum", i+1)
			s.ChannelMessageSend(channelID, text)
			helpers.Sleep(240, fastMode)
		}
	}
}

func checkToken(s *discordgo.Session, m *discordgo.MessageCreate, message string) {
	send := commands.SendFarmMessage(message)

	if !send {
		s.ChannelMessageSend(m.ChannelID, "mesaj gönderilemedi. token kontrol ediniz.")
	}

	s.MessageReactionAdd(m.ChannelID, m.ID, "\U0001F44D")
}
