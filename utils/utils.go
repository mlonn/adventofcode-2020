package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

// Input returns the input for the specified year and day as a string,
// downloading it if it does not already exist on disk.
func Input(year, day int) string {
	os.Mkdir("inputs", 0755)
	filename := fmt.Sprintf("inputs/%v.txt", day)
	if _, err := os.Stat(filename); err != nil {
		est, err := time.LoadLocation("EST")
		if err != nil {
			panic(err)
		}
		if t := time.Date(year, time.December, day, 0, 0, 0, 0, est); time.Until(t) > 0 {
			fmt.Printf("Puzzle not unlocked yet! Sleeping for %v...\n", time.Until(t))
			time.Sleep(time.Until(t) + 3*time.Second) // don't want to fire too early
		}
		fmt.Println("Downloading input...")
		req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("https://adventofcode.com/%v/day/%v/input", year, day), nil)
		req.AddCookie(&http.Cookie{
			Name:  "session",
			Value: os.Getenv("AOC_SESSION"),
		})
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		if err := ioutil.WriteFile(filename, data, 0660); err != nil {
			panic(err)
		}
	}
	return ReadInput(filename)
}

// ReadInput returns the contents of filename as a string.
func ReadInput(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(bytes.TrimSpace(data))
}

