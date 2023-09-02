package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func stringf(t time.Duration) string {
	dStr := t.String()
	var h, m, s string
	var split []string
	if strings.Contains(dStr, "h") {
		split = strings.Split(dStr, "h")
		h = split[0] + ":"
		dStr = split[1]
	}
	if strings.Contains(dStr, "m") {
		split = strings.Split(dStr, "m")
		m = split[0] + ":"
		dStr = split[1]
	}
	if strings.Contains(dStr, "s") {
		split = strings.Split(dStr, "s")
		s = split[0]
	}
	return fmt.Sprintf("%s%s%s", h, m, s)
}

func Start(duration time.Duration) {
	now := time.Now()
	timer, err := GetTimer()
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
	SetTime(Timer{Duration: d, Start: now})
}

func Pause() {
	timer, err := GetTimer()
	if err != nil || timer.IsNil() {
		fmt.Fprintln(os.Stderr, "No timer running")
		os.Exit(1)
	}
	duration := time.Until(timer.Start.Add(timer.Duration))
	duration = duration.Round(time.Second)
	SetTime(Timer{
		Duration: duration,
		IsPaused: true,
		Start:    timer.Start,
	})
}

func Stop() {
	t, _ := GetTimer()
	if t.IsNil() {
		fmt.Fprintln(os.Stdout, "No timer running")
		os.Exit(1)
	}
	err := Clear()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Couldn't stop timer")
		os.Exit(1)
	}
	fmt.Fprintln(os.Stdout, "Timer stopped")
}

func ShowTime() {
	timer, err := GetTimer()
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
