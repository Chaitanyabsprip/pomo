package pomo

import (
	"fmt"
	"os"
	"time"
)

// Start function  
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

// Pause function  
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

// Stop function  
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

// ShowTime function  
func ShowTime(alertDuration time.Duration) {
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
	var durationStr string
	if duration < 0 {
		durationStr = " - "
	} else {
		duration = duration.Round(time.Second)
		durationStr = duration.String()
	}
	fmt.Fprintf(os.Stdout, "%s\n", durationStr)
}
