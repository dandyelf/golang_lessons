package readdb

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
)

type Recipes struct {
	Cake []struct {
		Name        string `json:"name,omitempty"`
		Time        string `json:"time,omitempty"`
		Ingredients []struct {
			IngredientName  string `json:"ingredient_name,omitempty"`
			IngredientCount string `json:"ingredient_count,omitempty"`
			IngredientUnit  string `json:"ingredient_unit,omitempty"`
		} `json:"ingredients,omitempty"`
	} `json:"cake,omitempty"`
}

func (r *Recipes) RecipesRead {

}

func readDb(file_name string) {
	jsonFile, err := readFile(file_name)

	var json_base RecipesJson

	json.Unmarshal(jsonFile, &json_base)
	for i := 0; i < len(json_base.Cake); i++ {
		fmt.Println("User Type: " + json_base.Cake[i].Name)
		fmt.Println("User Age: " + json_base.Cake[i].Time)
		fmt.Println("User Name: " + json_base.Cake[i].Ingredients[i].IngredientName)
		fmt.Println("Facebook Url: " + json_base.Cake[i].Ingredients[i].IngredientUnit)
	}

}

func readFile(file_name string) ([]byte, error) {
	file, err := os.ReadFile(file_name)
	if err != nil {
		return nil, fmt.Errorf("open file error: %w", err)
	}

	return file, nil
}
