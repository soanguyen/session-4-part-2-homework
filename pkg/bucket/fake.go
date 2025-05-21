package bucket

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func NewFake() *fake {
	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Error getting current working directory: %v\n", err)
		// Fallback to relative path
		cwd = "."
	}

	// Create the directory if it doesn't exist
	uploadDir := filepath.Join(cwd, "static", "images")
	fmt.Printf("DEBUG: Upload directory path: %s\n", uploadDir)

	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		fmt.Printf("Error creating upload directory: %v\n", err)
	}
	return &fake{uploadDir: uploadDir}
}

type fake struct {
	uploadDir string
}

func (f *fake) SaveImage(ctx context.Context, name string, r io.Reader) (string, error) {
	// Generate file path
	fileName := filepath.Base(name)

	ext := strings.ToLower(filepath.Ext(fileName))
	validExtensions := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
		".bmp":  true,
		".webp": true,
	}
	if !validExtensions[ext] {
		return "", fmt.Errorf("unsupported image format: %s. Supported formats: jpg, jpeg, png, gif, bmp, webp", ext)
	}

	filePath := filepath.Join(f.uploadDir, fileName)

	fmt.Printf("DEBUG: Saving file to path: %s\n", filePath)

	// Create the file
	out, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("DEBUG: Error creating file: %v\n", err)
		return "", fmt.Errorf("failed to create file: %w", err)
	}
	defer out.Close()

	// Copy the uploaded file to the destination file
	bytesWritten, err := io.Copy(out, r)
	if err != nil {
		fmt.Printf("DEBUG: Error writing file: %v\n", err)
		return "", fmt.Errorf("failed to write file: %w", err)
	}

	fmt.Printf("DEBUG: Successfully wrote %d bytes to file: %s\n", bytesWritten, filePath)

	// Return the relative URL path
	return fmt.Sprintf("/static/images/%s", fileName), nil
}
