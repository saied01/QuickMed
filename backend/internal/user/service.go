package user

import (
	"errors"
	"net/mail"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo *UserRepository
}

func NewUserService(repo *UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (r *UserService) RegisterUser(email string, name string, password string) error {
	err := valid_mail(email)
	if err != nil {
		return err
	}

	if u, _ := r.repo.GetByEmail(email); u != nil {
		return errors.New("Email already registered.")
	}

	// TODO: password validator

	hash_pass, errhash := HashPassword(password)
	if errhash != nil {
		return errhash
	}

	u := &User{
		Email:    email,
		Name:     name,
		Password: hash_pass,
	}
	err2 := r.repo.Create(u)

	return err2
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func valid_mail(email string) error {
	_, err := mail.ParseAddress(email)
	return err
}

func (r *UserService) GetUserByID(id uint) (*User, error) {
	user, err := r.repo.GetByID(id)
	return user, err
}

func (r *UserService) UpdateUser(user *User) error {
	return r.repo.Update(user)
}

func (r *UserService) DeleteUser(user *User) error {
	return r.repo.Delete(user)
}
