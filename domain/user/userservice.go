package user

import (
	"fmt"

	"github.com/hiroki-Fukumoto/farm2/apierror"
	"github.com/hiroki-Fukumoto/farm2/database"
	"github.com/hiroki-Fukumoto/farm2/database/model"
	"github.com/hiroki-Fukumoto/farm2/logger"
)

type UserService interface {
	RegisterUser(accountID, name string) (*UserResponse, *apierror.APIError)
	FindUser(accountID string) (*UserResponse, *apierror.APIError)
}

type userService struct {
	userRepository UserRepository
}

func NewUserService(r UserRepository) UserService {
	return &userService{
		userRepository: r,
	}
}

func (s *userService) RegisterUser(accountID, name string) (*UserResponse, *apierror.APIError) {
	existUser, err := s.userRepository.FindUserByAccountID(accountID)
	if err != nil {
		logger.Err(err)
		if !database.IsNoRowsError(err) {
			return nil, apierror.NewAPIError(apierror.ErrInternalServerError, apierror.InternalServerErrCode, err.Error())
		}
	}
	if existUser != nil {
		return nil, apierror.NewAPIError(apierror.ErrConflict, apierror.DuplicateDataErrCode, "The specified account ID is not available.")
	}

	req := &model.User{
		AccountID: accountID,
		Name:      name,
	}
	err = s.userRepository.CreateUser(req)
	if err != nil {
		return nil, apierror.NewAPIError(apierror.ErrInternalServerError, apierror.InternalServerErrCode, err.Error())
	}

	res := &UserResponse{
		AccountID: accountID,
		Name:      name,
	}

	return res, nil
}

func (s *userService) FindUser(accountID string) (*UserResponse, *apierror.APIError) {
	user, err := s.userRepository.FindUserByAccountID(accountID)

	if err != nil {
		logger.Err(err)
		if database.IsNoRowsError(err) {
			return nil, apierror.NewAPIError(apierror.ErrNotFound, apierror.NotFoundDataErrCode, fmt.Sprintf("account_id=%s: user not found.", accountID))
		} else {
			return nil, apierror.NewAPIError(apierror.ErrInternalServerError, apierror.InternalServerErrCode, err.Error())
		}
	}

	res := &UserResponse{
		AccountID: user.AccountID,
		Name:      user.Name,
	}

	return res, nil
}
