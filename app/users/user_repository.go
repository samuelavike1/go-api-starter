package users

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	DB *pgx.Conn
}

func NewUserRepository(db *pgx.Conn) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetUserByID(ctx context.Context, id int) (*User, error) {
	user := &User{}
	err := r.DB.QueryRow(ctx, "SELECT id, name, email FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return user, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, user User) (int, error) {
	var id int

	err := r.DB.QueryRow(ctx, "INSERT INTO users (name, email) VALUES ($1, $2) returning id", user.Name, user.Email).Scan(&id)

	if err != nil {
		return 0, fmt.Errorf("failed to create user: %w", err)
	}

	return id, nil
}

func (r *UserRepository) GetAllUsers(ctx context.Context) ([]User, error) {
	rows, err := r.DB.Query(ctx, "SELECT id, name, email FROM users")
	if err != nil {
		return nil, fmt.Errorf("failed to get all users: %w", err)
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email); err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during rows iteration %w", err)
	}

	return users, nil
}
