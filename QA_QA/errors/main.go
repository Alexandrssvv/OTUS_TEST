package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	//err := file()
	//if err != nil {
	//	panic(err.Error())
	//}

	err := httpCall()
	if err != nil {
		panic(err.Error())
	}
}

func httpCall() error {
	resp, err := http.Get("https://google.com")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	fmt.Println(string(body))

	return nil
}

func file() error {
	file, err := os.OpenFile("./errors/text.txt", os.O_APPEND|os.O_RDWR, os.ModeAppend)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString("Hello World\n")
	if err != nil {
		return err
	}

	readFile, err := os.Open("./errors/text.txt")
	if err != nil {
		return err
	}
	defer readFile.Close()
	scanner := bufio.NewScanner(readFile)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		return err
	}

	return nil
}
