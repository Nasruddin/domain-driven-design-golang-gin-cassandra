package http

import (
	"github.com/gin-gonic/gin"
	"golang-gin-cassandra/src/domain/users/model"
	"golang-gin-cassandra/src/domain/users/service"
	"golang-gin-cassandra/src/utils/errors"
	"net/http"
	"strings"
)

type UserHandler interface {
	GetById(ctx *gin.Context)
	Create(ctx *gin.Context)
}

type userHandler struct {
	userService service.UserService
}

func (userHandler userHandler) GetById(ctx *gin.Context) {
	userID := strings.TrimSpace(ctx.Param("user_id"))
	user, err := userHandler.userService.GetByID(userID)
	if err != nil {
		ctx.JSON(err.ErrStatus, err)
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func NewHandler(userService service.UserService) UserHandler {
	return &userHandler{
		userService: userService,
	}
}


func (userHandler *userHandler) Create(ctx *gin.Context)  {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		ctx.JSON(restErr.ErrStatus, restErr)
	}
	if _, err := userHandler.userService.Create(user); err != nil {
		ctx.JSON(err.ErrStatus, err)
		return
	}

	ctx.JSON(http.StatusCreated, user)
}
