package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Была паника: ", r)
		}
	}()

	err := something()
	if err != nil {
		var ourErr Error

		if errors.As(err, &ourErr) {
			fmt.Println(ourErr.Text)

		}

		panic(err.Error())
	}
}

func something() error {
	_, err := os.Open("./errors/text.txt")
	if err != nil {
		return Error{
			Text:        "Проблема с файлом",
			Description: err.Error(),
			NumberCode:  10,
		}
	}

	return nil

}

type Error struct {
	Text        string
	Description string
	NumberCode  int
}

func (e Error) Error() string {
	return fmt.Sprintf("Text error: %s; Number code %d; Description: %s",
		e.Text,
		e.NumberCode,
		e.Description)
}
