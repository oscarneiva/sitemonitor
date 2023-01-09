package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	printIntro()
	for {
		printMenu()
		command := readCommand()
		switch command {
		case 1:
			fmt.Println("Monitor...")
			startMonitor()
		case 2:
			fmt.Println("Logs...")
			printLogs()
		case 0:
			fmt.Println("Exit...")
			os.Exit(0)
		default:
			fmt.Println("Command doesn't exist...")
			os.Exit(-1)
		}
	}
}

func printIntro() {
	name := "oscar"
	version := 1.11

	fmt.Println("Hello Mr. ", name)
	fmt.Println("This program is at version ", version)
}

func printMenu() {
	fmt.Println("1 - Start monitor")
	fmt.Println("2 - Logs")
	fmt.Println("3 - Exit")
}

func readCommand() int {
	var command int
	fmt.Scan(&command)

	fmt.Println("The selected command was: ", command)
	return command
}

func startMonitor() {
	fmt.Println("Starting monitor...")
	urls := readUrls()

	for i := 0; i < len(urls); i++ {
		resp, err := http.Get(urls[i])
		if err != nil {
			fmt.Println("Error ", err)
		}
		if resp.StatusCode == 200 {
			fmt.Println(urls[i], " RUNNING")
			createLogs(urls[i], true)
		} else {
			fmt.Println(urls[i], " DOWN")
			createLogs(urls[i], false)
		}
	}
}

func readUrls() []string {
	var urls []string
	file, err := os.Open("./main/urls.txt")
	if err != nil {
		fmt.Println("Error ", err)
	}
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		urls = append(urls, line)
		if err == io.EOF {
			break
		}
	}
	file.Close()
	return urls
}

func createLogs(url string, status bool) {
	file, err := os.OpenFile("./main/logs.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}
	file.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + url + " online: " + strconv.FormatBool(status) + "\n")
	file.Close()
}

func printLogs() {
	file, err := ioutil.ReadFile("./main/logs.txt")
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(string(file))
}
