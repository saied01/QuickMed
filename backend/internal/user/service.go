package user

import (
	"errors"

	"quickmed/pkg/security"
	"quickmed/pkg/validation"
)

type UserService struct {
	repo *UserRepository
}

func NewUserService(repo *UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (r *UserService) RegisterUser(email string, name string, password string) error {
	err := validation.ValidMail(email)
	if err != nil {
		return err
	}

	if u, _ := r.repo.GetByEmail(email); u != nil {
		return errors.New("email already registered")
	}

	// TODO: password validator

	hashPass, errhash := security.HashPassword(password)
	if errhash != nil {
		return errhash
	}

	u := &User{
		Email:    email,
		Name:     name,
		Password: hashPass,
	}
	err2 := r.repo.Create(u)

	return err2
}

func (r *UserService) GetUserByID(id uint) (*User, error) {
	user, err := r.repo.GetByID(id)
	return user, err
}

func (r *UserService) UpdateUser(id uint, updated *User) error {
	// get user and update changed values
	existing, err := r.repo.GetByID(id)
	if err != nil {
		return err
	}

	if updated.Email != "" && updated.Email != existing.Email {
		return errors.New("email cannot be changed")
	}

	if updated.Name != "" && updated.Name != existing.Name {
		existing.Name = updated.Name
	}

	if updated.Password != "" {
		hash, err := security.HashPassword(updated.Password)
		if err != nil {
			return err
		}
		existing.Password = hash
	}

	return r.repo.Update(existing)
}

func (r *UserService) DeleteUser(user *User) error {
	return r.repo.Delete(user)
}
