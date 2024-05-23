package sqlite

import (
	"context"
	"database/sql"
	"errors"

	"github.com/a-reshatniak-itechart/projectlayout/internal/app"
)

func NewUserRepository(conn *sql.DB) UserRepository {
	return UserRepository{conn: conn}
}

type UserRepository struct {
	conn *sql.DB
}

func (r UserRepository) GetByID(ctx context.Context, id int) (app.User, error) {
	var u app.User
	q := `select id, name from users where id = ?`
	err := r.conn.QueryRowContext(ctx, q, id).Scan(&u.ID, &u.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return u, app.ErrUserNotFound
		}

		return u, err
	}

	return u, nil
}
