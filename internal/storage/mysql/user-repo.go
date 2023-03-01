package mysql

import (
	usr "xm/pkg/user"
	"context"
	"database/sql"
)

type user struct {
	db *sql.DB
}

func NewUserRepo(d *sql.DB) usr.Repository {
	return &user{
		db: d,
	}
}

//Get user info
func (u *user) GetUser(ctx context.Context, email string) (*usr.User, error) {
	var user = usr.User{}
	SQL := "SELECT * FROM deployments d WHERE d.deploymentId = ?"
	row := u.db.QueryRowContext(ctx, SQL, email)
	err := row.Scan(&user.Name, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *user) getUserObject() *user {
	return u
}
