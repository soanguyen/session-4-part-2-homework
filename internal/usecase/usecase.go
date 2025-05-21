package usecase

import (
	"context"
	"ct-backend-course-baonguyen/internal/entity"
	"ct-backend-course-baonguyen/pkg/auth"
	"fmt"
	"io"
	"time"
)

type UserStore interface {
	Save(info entity.UserInfo) error
	Get(username string) (entity.UserInfo, error)
	UpdateImageURL(username string, imageURL string) error
}

type ImageBucket interface {
	SaveImage(ctx context.Context, name string, r io.Reader) (string, error)
}

func NewUseCase(userStore UserStore, imageBucket ImageBucket) *ucImplement {
	return &ucImplement{
		userStore: userStore,
		imgBucket: imageBucket,
	}
}

type ucImplement struct {
	userStore UserStore
	imgBucket ImageBucket
}

func (uc *ucImplement) Register(ctx context.Context, req *entity.RegisterRequest) (*entity.RegisterResponse, error) {
	if err := uc.userStore.Save(entity.UserInfo{
		Username: req.Username,
		Password: req.Password,
		FullName: req.FullName,
		Address:  req.Address,
	}); err != nil {
		return nil, err
	}

	return &entity.RegisterResponse{UserId: req.Username}, nil
}

func (uc *ucImplement) Login(ctx context.Context, req *entity.LoginRequest) (*entity.LoginResponse, error) {
	userInfo, err := uc.userStore.Get(req.Username)
	if err != nil {
		return nil, err
	}

	if userInfo.Password != req.Password {
		return nil, fmt.Errorf("invalid password")
	}

	token, err := auth.GenerateToken(req.Username, 1*time.Hour)
	if err != nil {
		return nil, err
	}

	return &entity.LoginResponse{Token: token}, nil
}

func (uc *ucImplement) Self(ctx context.Context, req *entity.SelfRequest) (*entity.SelfResponse, error) {
	userInfo, err := uc.userStore.Get(req.Username)
	if err != nil {
		return nil, err
	}

	return &entity.SelfResponse{
		Username: userInfo.Username,
		FullName: userInfo.FullName,
		Address:  userInfo.Address,
		ImageURL: userInfo.ImageURL,
	}, nil
}

func (uc *ucImplement) UploadImage(ctx context.Context, req *entity.UploadImageRequest) (*entity.UploadImageResponse, error) {
	imageURL, err := uc.imgBucket.SaveImage(ctx, req.ImageName, req.ImageData)
	if err != nil {
		return nil, fmt.Errorf("failed to save image: %w", err)
	}

	err = uc.userStore.UpdateImageURL(req.Username, imageURL)
	if err != nil {
		return nil, fmt.Errorf("failed to update user image URL: %w", err)
	}

	return &entity.UploadImageResponse{
		ImageURL: imageURL,
	}, nil
}
