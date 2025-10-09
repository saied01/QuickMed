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

func (r *UserService) RegisterUser(email string, name string, password string) (*User, error) {
	err := validation.ValidMail(email)
	if err != nil {
		return nil, err
	}

	if u, _ := r.repo.GetByEmail(email); u != nil {
		return nil, errors.New("email already registered")
	}

	// TODO: password validator

	hashPass, errhash := security.HashPassword(password)
	if errhash != nil {
		return nil, errhash
	}

	u := &User{
		Email:    email,
		Name:     name,
		Password: hashPass,
	}
	err2 := r.repo.Create(u)

	return u, err2
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

func (s *UserService) DeleteUser(id uint) error {
	user, err := s.repo.GetByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(user)
}
