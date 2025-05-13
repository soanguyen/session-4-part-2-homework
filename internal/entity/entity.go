package entity

import "io"

type RegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
	FullName string `json:"full_name"`
	Address  string `json:"address"`
}

type RegisterResponse struct {
	UserId string `json:"user_id"`
}

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type SelfRequest struct {
	Username string `json:"username"`
}

type SelfResponse struct {
	Username string `json:"username"`
	FullName string `json:"full_name"`
	Address  string `json:"address"`
	ImageURL string `json:"image_url,omitempty"`
}

type UploadImageRequest struct {
	Username string    `json:"username"`
	ImageData io.Reader `:"-"`
	ImageName string    `json:"image_name"`
}

type UploadImageResponse struct {
	ImageURL string `json:"image_url"`
}

type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
	Address  string `json:"address"`
	ImageURL string `json:"image_url,omitempty"`
}

type ImageInfo struct {
	Username string `json:"username"`
	ImageURL string `json:"image_url"`
	UploadedAt string `json:"uploaded_at"`
}
