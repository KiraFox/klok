package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"time"
)

var dir = fmt.Sprintf("%s/.local/share/klok", os.Getenv("HOME"))

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
	yr, wk := t.ISOWeek()
	filename := fmt.Sprintf("%d-wk%d.txt", yr, wk)
	fullPath := path.Join(dir, filename)

	err := os.MkdirAll(dir, 0755)
	checkError(err)

	file, err := os.OpenFile(fullPath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
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
