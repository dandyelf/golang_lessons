package main

import (
	"day01/readdb"
	"flag"
	"fmt"
	"path/filepath"
)

var fileName string

func main() {
	filename := flag.String("f", "", "json or xml file")
	flag.Parse()
	fmt.Println(*filename)
	format, err := checkFormatFile(*filename)
	if err != nil {
		fmt.Println(err.Error())

		return
	}
	db, err := readdb.ParsFile(*filename, format)
	if err != nil {
		fmt.Println(err.Error())

		return
	}
	fmt.Println(db)
}

func checkFormatFile(file_name string) (string, error) {
	if file_name == "" {

		return "", fmt.Errorf("ошибка: отсутствует файл")
	}
	if filepath.Ext(file_name) == ".json" {

		return "json", nil
	}
	if filepath.Ext(file_name) == ".xml" {

		return "xml", nil
	}

	return "", fmt.Errorf("ошибка: файл должен быть json или xml")
}
