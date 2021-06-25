// Package migrate implements postgres migrations
package migrate

import (
	"github.com/changyoungkwon/gxample/internal/database"
	"github.com/changyoungkwon/gxample/internal/models"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func getMigrations() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		{
			ID: "202105162358",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.AutoMigrate(&models.Tag{}); err != nil {
					return err
				}
				if err := tx.AutoMigrate(&models.RecipeCategory{}); err != nil {
					return err
				}
				if err := tx.AutoMigrate(&models.Ingredient{}); err != nil {
					return err
				}
				if err := tx.AutoMigrate(&models.User{}); err != nil {
					return err
				}
				if err := tx.AutoMigrate(&models.Recipe{}); err != nil {
					return err
				}
				if err := tx.AutoMigrate(&models.IngredientQuantity{}); err != nil {
					return err
				}
				tx.Debug().Create(&models.User{
					ID:          "0",
					Name:        "관리자",
					Description: "관리자",
				})
				return nil
			},
			Rollback: func(tx *gorm.DB) error {
				if err := tx.Migrator().DropTable("recipes"); err != nil {
					return err
				}
				if err := tx.Migrator().DropTable("recipe_categories"); err != nil {
					return err
				}
				if err := tx.Migrator().DropTable("ingredients"); err != nil {
					return err
				}
				return nil
			},
		},
	}
}

// Migrate run gormigrate migrations
func Migrate() error {
	db := database.DBConn()
	m := gormigrate.New(db, gormigrate.DefaultOptions, getMigrations())
	return m.Migrate()
}

// RollbackLast run gormigrate undo
func RollbackLast() error {
	db := database.DBConn()
	m := gormigrate.New(db, gormigrate.DefaultOptions, getMigrations())
	return m.RollbackLast()
}
