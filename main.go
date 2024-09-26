package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"dawar.com/note"
)

func main() {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")

	if title == "" || content == "" {
		fmt.Println("title and content are required")
		os.Exit(1)
		//end the program
		panic("title and content are required")
	}

	note := note.New(title, content)
	data := note.Print()
	fmt.Println(data)
	err := note.SaveAsJSON()
	if err != nil {
		fmt.Println("Error saving note")
		os.Exit(1)
	}

	fmt.Println("Note saved successfully")

}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	reader := bufio.NewReader(os.Stdin)
	value, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	text := strings.TrimSuffix(value, "\n")
	text = strings.TrimSuffix(text, "\r")
	return text
}
