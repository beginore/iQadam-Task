package repository

import (
	"context"
	"database/sql"
	"ums/internal/domain/model"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
	GetByID(ctx context.Context, id int64) (*model.User, error)
	GetByUsername(ctx context.Context, username string) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id int64) error
	GetAll(ctx context.Context) ([]*model.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(ctx context.Context, user *model.User) error {
	query := `INSERT INTO users (username, password, full_name, email, role) 
		VALUES ($1, $2, $3, $4, $5) 
		RETURNING id, created_at`
	return r.db.QueryRowContext(ctx, query,
		user.Username,
		user.Password,
		user.FullName,
		user.Email,
		user.Role,
	).Scan(&user.ID, &user.CreatedAt)
}

func (r *userRepository) GetByID(ctx context.Context, id int64) (*model.User, error) {
	query := `SELECT id, username, password, full_name, email, role, created_at 
             FROM users WHERE id = $1`
	user := &model.User{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.FullName,
		&user.Email,
		&user.Role,
		&user.CreatedAt,
	)
	return user, err
}

func (r *userRepository) GetByUsername(ctx context.Context, username string) (*model.User, error) {
	query := `SELECT id, username, password, full_name, email, role, created_at 
		FROM users WHERE username = $1`
	user := &model.User{}
	err := r.db.QueryRowContext(ctx, query, username).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.FullName,
		&user.Email,
		&user.Role,
		&user.CreatedAt,
	)
	return user, err
}

// Исправленные методы с правильным ресивером
func (r *userRepository) Update(ctx context.Context, user *model.User) error {
	query := `UPDATE users 
		SET username = $1, full_name = $2, email = $3, role = $4 
		WHERE id = $5`
	_, err := r.db.ExecContext(ctx, query,
		user.Username,
		user.FullName,
		user.Email,
		user.Role,
		user.ID,
	)
	return err
}

func (r *userRepository) Delete(ctx context.Context, id int64) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *userRepository) GetAll(ctx context.Context) ([]*model.User, error) {
	query := `SELECT id, username, full_name, email, role, created_at FROM users`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*model.User
	for rows.Next() {
		var user model.User
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.FullName,
			&user.Email,
			&user.Role,
			&user.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}
