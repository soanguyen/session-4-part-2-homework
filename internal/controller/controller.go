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
	Self(ctx context.Context, req *entity.SelfRequest) (*entity.SelfResponse, error)
	UploadImage(ctx context.Context, req *entity.UploadImageRequest) (*entity.UploadImageResponse, error)
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
	var req entity.LoginRequest
	if err := c.Bind(&req); err != nil {
		return fmt.Errorf("bind: %w", err)
	}

	resp, err := h.uc.Login(context.TODO(), &req)
	if err != nil {
		return fmt.Errorf("uc.Login: %w", err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) Self(c echo.Context) error {
	// Extract username from JWT token (set by middleware)
	username := c.Get("username").(string)

	// Prepare request
	req := &entity.SelfRequest{
		Username: username,
	}

	// Call usecase
	resp, err := h.uc.Self(context.TODO(), req)
	if err != nil {
		return fmt.Errorf("uc.Self: %w", err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *Handler) UploadImage(c echo.Context) error {
	// Parse multipart form
	file, err := c.FormFile("image")
	if err != nil {
		return fmt.Errorf("form file: %w", err)
	}

	// Open the file
	src, err := file.Open()
	if err != nil {
		return fmt.Errorf("file open: %w", err)
	}
	defer src.Close()

	// Extract username from JWT token (set by middleware)
	username := c.Get("username").(string)

	// Prepare request
	req := &entity.UploadImageRequest{
		Username: username,
		ImageData: src,
		ImageName: file.Filename,
	}

	// Call usecase
	resp, err := h.uc.UploadImage(context.TODO(), req)
	if err != nil {
		return fmt.Errorf("uc.UploadImage: %w", err)
	}

	return c.JSON(http.StatusOK, resp)
}
