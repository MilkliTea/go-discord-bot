package helpers

import (
	"fmt"
	"go-discord-bot/commands"
	"math/rand"
	"time"
)

func GenerateRandomText(sleepTime int, length int) {
	rand.Seed(time.Now().UnixNano())

	// Rastgele harfleri içeren bir karakter dizisi
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// Metin uzunluğuna kadar döngü
	randomText := make([]byte, length)
	for i := 0; i < length; i++ {
		// Rastgele bir indeks seç ve karakter dizisinden bir harf al
		randomIndex := rand.Intn(len(charset))
		randomChar := charset[randomIndex]

		// Seçilen harfi metin dizisine ekle
		randomText[i] = randomChar
	}

	text := fmt.Sprintf("%d sn cooldown. %s", sleepTime, randomText)

	commands.SendFarmMessageToMainChannel(text)
}

func GenerateRandomNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano())

	sleepTime := rand.Intn(max-min+1) + min

	return sleepTime
}
