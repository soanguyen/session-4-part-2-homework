// You can edit this code!
// Click here and start typing.
package main

import (
	"ct-backend-course-baonguyen/internal/controller"
	inmemory "ct-backend-course-baonguyen/internal/storage/in-memory"
	"ct-backend-course-baonguyen/internal/usecase"
	"ct-backend-course-baonguyen/pkg/auth"
	"ct-backend-course-baonguyen/pkg/bucket"
	"ct-backend-course-baonguyen/pkg/validator"
	"fmt"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func main() {
	userStore := inmemory.NewUserStore()
	imgBucket := bucket.NewFake()
	uc := usecase.NewUseCase(userStore, imgBucket)
	hdl := controller.NewHandler(uc)

	srv := newServer(hdl)
	if err := srv.Start(":8090"); err != nil {
		log.Error(err)
	}
}

func newServer(hdl *controller.Handler) *echo.Echo {
	e := echo.New()
	e.Validator = validator.NewCustomValidator()

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// Serve static files
	cwd, err := os.Getwd()
	if err != nil {
		log.Errorf("Error getting current working directory: %v", err)
		cwd = "."
	}
	staticDir := filepath.Join(cwd, "static")
	fmt.Printf("DEBUG: Static files directory: %s\n", staticDir)
	e.Static("/static", staticDir)

	public := e.Group("/api/public")
	private := e.Group("/api/private", auth.JWTMiddleware())

	public.POST("/register", hdl.Register)
	public.POST("/login", hdl.Login)

	private.GET("/self", hdl.Self)
	private.POST("/upload", hdl.UploadImage)

	return e
}
