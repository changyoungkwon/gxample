package cli

import (
	"github.com/changyoungkwon/gxample/internal/database"
	"github.com/changyoungkwon/gxample/internal/logging"
	"github.com/changyoungkwon/gxample/internal/models"
	"github.com/spf13/cobra"
	"gorm.io/gorm"
)

var seedingUsers = []*models.User{
	{
		ID:          "123-123",
		Name:        "이두희",
		Description: "두희입니다",
		ImagePath:   "/static/user/1/profile.png",
	},
	{
		ID:          "1234-1234",
		Name:        "지완규",
		Description: "완규입니다",
		ImagePath:   "/static/user/2/profile.png",
	},
	{
		ID:          "12345-12345",
		Name:        "박상오",
		Description: "상오입니다",
		ImagePath:   "/static/user/3/profile.png",
	},
}

var seedingIngredients = []*models.Ingredient{
	{
		Name:      "닭고기",
		ImagePath: "/static/ingredient/1/thumb.png",
	},
	{
		Name:      "소고기",
		ImagePath: "/static/ingredient/2/thumb.png",
	},
	{
		Name:      "돼지고기",
		ImagePath: "/static/ingredient/1/thumb.png",
	},
}

var seedingRecipeCategories = []*models.RecipeCategory{
	{
		Name:      "양식",
		ImagePath: "/static/category/1/thumb.png",
	},
	{
		Name:      "중식",
		ImagePath: "/static/category/2/thumb.png",
	},
	{
		Name:      "한식",
		ImagePath: "/static/category/1/thumb.png",
	},
}

// cleanDatabase
func cleanDatabase(db *gorm.DB) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		stmts := []string{
			"DELETE FROM INGREDIENT_QUANTITY CASCADE",
			"DELETE FROM RECIPE_TAGS CASCADE",
			"DELETE FROM USERS CASCADE",
			"DELETE FROM INGREDIENTS CASCADE",
			"DELETE FROM RECIPE_CATEGORIES CASCADE",
			"DELETE FROM RECIPES CASCADE",
			"DELETE FROM TAGS CASCADE",
		}
		for _, stmt := range stmts {
			err := tx.Exec(stmt).Error
			if err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

// seedDatabase
func seedDatabase(db *gorm.DB) error {
	err := db.Transaction(func(tx *gorm.DB) error {
		if err := cleanDatabase(db); err != nil {
			return err
		}
		ingredientStore := database.NewIngredientStore(tx)
		recipeCategoryStore := database.NewRecipeCategoryStore(tx)
		userStore := database.NewUserStore(tx)
		for _, i := range seedingIngredients {
			if err := ingredientStore.Add(i); err != nil {
				return err
			}
		}

		for _, r := range seedingRecipeCategories {
			if err := recipeCategoryStore.Add(r); err != nil {
				return err
			}
		}
		for _, u := range seedingUsers {
			if err := userStore.Add(u); err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

// SeedCmd seeds based on seed.yml
var SeedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Seed",
	Run: func(cmd *cobra.Command, args []string) {
		db := database.DBConn()
		if err := seedDatabase(db); err != nil {
			logging.Logger.Errorf("cannot seed on database, %v", err)
			panic(err)
		}
		logging.Logger.Info("seed database succesfully")
	},
}
