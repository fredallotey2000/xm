package user

import (
	"context"

	"golang.org/x/crypto/bcrypt"
)

type service struct {
	repo Repository
}

type Service interface {
	GetUser(ctx context.Context, email string) (*User, error)
	AuthUser(ctx context.Context, email, password string) (*User, error)
}

// NewService creates a new DNS service with an instance of a repository
func NewUserService(r Repository) Service {
	o := &service{
		repo: r,
	}
	return o
}

// CalculatePath gets the data bank location for a repository
func (s *service) GetUser(ctx context.Context, email string) (*User, error) {

	user, err := s.repo.GetUser(ctx, email)
	if err != nil {
		return nil, err
	}
	return user, nil

}

func (s *service) AuthUser(ctx context.Context, email, password string) (*User, error) {

	user, err := s.repo.GetUser(ctx, email)
	if err != nil {
		return nil, err
	}
	err = checkPassword(user.Password, password)
	if err != nil {
		return nil, err
	}
	return user, nil

}

// func HashPassword(password string) error {
// 	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
// 	if err != nil {
// 		return err
// 	}
// 	user.Password = string(bytes)
// 	return nil
// }

func checkPassword(password, providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
