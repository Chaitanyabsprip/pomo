package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/chaitanyabsprip/pomo/cache"
	"github.com/chaitanyabsprip/pomo/handler"
)

func parseDuration(args []string) time.Duration {
	duration, _ := time.ParseDuration("52m")
	if len(args) > 1 {
		dStr := args[1]
		if dStr == "hr" || dStr == "hour" {
			now := time.Now()
			nextHour := time.Date(now.Year(), now.Month(), now.Day(), now.Hour()+1, 0, 0, 0, time.Local)
			duration = time.Until(nextHour)
		} else {
			duration, _ = time.ParseDuration(dStr)
		}
	}
	return duration
}

func main() {
	flag.Parse()

	args := flag.Args()
	var cmd string
	if len(args) > 0 {
		cmd = flag.Args()[0]
	}

	cache.Setup()
	switch cmd {
	case "start":
		duration := parseDuration(args)
		handler.Start(duration)
	case "pause":
		handler.Pause()
	case "stop":
		handler.Stop()
	case "":
		handler.ShowTime()
	default:
		fmt.Fprintln(os.Stderr, "Invalid command")
	}
}
