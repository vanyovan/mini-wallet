package repo

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/vanyovan/mini-wallet.git/internal/entity"
)

type Repo struct {
	db *sql.DB
}

type UserRepo interface {
	CreateUser(ctx context.Context, userId string) (token string, err error)
	GetUserByUserId(userId string) (result entity.User, err error)
	GetUserByToken(ctx context.Context, token string) (result entity.User, err error)
}

func NewUserRepo(db *sql.DB) UserRepo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) CreateUser(ctx context.Context, userId string) (token string, err error) {
	tx, err := r.db.Begin()
	if err != nil {
		tx.Rollback()
		return "", errors.New("failed to begin database transaction")
	}

	//generate token
	token = generateToken()
	_, err = tx.ExecContext(ctx, "INSERT INTO mst_user (user_id, token) VALUES (?, ?)", userId, token)
	if err != nil {
		tx.Rollback()
		return "", fmt.Errorf("failed to create user: %w", err)
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return "", errors.New("failed to commit database transaction")
	}
	return token, nil
}

func (r *Repo) GetUserByUserId(userId string) (result entity.User, err error) {
	query := "SELECT user_id, token FROM mst_user WHERE user_id = ?"
	row := r.db.QueryRow(query, userId)
	result = entity.User{}
	err = row.Scan(&result.CustomerXid, &result.Token)
	if err != nil {
		if err == sql.ErrNoRows {
			return result, nil
		} else {
			fmt.Println("Failed to retrieve row:", err)
		}
		return result, err
	}
	return result, nil
}

func (r *Repo) GetUserByToken(ctx context.Context, token string) (result entity.User, err error) {
	query := "SELECT user_id, token FROM mst_user WHERE token = ?"
	row := r.db.QueryRow(query, token)
	result = entity.User{}
	err = row.Scan(&result.CustomerXid, &result.Token)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No rows found.")
		} else {
			fmt.Println("Failed to retrieve row:", err)
		}
		return result, err
	}
	return result, nil
}
func generateToken() string {
	token := make([]byte, 16)
	if _, err := rand.Read(token); err != nil {
		return ""
	}
	tokenString := hex.EncodeToString(token)
	return tokenString
}
