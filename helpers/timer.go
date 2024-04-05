package helpers

import "time"

func Sleep(duration int, isFastMod bool) {
	if isFastMod {
		time.Sleep(time.Duration(13) * time.Second)
		return
	}

	time.Sleep(time.Duration(duration) * time.Second)
}
