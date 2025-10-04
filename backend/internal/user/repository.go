package user

import (
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *User) error {
	result := r.db.Create(user)
	return result.Error
}

func (r *UserRepository) GetByID(id uint) (*User, error) {
	var user User

	result := r.db.First(&user, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

func (r *UserRepository) Update(user *User) error {
	result := r.db.Save(user)
	return result.Error
}

func (r *UserRepository) Delete(user *User) error {
	result := r.db.Delete(user)
	return result.Error
}

func (r *UserRepository) GetByEmail(email string) (*User, error) {
	var user User
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
