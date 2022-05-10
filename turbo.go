package main

import (
	"time"
)

func Turbo() {
	for {
		var username string = <- channel
		if usernameAvailable(username) && !lock {
			lock = true
			mutex.Lock()
			if claimUsername(username) {
				claimedUsername(username)
			} else {
				isDeactivated(username)
			}
			time.Sleep(time.Second * 2)
			mutex.Unlock()
			lock = false
		}
		attempts++
	}
}