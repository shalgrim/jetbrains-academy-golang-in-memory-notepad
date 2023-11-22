package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter the maximum number of notes:")
	scanner.Scan()
	maxNotes, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal("Unable to get maxNotes")
	}
	notes := make([]string, 0, maxNotes)

	for {
		fmt.Println("Enter a command and data:")
		scanner.Scan()
		input := strings.Split(scanner.Text(), " ")
		command, data := input[0], strings.Join(input[1:], " ")
		switch command {
		case "create":
			if len(notes) == maxNotes {
				fmt.Println("[Error] Notepad is full")
			} else {
				if data == "" {
					fmt.Println("[Error] Missing note argument")
				} else {
					notes = append(notes, data)
					fmt.Println("[OK] The note was successfully created")
				}
			}
		case "list":
			if len(notes) == 0 {
				fmt.Println("[Info] Notepad is empty")
			} else {
				for i, n := range notes {
					fmt.Printf("[Info] %d: %s\n", i+1, n)
				}
			}
		case "clear":
			notes = make([]string, 0, 5)
			fmt.Println("[OK] All notes were successfully deleted")
		case "exit":
			fmt.Println("[Info] Bye!")
			os.Exit(0)
		default:
			fmt.Println("[Error] Unknown command")
		}
	}
}
