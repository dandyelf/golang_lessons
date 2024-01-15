package main

import (
	"day01/readdb"
	"flag"
	"fmt"
	"path/filepath"
)

func main() {
	filename := flag.String("f", "", "json or xml file")
	flag.Parse()
	fmt.Println(*filename)

	db, err := readdb.ParsFile(*filename, format)
	if err != nil {
		fmt.Println(err.Error())

		return
	}
	readdb.PrintRecipes(db, format)
}
