package user

import "context"

type Repository interface {
	GetUser(ctx context.Context, email string) (*User, error)
}
