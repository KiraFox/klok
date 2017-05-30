package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

var loc = fmt.Sprintf("%s/.local/share/klok/timesheet.txt", os.Getenv("HOME"))

func main() {
	command := os.Args[1]

	switch command {
	case "in":
		logTime("in")
	case "out":
		logTime("out")
	case "today":
		fmt.Println("got today")
	case "week":
		fmt.Println("got week")
	default:
		fmt.Println("What?")
	}
}

func logTime(key string) {
	t := time.Now()

	file, err := os.OpenFile(loc, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	checkError(err)
	defer file.Close()

	writer := bufio.NewWriter(file)

	_, err = writer.WriteString(key + ": " + t.Format(time.RFC3339) + "\n")
	checkError(err)

	err = writer.Flush()
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
