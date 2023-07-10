package user_test

import (
	"database/sql"
	"errors"
	"net/http"
	"testing"

	"github.com/hiroki-Fukumoto/farm2/apierror"
	"github.com/hiroki-Fukumoto/farm2/database/model"
	"github.com/hiroki-Fukumoto/farm2/domain/user"
	"github.com/stretchr/testify/assert"
)

type userRepositoryMock struct {
	user.UserRepository
	FakeCreateUser          func(user *model.User) error
	FakeFindUserByAccountID func(accountID string) (*model.User, error)
}

func (r *userRepositoryMock) CreateUser(user *model.User) error {
	return r.FakeCreateUser(user)
}

func (r *userRepositoryMock) FindUserByAccountID(accountID string) (*model.User, error) {
	return r.FakeFindUserByAccountID(accountID)
}

func TestRegisterUserService(t *testing.T) {
	tests := []struct {
		name           string
		accountID      string
		userName       string
		repository     *userRepositoryMock
		expectResponse *user.UserResponse
		expectError    *apierror.APIError
	}{
		{
			name:      "Should confirm created user.",
			accountID: "accountID1",
			userName:  "user1",
			repository: &userRepositoryMock{
				FakeCreateUser: func(user *model.User) error {
					return nil
				},
				FakeFindUserByAccountID: func(accountID string) (*model.User, error) {
					return nil, sql.ErrNoRows
				},
			},
			expectResponse: &user.UserResponse{
				AccountID: "accountID1",
				Name:      "user1",
			},
		},
		{
			name:      "Should confirm error create user.",
			accountID: "accountID1",
			userName:  "user1",
			repository: &userRepositoryMock{
				FakeCreateUser: func(user *model.User) error {
					return errors.New("sql error.")
				},
				FakeFindUserByAccountID: func(accountID string) (*model.User, error) {
					return nil, sql.ErrNoRows
				},
			},
			expectError: &apierror.APIError{
				StatusCode: http.StatusInternalServerError,
				ErrorCode:  apierror.InternalServerErrCode,
				Error:      apierror.ErrInternalServerError.Error(),
				Message:    "sql error.",
			},
		},
		{
			name:      "Should confirm already used accountID.",
			accountID: "accountID1",
			userName:  "user1",
			repository: &userRepositoryMock{
				FakeCreateUser: func(user *model.User) error {
					return nil
				},
				FakeFindUserByAccountID: func(accountID string) (*model.User, error) {
					return &model.User{}, sql.ErrNoRows
				},
			},
			expectError: &apierror.APIError{
				StatusCode: http.StatusConflict,
				ErrorCode:  apierror.DuplicateDataErrCode,
				Error:      apierror.ErrConflict.Error(),
				Message:    "The specified account ID is not available.",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := user.NewUserService(tt.repository)

			res, err := service.RegisterUser(tt.accountID, tt.userName)
			if err != nil {
				assert.Equal(t, tt.expectError, err)
				return
			}
			assert.Equal(t, tt.expectResponse, res)
		})
	}
}

func TestFindUserService(t *testing.T) {
	tests := []struct {
		name           string
		accountID      string
		repository     *userRepositoryMock
		expectResponse *user.UserResponse
		expectError    *apierror.APIError
	}{
		{
			name:      "Should confirm success.",
			accountID: "accountID1",
			repository: &userRepositoryMock{
				FakeFindUserByAccountID: func(accountID string) (*model.User, error) {
					return &model.User{
						AccountID: "accountID1",
						Name:      "user1",
					}, nil
				},
			},
			expectResponse: &user.UserResponse{
				AccountID: "accountID1",
				Name:      "user1",
			},
		},
		{
			name:      "Should confirm not found user.",
			accountID: "accountID1",
			repository: &userRepositoryMock{
				FakeFindUserByAccountID: func(accountID string) (*model.User, error) {
					return nil, sql.ErrNoRows
				},
			},
			expectError: apierror.NewAPIError(apierror.ErrNotFound, apierror.NotFoundDataErrCode, "account_id=accountID1: user not found."),
		},
		{
			name:      "Should confirm sql error.",
			accountID: "accountID1",
			repository: &userRepositoryMock{
				FakeFindUserByAccountID: func(accountID string) (*model.User, error) {
					return nil, errors.New("sql error")
				},
			},
			expectError: apierror.NewAPIError(apierror.ErrInternalServerError, apierror.InternalServerErrCode, "sql error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := user.NewUserService(tt.repository)

			res, err := service.FindUser(tt.accountID)
			if err != nil {
				assert.Equal(t, tt.expectError, err)
				return
			}
			assert.Equal(t, tt.expectResponse, res)
		})
	}
}
