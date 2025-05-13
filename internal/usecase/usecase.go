package usecase

import (
	"context"
	"ct-backend-course-baonguyen/internal/entity"
	"io"
)

type UserStore interface {
	Save(info entity.UserInfo) error
	Get(username string) (entity.UserInfo, error)
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
	panic("TODO implement me")
}

func (uc *ucImplement) Self(ctx context.Context, req *entity.SelfRequest) (*entity.SelfResponse, error) {
	panic("TODO implement me")
}

func (uc *ucImplement) UploadImage(ctx context.Context, req *entity.RegisterRequest) (*entity.RegisterResponse, error) {
	panic("TODO implement me")
}
