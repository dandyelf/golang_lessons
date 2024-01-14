package readdb

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
)

type Recipes struct {
	Cake []struct {
		Name        string `xml:"name" json:"name"`
		Time        string `xml:"stovetime" json:"time"`
		Ingredients []struct {
			IngredientName  string `xml:"itemname" json:"ingredient_name"`
			IngredientCount string `xml:"itemcount" json:"ingredient_count"`
			IngredientUnit  string `xml:"itemunit" json:"ingredient_unit,omitempty"`
		} `xml:"ingredients>item" json:"ingredients"`
	} `xml:"cake" json:"cake"`
}

func ParsFile(file_name string, format string) (Recipes, error) {
	file, err := ReadFile(file_name)
	var base Recipes
	if err != nil {

		return base, fmt.Errorf("open file error: %w", err)
	}
	if format == "json" {
		err = json.Unmarshal(file, &base)
	} else if format == "xml" {
		err = xml.Unmarshal(file, &base)
	} else {
		return base, fmt.Errorf("unknown format: %w", err)
	}
	if err != nil {

		return base, fmt.Errorf("pars file error: %w", err)
	}

	return base, nil
}

func ReadFile(file_name string) ([]byte, error) {
	file, err := os.ReadFile(file_name)
	if err != nil {

		return nil, fmt.Errorf("open file error: %w", err)
	}

	return file, nil
}
