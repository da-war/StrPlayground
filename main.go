package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"dawar.com/note"
	"dawar.com/todo"
)

type saver interface {
	Save() error
}

func main() {

	number := 10
	dawar := "Dawar"
	floatNumber := 10.5

	printAnyThing(number)
	printAnyThing(dawar)
	printAnyThing(floatNumber)

	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")
	todoText := getUserInput("Todo text: ")

	todo := todo.New(todoText)
	todo.Print()

	if title == "" || content == "" {
		fmt.Println("title and content are required")
		os.Exit(1)
		//end the program
		panic("title and content are required")
	}

	note := note.New(title, content)
	data := note.Print()
	fmt.Println(data)
	err := saveData(note)
	secondErr := saveData(todo)
	if err != nil || secondErr != nil {
		fmt.Println("Error saving note")
		os.Exit(1)
	}

	fmt.Println("Note saved successfully")

}
func saveData(data saver) error {
	err := data.Save()
	if err != nil {
		return err
	}
	fmt.Println("Data saved successfully")
	return nil
}

// first way
// func printAnyThing(data interface{}) {
// 	fmt.Println(data)
// }

// second way
func printAnyThing(data any) {
	// switch data.(type) {
	// case int:
	// 	fmt.Println("This is an integer")
	// case string:
	// 	fmt.Println("This is a string")
	// case float64:
	// 	fmt.Println("This is a float")
	// }
	// typedValue, ok := data.(int)
	// fmt.Println(typedValue, ok)
	typedValue, ok := data.(string)
	if ok {
		fmt.Println("This is a string", typedValue)
	}
	typedValueInt, ok := data.(int)
	if ok {
		fmt.Println("This is an integer", typedValueInt)
	}
	typedValueFloat, ok := data.(float64)
	if ok {
		fmt.Println("This is a float", typedValueFloat)
	}

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
