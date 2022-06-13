package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	SaveAvatar(id int, fileLocation string) (User, error)
	GetUserById(id int) (User, error)
	GetAllUsers() ([]User, error)
}

type service struct {
	repository Repository
}

// User service instance
func NewService(repository Repository) *service {
	return &service{repository}
}

// Service to register a new user
func (s *service) RegisterUser(input RegisterUserInput) (User, error) {
	// mapping input struct to User struct
	user := User{}
	user.Name = input.Name
	user.Occupation = input.Occupation
	user.Email = input.Email
	user.Role = "user"

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.PasswordHash = string(passwordHash)

	// save User struct via repository
	newUser, err := s.repository.Create(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

// Service to login a user
func (s *service) Login(input LoginInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User not found with that email")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

// Service to check if email is available
func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	email := input.Email

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

// Service to save avatar
func (s *service) SaveAvatar(id int, fileLocation string) (User, error) {
	// get user by id
	user, err := s.repository.FindById(id)
	if err != nil {
		return user, err
	}

	// update user avatar_file_name
	user.AvatarFileName = fileLocation

	// save avatar changes to db
	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

// Service to get user by id
func (s *service) GetUserById(id int) (User, error) {
	user, err := s.repository.FindById(id)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User not found with that ID")
	}

	return user, nil
}

// Service to get all users
func (s *service) GetAllUsers() ([]User, error) {
	users, err := s.repository.FindAll()
	if err != nil {
		return users, err
	}

	return users, nil
}
