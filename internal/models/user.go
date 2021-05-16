package models

// User to specify the user
type User struct {
	ID             string `gorm:"primaryKey"`
	Name           string `gorm:"unique;not null"`
	Description    string
	ImagePath      string
	MyRecipes      []Recipe  `gorm:"foreignKey:WriterID"`
	ClippedRecipes []*Recipe `gorm:"many2many:recipe_clippers;joinForeignKey:ClipperID"`
	Subscribers    []User    `gorm:"many2many:user_subscribers"`
}
