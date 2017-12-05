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
		today()
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

func today() {
	t := time.Now()
	today := t.Weekday()

	_, dayTotal, err := scanFile()
	checkError(err)

	for i := 0; i < len(dayTotal); i++ {
		if i != int(today) {
			continue
		} else {
			fmt.Println("Total time today so far: ", dayTotal[i].String())
		}
	}

}

func week() {
	totalTime, dayTotal, err := scanFile()
	checkError(err)

	for i := 0; i < 7; i++ {
		weekday := time.Weekday(i)
		total := dayTotal[i]

		if dayTotal[i] > 0 {
			fmt.Println(weekday, total)
		}
	}

	fmt.Println("Total time this week so far: ", totalTime.String())
}

func scanFile() (totalTime time.Duration, dayTotal [7]time.Duration, err error) {
	t := time.Now()

	file, err := os.Open(fullPath(t))
	checkError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var in []time.Time
	var out []time.Time

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
		diff := out[i].Sub(in[i])
		totalTime += diff
		dayTotal[in[i].Weekday()] += diff
	}

	if len(in) > len(out) {
		diff := t.Sub(in[len(in)-1])
		totalTime += diff
		dayTotal[in[len(in)-1].Weekday()] += diff
	}
	return
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
