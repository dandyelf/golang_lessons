package main

import (
	"day01/readdb"
	"flag"
	"fmt"
)

type DBReader interface {
	ReadDb() (db readdb.Recipes, err error)
}

type XmlDb struct {
	filename string
}

func (b XmlDb) ReadDb() (db readdb.Recipes, err error) {
	return
}

type JsonDb struct {
	filename string
}

func (c JsonDb) ReadDb() (db readdb.Recipes, err error) {
	return
}

func main() {

	filename := flag.String("f", "", "json or xml file")
	flag.Parse()

	format, err := readdb.CheckFormatFile(*filename)

	if err != nil {
		fmt.Println(err.Error())

		return base, err, format
	}

	file, err := ReadFile(file_name)

	xml := XmlDb{*filename}
	json := JsonDb{*filename}

	db, err := WriteLog(xml)
	db, err = WriteLog(json)
	// fmt.Println(*filename)
	if err != nil {
		fmt.Println(err.Error())

		return
	}
	fmt.Println(db)
	// readdb.PrintRecipes(db, format)
}

func WriteLog(s DBReader) (db readdb.Recipes, err error) {
	return s.ReadDb()
}
