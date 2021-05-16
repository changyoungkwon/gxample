package models

import (
	"time"
)

// RecipeCategory application struct
type RecipeCategory struct {
	ID            int       `json:"-"`
	Name          string    `json:"name" pg:"unique,notnull"`
	ThumbnailPath string    `json:"thumbnail_path"`
	UpdatedAt     time.Time `json:"updated_at"`
}
