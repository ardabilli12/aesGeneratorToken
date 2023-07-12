package repo

import (
	"log"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type UserRepositoryContract interface {
	GetDetailUserByEmail(email string) Users
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) GetDetailUserByEmail(email string) Users {
	var user Users
	err := u.db.Where("email = ?", email).First(&user).Error
	if err != nil {
		log.Fatalf("Error querying database: %v", err)
	}

	return user
}

// Model

type Users struct {
	Uuid      string `json:"uuid"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Photo     string `json:"photo"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	CompanyId string `json:"company_id"`
}
