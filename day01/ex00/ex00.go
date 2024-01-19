package main

import (
	"day01/readdb"
	"flag"
	"fmt"
)

// example start
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

type DBReader interface {
	ReadDb() (db readdb.Recipes, err error)
}

func WriteLog(s DBReader) (db readdb.Recipes, err error) {
	return s.ReadDb()
}

// example end

func main() {

	filename := flag.String("f", "", "json or xml file")
	flag.Parse()

	xml := XmlDb{*filename}
	db, err := WriteLog(xml)

	// Инициализируем объект Count и передаём в WriteLog().
	json := JsonDb{*filename}
	db, err = WriteLog(json)
	// fmt.Println(*filename)
	if err != nil {
		fmt.Println(err.Error())

		return
	}
	fmt.Println(db)
	// readdb.PrintRecipes(db, format)
}
