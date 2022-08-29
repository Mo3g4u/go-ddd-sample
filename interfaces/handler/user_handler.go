package handler

import (
	"errors"
	"net/http"

	"github.com/Mo3g4u/go-ddd-sample/application/usecase"
	"github.com/Mo3g4u/go-ddd-sample/domain/model"
	"github.com/Mo3g4u/go-ddd-sample/domain/repository"
	"github.com/Mo3g4u/go-ddd-sample/domain/service"
	"github.com/Mo3g4u/go-ddd-sample/interfaces/request"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	u *usecase.UserUsecase
}

func NewUserHandler(r repository.Repository, us service.UserService) UserHandler {
	return UserHandler{
		u: usecase.NewUserUsecase(r, us),
	}
}

func (u *UserHandler) Create(c *gin.Context) {
	req := request.UserCreateRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := u.u.Create(req); err != nil {
		var reErr *model.UserExistsError
		if errors.As(err, &reErr) {
			// UserExistsError の場合は 409 conflict を返す
			c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
