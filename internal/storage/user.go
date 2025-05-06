package storage

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/AndreySirin/avito-backend-assignment-2023/internal/entity"
	"github.com/AndreySirin/avito-backend-assignment-2023/internal/logger"
	"log"
)

type UserStorage struct {
	lg *logger.MyLogger
	db *sql.DB
}

func NewUser(db *Storage) *UserStorage {
	return &UserStorage{
		lg: db.lg,
		db: db.db,
	}
}
func (u *UserStorage) CreateUser(ctx context.Context, user entity.User) (int, error) {
	err := u.db.QueryRowContext(ctx, "INSERT INTO users (full_name,gender,date_of_birth) VALUES ($1,$2,$3)RETURNING id_user ",
		user.FullName, user.Gender, user.DateOfBirth).Scan(&user.ID)
	if err != nil {
		return 0, fmt.Errorf("create user: %v", err)
	}
	return user.ID, nil
}

func (u *UserStorage) UpdateUser(ctx context.Context, user entity.User) (err error) {
	tx, err := u.db.BeginTx(ctx, nil)
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
	err = tx.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM users WHERE id_user = $1)", user.ID).Scan(&exists)
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

func (u *UserStorage) DeleteUser(ctx context.Context, user entity.User) error {
	rows, err := u.db.ExecContext(ctx, "DELETE FROM users WHERE id_user = $1", user.ID)
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
