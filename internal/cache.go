// Package pomo provides pomo  
package pomo

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// Timer struct  
type Timer struct {
	Start    time.Time
	Duration time.Duration
	IsPaused bool
}

// IsNil method  
func (t *Timer) IsNil() bool { return t.Duration == 0 && t.Start.IsZero() }

var cacheFile = os.Getenv("HOME") + "/.cache/pomodoro/" + filepath.Base(os.Args[0]) + "timer"

// Setup function  
func Setup() error {
	if _, err := os.Stat(filepath.Dir(cacheFile)); os.IsNotExist(err) {
		if err := os.Mkdir(filepath.Dir(cacheFile), 0o755); err != nil {
			return err
		}
	}
	if _, err := os.Stat(cacheFile); os.IsNotExist(err) {
		fmt.Println(cacheFile)
		if err := os.WriteFile(cacheFile, []byte{}, 0o644); err != nil {
			return err
		}
	}
	return nil
}

func parseTimer(cache map[string]string) (Timer, error) {
	var duration time.Duration
	var isPaused bool
	var err error
	if duration, err = time.ParseDuration(cache["duration"]); err != nil {
		fmt.Printf("duration %s\n", err.Error())
		return Timer{}, err
	}
	if isPaused, err = strconv.ParseBool(cache["is_paused"]); err != nil {
		fmt.Printf("is_paused %s\n", err.Error())
		return Timer{}, err
	}
	unixTime, err := strconv.ParseInt(cache["start"], 10, 64)
	if err != nil {
		fmt.Printf("unix time %s\n", err.Error())
		return Timer{}, err
	}
	start := time.Unix(unixTime, 0)
	return Timer{
		Duration: duration,
		IsPaused: isPaused,
		Start:    start,
	}, nil
}

// GetTimer function  
func GetTimer() (Timer, error) {
	file, err := os.Open(cacheFile)
	if err != nil {
		fmt.Println("couldn't open file")
		return Timer{}, err
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil || stat.Size() == 0 {
		return Timer{}, err
	}
	scanner := bufio.NewScanner(file)
	cache := make(map[string]string)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ":")
		key, value := line[0], line[1]
		cache[key] = value
	}
	return parseTimer(cache)
}

// SetTime function  
func SetTime(t Timer) error {
	// 0644 is used to denote that only the own can read and write the file, other
	// users can only read the file.
	file, err := os.OpenFile(cacheFile, os.O_RDWR, 0o644)
	if err != nil {
		return err
	}
	defer file.Close()
	w := bufio.NewWriter(file)
	fmt.Fprintf(w, "duration:%s\n", t.Duration.String())
	fmt.Fprintf(w, "start:%s\n", fmt.Sprint(t.Start.Unix()))
	fmt.Fprintf(w, "is_paused:%s\n", strconv.FormatBool(t.IsPaused))
	w.Flush()
	return nil
}

// Clear function  
func Clear() error {
	if err := os.Truncate(cacheFile, 0); err != nil {
		return err
	}
	return nil
}
