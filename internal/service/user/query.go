package services

import (
	"grocery-management/internal/models"
	"grocery-management/pkg/logger"
	"grocery-management/pkg/security"
	"grocery-management/pkg/wrapper"

	"github.com/gin-gonic/gin"
)

func (s *UserService) GetUserByUsername(ctx *gin.Context, username string) (*models.User, *wrapper.ErrorWrapper) {
	if username == "" {
		return nil, &wrapper.ErrorWrapper{
			Code:   wrapper.A_BAD_REQUEST.Code,
			Msg:    wrapper.A_BAD_REQUEST.Msg,
			Detail: wrapper.StringWrapper("username cannot be empty"),
		}
	}

	user, err := s.repo.GetByUsername(ctx, username)
	if err != nil {
		return nil, &wrapper.ErrorWrapper{
			Code: wrapper.A_INTERNAL.Code,
			Msg:  wrapper.A_INTERNAL.Msg,
		}
	}

	if user == nil {
		return nil, &wrapper.ErrorWrapper{
			Code: wrapper.A_NOT_FOUND.Code,
			Msg:  wrapper.A_NOT_FOUND.Msg,
		}
	}

	return user, nil
}

func (s *UserService) Login(ctx *gin.Context, username, password string) (*string, *wrapper.ErrorWrapper) {
	if username == "" || password == "" {
		return nil, &wrapper.ErrorWrapper{
			Code:   wrapper.A_BAD_REQUEST.Code,
			Msg:    wrapper.A_BAD_REQUEST.Msg,
			Detail: wrapper.StringWrapper("username or password cannot be empty"),
		}
	}

	user, err := s.repo.GetByUsername(ctx, username)
	if err != nil {
		return nil, &wrapper.ErrorWrapper{
			Code: wrapper.A_INTERNAL.Code,
			Msg:  wrapper.A_INTERNAL.Msg,
		}
	}

	if user == nil {
		return nil, &wrapper.ErrorWrapper{
			Code: wrapper.A_NOT_FOUND.Code,
			Msg:  wrapper.A_NOT_FOUND.Msg,
		}
	}

	if user.IsActive == nil || !*user.IsActive {
		return nil, &wrapper.ErrorWrapper{
			Code: wrapper.A_NOT_FOUND.Code,
			Msg:  wrapper.A_NOT_FOUND.Msg,
		}
	}

	if !security.CheckPasswordHash(password, user.PasswordHash) {
		logger.Log.Error("invalid credentials")
		return nil, &wrapper.ErrorWrapper{
			Code: wrapper.A_CREDENTIALS.Code,
			Msg:  wrapper.A_CREDENTIALS.Msg,
		}
	}

	jwt, err := security.GenerateJWT(user)
	if err != nil {
		return nil, &wrapper.ErrorWrapper{
			Code: wrapper.A_INTERNAL.Code,
			Msg:  wrapper.A_INTERNAL.Msg,
		}
	}

	return &jwt, nil

}
