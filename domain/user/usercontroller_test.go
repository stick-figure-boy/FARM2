package user_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/hiroki-Fukumoto/farm2/apierror"
	"github.com/hiroki-Fukumoto/farm2/domain/user"
	"github.com/stretchr/testify/assert"
)

type userServiceMock struct {
	user.UserService
	FakeRegisterUser func(accountID, name string) (*user.UserResponse, *apierror.APIError)
	FakeFindUser     func(accountID string) (*user.UserResponse, *apierror.APIError)
}

func (s *userServiceMock) RegisterUser(accountID, name string) (*user.UserResponse, *apierror.APIError) {
	return s.FakeRegisterUser(accountID, name)
}

func (s *userServiceMock) FindUser(accountID string) (*user.UserResponse, *apierror.APIError) {
	return s.FakeFindUser(accountID)
}

func TestRegisterUserController(t *testing.T) {
	tests := []struct {
		name             string
		service          *userServiceMock
		request          user.RegisterUserRequest
		expectResponse   *user.UserResponse
		expectErrCode    string
		expectStatusCode int
	}{
		{
			name: "Should confirm success.",
			request: user.RegisterUserRequest{
				AccountID: "accountID1",
				Name:      "user1",
			},
			service: &userServiceMock{
				FakeRegisterUser: func(accountID, name string) (*user.UserResponse, *apierror.APIError) {
					return &user.UserResponse{
						AccountID: "accountID1",
						Name:      "user1",
					}, nil
				},
			},
			expectResponse: &user.UserResponse{
				AccountID: "accountID1",
				Name:      "user1",
			},
			expectStatusCode: http.StatusCreated,
		},
		{
			name: "Should confirm AccountID validation error.",
			request: user.RegisterUserRequest{
				AccountID: "",
				Name:      "user1",
			},
			service: &userServiceMock{
				FakeRegisterUser: func(accountID, name string) (*user.UserResponse, *apierror.APIError) {
					return nil, nil
				},
			},
			expectStatusCode: http.StatusBadRequest,
			expectErrCode:    apierror.ValidationErrCode,
		},
		{
			name: "Should confirm Name validation error.",
			request: user.RegisterUserRequest{
				AccountID: "accountID1",
				Name:      "",
			},
			service: &userServiceMock{
				FakeRegisterUser: func(accountID, name string) (*user.UserResponse, *apierror.APIError) {
					return nil, nil
				},
			},
			expectStatusCode: http.StatusBadRequest,
			expectErrCode:    apierror.ValidationErrCode,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, _ := json.Marshal(tt.request)
			resp := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(resp)
			ctx.Request, _ = http.NewRequest(
				http.MethodPost,
				"/users",
				bytes.NewBuffer(b),
			)

			cnt := user.NewUserController(tt.service)
			cnt.RegisterUser(ctx)

			assert.Equal(t, tt.expectStatusCode, resp.Code)

			if tt.expectResponse != nil {
				var res *user.UserResponse
				err := json.Unmarshal(resp.Body.Bytes(), &res)
				if err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, tt.expectResponse, res)
				return
			}

			var eres *apierror.ErrorResponse
			err := json.Unmarshal(resp.Body.Bytes(), &eres)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.expectErrCode, eres.ErrorCode)
		})
	}
}

func TestFindUserController(t *testing.T) {
	tests := []struct {
		name             string
		accountID        string
		service          *userServiceMock
		expectResponse   *user.UserResponse
		expectErrCode    string
		expectStatusCode int
	}{
		{
			name:      "Should confirm success.",
			accountID: "accountID1",
			service: &userServiceMock{
				FakeFindUser: func(accountID string) (*user.UserResponse, *apierror.APIError) {
					return &user.UserResponse{
						AccountID: "accountID1",
						Name:      "user1",
					}, nil
				},
			},
			expectResponse: &user.UserResponse{
				AccountID: "accountID1",
				Name:      "user1",
			},
			expectStatusCode: http.StatusOK,
		},
		{
			name:      "Should confirm not set accountID.",
			accountID: "",
			service: &userServiceMock{
				FakeFindUser: func(accountID string) (*user.UserResponse, *apierror.APIError) {
					return nil, nil
				},
			},
			expectStatusCode: http.StatusBadRequest,
			expectErrCode:    apierror.ValidationErrCode,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(resp)
			ctx.Request, _ = http.NewRequest(
				http.MethodGet,
				"/users/accountID",
				nil,
			)

			ctx.AddParam("accountID", tt.accountID)
			cnt := user.NewUserController(tt.service)
			cnt.FindUser(ctx)

			assert.Equal(t, tt.expectStatusCode, resp.Code)

			if tt.expectResponse != nil {
				var res *user.UserResponse
				err := json.Unmarshal(resp.Body.Bytes(), &res)
				if err != nil {
					t.Fatal(err)
				}
				assert.Equal(t, tt.expectResponse, res)
				return
			}

			var eres *apierror.ErrorResponse
			err := json.Unmarshal(resp.Body.Bytes(), &eres)
			if err != nil {
				t.Fatal(err)
			}
			assert.Equal(t, tt.expectErrCode, eres.ErrorCode)
		})
	}
}
