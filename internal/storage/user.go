package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

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

func (s *Storage) GetUser(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	var user entity.User
	query, args, err := sq.Select(
		"id",
		"full_name",
		"gender",
		"date_of_birth",
	).From("users").
		Where(sq.And{
			sq.Eq{"id": id},
			sq.Expr("delete_at IS NULL"),
		}).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, fmt.Errorf("build query: %v", err)
	}
	err = s.db.QueryRowContext(ctx, query, args...).
		Scan(&user.ID, &user.FullName, &user.Gender, &user.DateOfBirth)
	if err != nil {
		return nil, fmt.Errorf("get user: %v", err)
	}
	return &user, nil
}

func (s *Storage) ListUsers(ctx context.Context) ([]entity.User, error) {
	query, args, err := sq.Select(
		"id",
		"full_name",
		"gender",
		"date_of_birth",
	).From("users").
		Where(sq.Eq{"delete_at": nil}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("build query: %v", err)
	}
	rows, err := s.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("query: %v", err)
	}

	defer func() {
		if errClose := rows.Close(); errClose != nil {
			s.lg.Error("error closing rows", "error", errClose)
			return
		}
	}()

	var users []entity.User
	for rows.Next() {
		var u entity.User
		err = rows.Scan(&u.ID, &u.FullName, &u.Gender, &u.DateOfBirth)
		if err != nil {
			return nil, fmt.Errorf("query rows: %v", err)
		}
		users = append(users, u)
	}
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %v", err)
	}
	return users, nil
}

func (s *Storage) UpdateUser(ctx context.Context, user entity.User) (err error) {
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("error for create tx: %v", err)
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	var exists bool
	err = tx.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)", user.ID).
		Scan(&exists)
	if err != nil {
		return fmt.Errorf("error verifying the user's existence: %v", err)
	}
	if !exists {
		return fmt.Errorf("user does not exist")
	}

	query, args, err := sq.Update("users").
		Set("full_name", user.FullName).
		Set("gender", user.Gender).
		Set("date_of_birth", user.DateOfBirth).
		Set("update_at", user.UpdatedAt).
		Where(sq.And{
			sq.Eq{"id": user.ID},
			sq.Expr("delete_at IS NULL"),
		}).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return fmt.Errorf("build query: %v", err)
	}
	_, err = tx.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("update user: %v", err)
	}
	return nil
}

func (s *Storage) DeleteUser(ctx context.Context, id uuid.UUID) error {
	t := time.Now()
	query, args, err := sq.Update("users").
		Set("delete_at", t).
		Where(sq.And{
			sq.Eq{"id": id},
			sq.Expr("delete_at IS NULL"),
		}).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return fmt.Errorf("build query: %v", err)
	}
	rows, err := s.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("delete user: %v", err)
	}
	rowsAffected, err := rows.RowsAffected()
	if err != nil {
		return fmt.Errorf("checking rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no user found with id %v", id)
	}
	return nil
}

func (s *Storage) CheckExistUser(ctx context.Context, tx *sql.Tx, subs *entity.Subscription) error {
	var exists bool
	err := tx.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)", subs.IDUser).
		Scan(&exists)
	if err != nil {
		return fmt.Errorf("error checking if user exists: %w", err)
	}
	if !exists {
		return fmt.Errorf("user does not exist")
	}
	return nil
}
