package mysql

import (
	"context"
	"database/sql"
	usr "xm/pkg/user"
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
	SQL := "SELECT * FROM users d WHERE d.email = ?"
	row := u.db.QueryRowContext(ctx, SQL, email)
	err := row.Scan(&user.Name, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
