package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	const MaxNotes = 5
	scanner := bufio.NewScanner(os.Stdin)
	notes := make([]string, 0, MaxNotes)

	for {
		fmt.Println("Enter a command and data:")
		scanner.Scan()
		input := strings.Split(scanner.Text(), " ")
		command, data := input[0], strings.Join(input[1:], " ")
		switch command {
		case "create":
			if len(notes) == MaxNotes {
				fmt.Println("[Error] Notepad is full")
			} else {
				notes = append(notes, data)
				fmt.Println("[OK] The note was successfully created")
			}
		case "list":
			for i, n := range notes {
				fmt.Printf("[Info] %d: %s\n", i+1, n)
			}
		case "clear":
			notes = make([]string, 0, 5)
			fmt.Println("[OK] All notes were successfully deleted")
		case "exit":
			fmt.Println("[Info] Bye!")
			os.Exit(0)
		}
	}
}
