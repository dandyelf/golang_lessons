package readdb

type DataBase struct {
	Cake []struct {
		Name        string `json:"name"`
		Time        string `json:"time"`
		Ingredients []struct {
			IngredientName  string `json:"ingredient_name"`
			IngredientCount string `json:"ingredient_count"`
			IngredientUnit  string `json:"ingredient_unit,omitempty"`
		} `json:"ingredients"`
	} `json:"cake"`
}

func readDb() {

}
