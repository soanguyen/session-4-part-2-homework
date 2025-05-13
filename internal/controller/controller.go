package controller

import (
	"context"
	"ct-backend-course-baonguyen/internal/entity"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UseCase interface {
	Login(ctx context.Context, req *entity.LoginRequest) (*entity.LoginResponse, error)
	Register(ctx context.Context, req *entity.RegisterRequest) (*entity.RegisterResponse, error)
	// TODO: implement more
}

func NewHandler(uc UseCase) *Handler {
	return &Handler{uc: uc}
}

type Handler struct {
	uc UseCase
}

func (h *Handler) Register(c echo.Context) error {
	var req entity.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return fmt.Errorf("bind: %w", err)
	}

	resp, err := h.uc.Register(context.TODO(), &req)
	if err != nil {
		return fmt.Errorf("uc.Login: %w", err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) Login(c echo.Context) error {
	panic("TODO implement me")
}

func (h *Handler) Self(c echo.Context) error {
	panic("TODO implement me")
}

func (h *Handler) UploadImage(c echo.Context) error {
	panic("TODO implement me")
}
