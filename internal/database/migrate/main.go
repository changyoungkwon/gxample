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
			ID: "202105162056",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&models.Ingredient{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable("ingredients")
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
