package todo

import (
	"encoding/json"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"text"`
}

func New(text string) Todo {
	if text == "" {
		panic("title and content are required")
	}
	return Todo{
		Text: text,
	}
}

// method to print note
func (todo Todo) Print() {
	fmt.Println(todo.Text)
}

//func to save note as a .json file

func (todo Todo) Save() error {
	fileName := "todo.json"

	json, err := json.MarshalIndent(todo, "", " ")

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}
