package note

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title      string    `json:"title"`
	Content    string    `json:"content"`
	CreateedAt time.Time `json:"created_at"`
}

func New(title string, content string) Note {
	if title == "" || content == "" {
		panic("title and content are required")
	}
	return Note{
		Title:      title,
		Content:    content,
		CreateedAt: time.Now(),
	}
}

// method to print note
func (n Note) Print() string {
	return fmt.Sprintf("Title: %s\nContent: %s\nCreated At: %s", n.Title, n.Content, n.CreateedAt)
}

//func to save note as a .json file

func (note Note) SaveAsJSON() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName)

	json, err := json.MarshalIndent(note, "", " ")

	if err != nil {
		return err
	}

	return os.WriteFile(fileName+".json", json, 0644)
}
