package main

import (
	"math/rand"
	"strings"
	"strconv"
	"bufio"
	"fmt"
	"os"


)

func loadFiles() {
	var usernameFile, _ = os.Open("./data/usernames.txt")
	var sessionFile, _ = os.Open("./data/sessions.txt")

	defer usernameFile.Close()
	defer sessionFile.Close()
	
	var usernameScan *bufio.Scanner = bufio.NewScanner(usernameFile)
	var sessionScan *bufio.Scanner = bufio.NewScanner(sessionFile)

	for usernameScan.Scan() {
		usernames = append(usernames, usernameScan.Text())
	}
	for sessionScan.Scan() {
		sessions = append(sessions, sessionScan.Text())
	}
	buildData()
}

func buildData() {
	index = rand.Intn(len(sessions))
	var unformattedSession []string = strings.Split(sessions[index], ":")

	session = "sessionid="+unformattedSession[2]+"; sessionid_ss="+unformattedSession[2]+"; install_id=7024964356334372614"
	channel = make(chan string, len(usernames))
}

func claimFile(username string) {
	var format string = "Username: %s\nEmail: %s\nPassword: %s\nSession: %s\n"
	var account []string = strings.Split(sessions[index], ":")

	var f, _ = os.Create("./data/claimed/" + username + ".txt")
	fmt.Fprintln(f, fmt.Sprintf(format, username, account[0], account[1], account[2]))
}

func claimedUsername(username string) {
	fmt.Printf("[%s] Successfully autoclaimed %s after %s attempts%s\n\n", green(">_<"), green("@" + username), green(formatNumber(attempts)), strings.Repeat(" ", 25))
	discordWebhook(AUTOCLAIMED_WEBHOOK, fmt.Sprintf(AUTOCLAIMED_USERNAME_TEMPLATE, username))
	
	claimFile(username)
	
	if len(sessions) == 0 {
		fmt.Printf("[%s] Session finished - Killing threads...\n", green(">_<"))
		os.Exit(0)
	}
	sessions = removeIndex(sessions, index); buildData()
}

func isDeactivated(username string) {
	for i := 0; i < 6; i++ {
		if claimUsername(username) {
			claimedUsername(username)
		} else if i == 5 {
			if usernameAvailable(username) {
				fmt.Printf("[%s] %s has been suspended\n\n", pink(">_<"), pink("@" + username))
				usernames = removeString(usernames, username)
				newFile(usernames, "usernames")

				discordWebhook(MONITOR_WEBHOOK, fmt.Sprintf(SUSPENDED_USERNAME_TEMPLATE, username))
			} else {
				fmt.Printf("[%s] %s has been swapped to %s%s\n\n", pink(">_<"), pink("@" + username), pink(usernameID(username)), strings.Repeat(" ", 25))
				discordWebhook(MONITOR_WEBHOOK, fmt.Sprintf(SWAPPED_USERNAME_TEMPLATE, username, usernameID(username)))
			}
		}	
	}
}

func newFile(slice []string, fn string) {
	var f, _ = os.Create("./data/" + fn + ".txt")
	defer f.Close()
	for _, v := range slice {
		fmt.Fprintln(f, v)
	}
}

func removeIndex(slice []string, s int) []string {
    return append(slice[:s], slice[s+1:]...)
}

func removeString(s []string, r string) []string {
    for i, v := range s {
        if v == r {
            return append(s[:i], s[i+1:]...)
        }
    }
    return s
}

func reverseString(str string) string {
	var runes []rune = []rune(str)
	for i, j := 0, len(runes) - 1; i < j; i, j = i + 1, j - 1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func formatNumber(n int64) string {
    in := strconv.FormatInt(n, 10)
    numOfDigits := len(in)
    if n < 0 {
        numOfDigits--
    }
    numOfCommas := (numOfDigits - 1) / 3

    out := make([]byte, len(in)+numOfCommas)
    if n < 0 {
        in, out[0] = in[1:], '-'
    }

    for i, j, k := len(in)-1, len(out)-1, 0; ; i, j = i-1, j-1 {
        out[j] = in[i]
        if i == 0 {
            return string(out)
        }
        if k++; k == 3 {
            j, k = j-1, 0
            out[j] = ','
        }
    }
}