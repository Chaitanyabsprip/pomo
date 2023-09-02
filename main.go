package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/chaitanyabsprip/pomo/cache"
)

var cacheFile = os.Getenv("HOME") + "/.cache/pomotimer"

func stringf(t time.Duration) string {
	dStr := t.String()
	h := strings.Split(dStr, "h")[0] + ":"
	m := strings.Split(dStr, "m")[0] + ":"
	s := strings.Split(dStr, "s")[0]
	return fmt.Sprintf("%s%s%s", h, m, s)
}

func start(duration time.Duration) {
	now := time.Now()
	timer, err := cache.GetTimer()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Something went wrong")
		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(1)
	}
	var d time.Duration
	if timer.IsNil() {
		d = duration
		defer fmt.Fprintf(os.Stdout, "Started timer for %s\n", duration)
	} else if timer.IsPaused {
		d = timer.Duration
		defer fmt.Fprintf(os.Stdout, "Resumed timer for %s\n", duration)
	} else {
		fmt.Fprintln(os.Stderr, "Timer already running", duration)
		os.Exit(1)
	}
	cache.SetTime(cache.Timer{Duration: d, Start: now})
}

func pause() {
	timer, err := cache.GetTimer()
	if err != nil || timer.IsNil() {
		fmt.Fprintln(os.Stderr, "No timer running")
		os.Exit(1)
	}
	duration := time.Until(timer.Start.Add(timer.Duration))
	duration = duration.Round(time.Second)
	cache.SetTime(cache.Timer{
		Duration: duration,
		IsPaused: true,
		Start:    timer.Start,
	})
}

func stop() {
	t, _ := cache.GetTimer()
	if t.IsNil() {
		fmt.Fprintln(os.Stdout, "No timer running")
		os.Exit(1)
	}
	err := cache.Clear()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't stop timer")
		os.Exit(1)
	}
	fmt.Fprintln(os.Stdout, "Timer stopped")
}

func showTime() {
	timer, err := cache.GetTimer()
	if err != nil || timer.IsNil() {
		fmt.Fprint(os.Stdout, "")
		return
	}
	if timer.IsPaused {
		fmt.Fprintf(os.Stdout, "⏸️ %s\n", timer.Duration)
		return
	}
	duration := time.Until(timer.Start.Add(timer.Duration))
	duration = duration.Round(time.Second)
	fmt.Fprintf(os.Stdout, "🍅%s\n", stringf(duration))
}

func main() {
	cmd := ""
	duration, _ := time.ParseDuration("52m")
	if len(os.Args) > 1 {
		cmd = os.Args[1]
	}

	cache.Setup()
	switch cmd {
	case "start":
		start(duration)
	case "pause":
		pause()
	case "stop":
		stop()
	case "":
		showTime()
	default:
		fmt.Fprintln(os.Stderr, "Invalid command")
	}
}
