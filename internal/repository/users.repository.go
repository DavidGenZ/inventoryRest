package repository

import (
	"context"

	"github.com/DavidGenZ/inventory-project/internal/entity"
)

func (r *repo) SaveUser(ctx context.Context, email, name, password string) error {
	return nil
}

func (r *repo) GetUser(ctx context.Context, email string) (*entity.User, error) {
	return nil, nil
}

