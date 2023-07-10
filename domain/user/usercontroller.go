package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hiroki-Fukumoto/farm2/apierror"
	"github.com/hiroki-Fukumoto/farm2/validator"
)

type UserController interface {
	RegisterUser(ctx *gin.Context)
	FindUser(ctx *gin.Context)
}

type userController struct {
	userService UserService
}

func NewUserController(s UserService) UserController {
	return &userController{
		userService: s,
	}
}

// @Summary Register new user.
// @Description Register new user.
// @Tags user
// @Accept json
// @Produce json
// @Param request body RegisterUserRequest true "new user info"
// @Success 201 {object} UserResponse
// @Failure 400 {object} apierror.ErrorResponse
// @Failure 500 {object} apierror.ErrorResponse
// @Router /v1/users [post]
func (c *userController) RegisterUser(ctx *gin.Context) {
	var request RegisterUserRequest
	ctx.BindJSON(&request)
	if err := validator.Validate(&request); err != nil {
		e := apierror.NewAPIError(apierror.ErrBadRequest, apierror.ValidationErrCode, err.Error())
		e.ResponseParser(ctx)
		return
	}

	res, err := c.userService.RegisterUser(request.AccountID, request.Name)
	if err != nil {
		err.ResponseParser(ctx)
		return
	}

	ctx.JSON(http.StatusCreated, res)
}

// @Summary Get user info.
// @Description Get user info matching the specified account ID.
// @Tags user
// @Accept json
// @Produce json
// @Param accountID path string true "account ID"
// @Success 200 {object} UserResponse
// @Failure 404 {object} apierror.ErrorResponse
// @Failure 500 {object} apierror.ErrorResponse
// @Router /v1/users/{accountID} [get]
func (c *userController) FindUser(ctx *gin.Context) {
	accountID := ctx.Param("accountID")
	if accountID == "" {
		err := apierror.NewAPIError(apierror.ErrBadRequest, apierror.ValidationErrCode, "Account ID not specified.")
		err.ResponseParser(ctx)
		return
	}

	user, err := c.userService.FindUser(accountID)
	if err != nil {
		err.ResponseParser(ctx)
		return
	}

	ctx.JSON(http.StatusOK, user)
}
