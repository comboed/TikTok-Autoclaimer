package main

import (
	"os/signal"
	"syscall"
	"strings"
	"time"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("[%s] Very Faste Turbowe xD %s\r\n\n", green("x_x"), pink("v3.0"))
	loadFiles()

	fmt.Printf("[%s] Sessions: %s\n", pink("x_x"), green(formatNumber(int64(len(sessions)))))
	fmt.Printf("[%s] Targets: %s\n", pink("x_x"), green(formatNumber(int64(len(usernames)))))

	fmt.Printf("[%s] Threads: ", pink("x_x"))
	fmt.Scanln(&threads)

	fmt.Printf("\n[%s] Press Enter to start", pink("x_x"))
	fmt.Scanln()

	fmt.Printf("\x1b[A                                      \x1b[A\n")
	
	httpClient = createTLSClient()

	var c chan os.Signal =  make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	for i := 0; i < threads; i++ {
		go Turbo()
	}

	fmt.Printf("[%s] All threads successfully initialized\n\n", green("x_x"))

	go func() {
		for {
			for _, v := range usernames {
				channel <- v
			}
		}
	}()


	go func() {
		for {
			var before int64 = attempts
			time.Sleep(time.Second * 1)
			rs = attempts - before
		}
	}()

	go func() {
		for {
			fmt.Printf("[%s] %s Attempts | RL: %s | R/S: %s%s\r", green("x_x"), formatNumber(attempts), formatNumber(rl), formatNumber(rs), strings.Repeat(" ", 10))
			time.Sleep(time.Millisecond * 100)
		}
	}()

	go func() {
		<-c
		fmt.Printf("[%s] Killing threads after {%s} attempts\n", green("x_x"), pink(formatNumber(attempts)))
		os.Exit(0)
	}()

	select {}
}