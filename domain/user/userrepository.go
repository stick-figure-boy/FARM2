package user

import (
	"github.com/hiroki-Fukumoto/farm2/database/model"
	"github.com/hiroki-Fukumoto/farm2/logger"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	CreateUser(user *model.User) error
	FindUserByAccountID(accountID string) (*model.User, error)
}

type userRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) UserRepository {
	return &userRepository{DB: db}
}

func (r *userRepository) CreateUser(user *model.User) error {
	query := "INSERT INTO users (account_id, name) VALUES (:account_id, :name);"
	_, err := r.DB.NamedExec(query, user)

	if err != nil {
		logger.Err(err)
		return err
	}

	return nil
}

func (r *userRepository) FindUserByAccountID(accountID string) (*model.User, error) {
	var user model.User
	query := "SELECT account_id, name FROM users WHERE account_id = ? AND deleted_at IS NULL"
	err := r.DB.Get(&user, query, accountID)

	if err != nil {
		logger.Err(err)
		return nil, err
	}

	return &user, nil
}
