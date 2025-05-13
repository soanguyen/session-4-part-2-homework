package entity

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
}

type LoginResponse struct {
}

type SelfRequest struct {
}

type SelfResponse struct {
}

type UploadImageRequest struct {
}

type UploadImageResponse struct {
}

type UserInfo struct {
	Username string `json:"username"`
	Password string `json:"password"`
	FullName string `json:"full_name"`
	Address  string `json:"address"`
}

type ImageInfo struct {
	// TODO
}
