package localfs

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/a-reshatniak-itechart/projectlayout/internal/app"
)

func NewUserRepository(dir string) UserRepository {
	return UserRepository{dir: dir}
}

type UserRepository struct {
	dir string
}

func (r UserRepository) GetByID(ctx context.Context, id int) (app.User, error) {
	var u app.User
	filepath := fmt.Sprintf("%s/%d.json", r.dir, id)
	data, err := os.ReadFile(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			return u, app.ErrUserNotFound
		}

		return u, err
	}

	if err = json.Unmarshal(data, &u); err != nil {
		return u, err
	}

	return u, nil
}
