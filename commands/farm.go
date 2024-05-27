package commands

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

func SendFarmMessage(content string) bool {
	url := os.Getenv("CHANNEL_URL")
	wcfMessage := map[string]interface{}{
		"content": content,
		"nonce":   time.Now().Format("202307080151269192"), // her seferinde farklı olmalı bir nevi message ID'si gibi bir şey
		"tts":     false,
	}

	jsonValue, _ := json.Marshal(wcfMessage)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonValue))

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "text/plain")
	req.Header.Set("Authorization", os.Getenv("BEARER_TOKEN"))

	client := &http.Client{}
	resp, _ := client.Do(req)

	if resp.StatusCode != 200 {
		return false
	}

	defer resp.Body.Close()

	defer ioutil.ReadAll(resp.Body)

	//log.Print(ioutil.ReadAll(resp.Body))
	return true
}

func SendBattleFarmText(isBattleFriends bool) {
	battleText := "owo b"
	if isBattleFriends {
		battleText = "owo b " + "<@" + os.Getenv("OTHER_AUTHOR_ID") + ">"
	}

	SendFarmMessage(battleText)
}

func SellWeapons() {
	SendFarmMessage("owo wc all")
	time.Sleep(time.Duration(5) * time.Second)
	SendFarmMessage("owo sell uncommonweapons")
	time.Sleep(time.Duration(3) * time.Second)
	SendFarmMessage("owo sell commonweapons")
	time.Sleep(time.Duration(3) * time.Second)
	SendFarmMessage("owo sell rareweapons")
	time.Sleep(time.Duration(3) * time.Second)
	SendFarmMessage("owo sell epicweapons")
}

func UpdateGems(inventory string) {
	var text string = "owo use "

	for i := 1; i < 5; i++ {
		if i == 2 {
			continue
		}

		//inventorydaki gemleri listeler ve en yüksek değerli olanı kullanmak için parse eder
		regexpString := fmt.Sprintf("(\\d+)`<:(?:c|u|l|r|e|m|f)?gem%d:\\d+>", i)
		re := regexp.MustCompile(regexpString)
		matches := re.FindAllStringSubmatch(inventory, -1)

		var result string
		for _, match := range matches {
			result += strings.Replace(match[1], "0", "", -1) + " "
		}

		nums := strings.Split(result, " ")

		if len(nums) < 2 {
			continue
		}

		text += " " + nums[len(nums)-2]
	}
	fmt.Println(text)

	SendFarmMessage(text)
}
