package user

type RegisterUserRequest struct {
	AccountID string `json:"account_id" validate:"required"`
	Name      string `json:"name" validate:"required"`
}
