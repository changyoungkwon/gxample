package cli

import (
	"github.com/changyoungkwon/gxample/internal/database"
	"github.com/changyoungkwon/gxample/internal/models"
	"github.com/spf13/cobra"
)

// GendocCmd represents the gendoc command
var GendocCmd = &cobra.Command{
	Use:   "gendoc",
	Short: "Generate project documentation",
	Run: func(cmd *cobra.Command, args []string) {
		genRoutesDoc()
	},
}

func genRoutesDoc() {
	db := database.DBConn()
	recipeStore := database.NewRecipeStore(db)
	quantity1 := []byte("{\"Amount\": 1, \"Unit\": 2}")
	quantity2 := []byte("{\"Amount\": \"약간\"}")
	recipeStore.Add(&models.Recipe{
		Title:            "빠라밤밤밥",
		ImagePath:        "/asdf/asdf",
		Ease:             "easy",
		PreparationTime:  15,
		RecipeCategoryID: 1,
		WriterID:         "123-123",
		IngredientQuantities: []*models.IngredientQuantity{
			{IngredientID: 2, Quantity: quantity1},
			{IngredientID: 3, Quantity: quantity2},
		},
		Steps: []models.RecipeStep{
			{Index: 0, Description: "0-이렇게", Tip: "", ImagePath: ""},
			{Index: 1, Description: "1-저렇게", Tip: "", ImagePath: ""},
			{Index: 2, Description: "2-이렇게", Tip: "", ImagePath: ""},
			{Index: 3, Description: "3-요렇게", Tip: "", ImagePath: ""},
		},
		Tags: []*models.Tag{
			{Name: "asdfaht"},
			{Name: "adsasdffwesome"},
		},
	})
}
