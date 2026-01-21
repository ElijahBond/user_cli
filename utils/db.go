package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

type User struct {
	ID   string `json:"ID"`
	User string `json:"user"`
}

func Database() []User {
	database_name := "database.json"
	CreateDB(database_name)
	return ReadDB(database_name)
}

func CreateDB(dbName string) {
	// смотрю есть ли файл, если нет создаю новый
	_, err := os.Open(dbName)
	log.Println("---Open Database---")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			file, _ := os.Create(dbName)
			log.Println("---Database created---")
			file.Close()
		}
	}
}

func ReadDB(dbName string) []User {
	data, err := os.ReadFile(dbName)
	if err != nil {
		fmt.Println(err)
	}

	var obj []User

	err = json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Println(err)
	}

	return obj
}

func SaveDB(usrs []User) {
	savedData, err := json.MarshalIndent(usrs, "", "  ")
	if err != nil {
		fmt.Println(err)
	}

	err = os.WriteFile("database.json", savedData, 0644)

	if err != nil {
		fmt.Println(err)
	}
	log.Println("---DB Saved---")
}
