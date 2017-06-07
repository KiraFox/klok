package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
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
		week()
	default:
		fmt.Println("What?")
	}
}

func logTime(key string) {
	t := time.Now()

	err := os.MkdirAll(dir, 0755)
	checkError(err)

	file, err := os.OpenFile(fullPath(t), os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	checkError(err)
	defer file.Close()

	writer := bufio.NewWriter(file)

	_, err = writer.WriteString(key + ": " + t.Format(time.RFC3339) + "\n")
	checkError(err)

	err = writer.Flush()
	checkError(err)
}

func week() {
	t := time.Now()

	file, err := os.Open(fullPath(t))
	checkError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var in []time.Time
	var out []time.Time
	var timeDiff []time.Duration
	var totalTime time.Duration

	for scanner.Scan() {
		text := scanner.Text()
		currIsIn, currTime, err := parseEntry(text)
		checkError(err)

		if currIsIn {
			in = append(in, currTime)
		} else {
			out = append(out, currTime)
		}
	}

	for i := 0; i < len(out); i++ {
		timeDiff = append(timeDiff, (out[i].Sub(in[i])))
	}

	for _, x := range timeDiff {
		totalTime += x
	}

	if len(in) > len(out) {
		totalTime += t.Sub(in[len(in)-1])
	}

	fmt.Println("Total time this week so far: ", totalTime.String())
}

func fullPath(t time.Time) string {
	yr, wk := t.ISOWeek()
	filename := fmt.Sprintf("%d-wk%d.txt", yr, wk)
	fullPath := path.Join(dir, filename)
	return fullPath
}

func parseEntry(line string) (isIn bool, timeStamp time.Time, err error) {
	isIn = strings.HasPrefix(line, "in: ")
	var timeText string

	if isIn {
		timeText = line[4:]
	} else if strings.HasPrefix(line, "out: ") {
		timeText = line[5:]
	} else {
		err = errors.New("Unknown entry type.")
		return
	}

	timeStamp, err = time.Parse(time.RFC3339, timeText)
	return

}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
