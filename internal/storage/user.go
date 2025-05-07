package storage

import (
	"context"
	"errors"
	"fmt"
	"log"

	sq "github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/AndreySirin/avito-backend-assignment-2023/internal/entity"
)

func (s *Storage) CreateUser(ctx context.Context, user entity.User) (uuid.UUID, error) {
	if err := user.Validate(); err != nil {
		return uuid.Nil, fmt.Errorf("%w: %v", ErrNotValid, err)
	}

	query, args, err := sq.Insert("users").
		Columns(
			"id",
			"full_name",
			"gender",
			"date_of_birth",
			"create_at",
			"update_at",
			"delete_at",
		).
		Values(
			user.ID,
			user.FullName,
			user.Gender,
			user.DateOfBirth,
			user.CreatedAt,
			user.UpdatedAt,
			user.DeletedAt,
		).Suffix("RETURNING id").
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return uuid.Nil, fmt.Errorf("build query: %v", err)
	}

	var userID uuid.UUID
	if err = s.db.QueryRowContext(
		ctx,
		query,
		args...,
	).Scan(&userID); err != nil {
		return uuid.Nil, fmt.Errorf("create user: %v", err)
	}

	if userID != user.ID {
		return uuid.Nil, errors.New("user ID mismatch")
	}

	return userID, nil
}

// TODO
// - GetUser
// - ListUsers

func (s *Storage) GetUser(ctx context.Context, user entity.User) (*entity.User, error) {
	// FIXME
	return nil, nil
}

func (s *Storage) ListUsers(ctx context.Context, user entity.User) ([]entity.User, error) {
	// FIXME
	return nil, nil
}

func (s *Storage) UpdateUser(ctx context.Context, user entity.User) (err error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("update user: %v", err)
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	var exists bool
	err = tx.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM users WHERE id_user = $1)", user.ID).
		Scan(&exists)
	if err != nil {
		return fmt.Errorf("update user: %v", err)
	}
	if !exists {
		return fmt.Errorf("user does not exist")
	}
	_, err = tx.ExecContext(ctx,
		`UPDATE users SET full_name=$1,gender=$2,date_of_birth=$3
		 WHERE id_user=$4`,
		user.FullName, user.Gender, user.DateOfBirth, user.ID)
	if err != nil {
		return fmt.Errorf("update user: %v", err)
	}
	return nil
}

func (s *Storage) DeleteUser(ctx context.Context, user entity.User) error {
	rows, err := s.db.ExecContext(ctx, "DELETE FROM users WHERE id_user = $1", user.ID)
	if err != nil {
		log.Printf("delete user: %v", err)
	}
	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return fmt.Errorf("checking rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no user found with id %v", user.ID)
	}
	return nil
}
