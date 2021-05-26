package database

import (
	"github.com/changyoungkwon/gxample/internal/models"
	"gorm.io/gorm"
)

// UserStore is the wrapper class for the database
type UserStore struct {
	db *gorm.DB
}

// NewUserStore provide store abstraction
func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

// Add create user, update i with automatically received value, then return err
func (s *UserStore) Add(i *models.User) error {
	result := s.db.Create(&i)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Get get user by key
func (s *UserStore) Get(key string) (*models.User, error) {
	i := &models.User{}
	if result := s.db.First(&i, key); result.Error != nil {
		return nil, result.Error
	}
	return i, nil
}

// List list all users
func (s *UserStore) List() ([]models.User, error) {
	var users []models.User
	result := s.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
