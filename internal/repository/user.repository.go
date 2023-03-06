package repository

import (
	"context"
	"database/sql"
	"github.com/iqbaleff214/go-fiber-api-aplikasi-pembayaran-spp/internal/entity"
)

type UserRepository interface {
	GetUserByUsername(ctx context.Context, username string) (entity.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return userRepository{db: db}
}

func (u userRepository) GetUserByUsername(ctx context.Context, username string) (entity.User, error) {
	var out entity.User
	query := `SELECT id, name, username, password, role FROM users WHERE username = ?`
	err := u.db.QueryRowContext(ctx, query, username).Scan(&out.ID, &out.Name, &out.Username, &out.Password, &out.Role)

	return out, err
}
