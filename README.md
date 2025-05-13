# Backend Course: User Authentication & Image Upload API

This project implements a REST API for user authentication and image upload using Go and the Echo framework.

## Features
- User registration and login
- JWT-based authentication
- User profile management
- Image upload for authenticated users

## Project Structure
The project follows a clean 3-layer architecture:
- **Controller Layer**: Handles HTTP requests and responses
- **UseCase Layer**: Contains business logic
- **Storage Layer**: Handles data persistence

## API Endpoints

### Public Endpoints
- `POST /api/public/register`: Register a new user
- `POST /api/public/login`: Authenticate and receive a JWT token

### Private Endpoints (Requires Authentication)
- `GET /api/private/self`: Get authenticated user's profile
- `POST /api/private/upload`: Upload an image for the authenticated user

## Running the Application

### Using Go directly
```shell
go mod download
go run main.go
```

### Using Docker
```shell
docker build -t user-image-api .
docker run -p 8090:8090 user-image-api
```

### Using Docker Compose
```shell
docker-compose up
```

## API Usage Examples

### Register a User
```shell
curl --location 'localhost:8090/api/public/register' \
--header 'Content-Type: application/json' \
--data '{
    "username": "testuser",
    "password": "password123",
    "full_name": "Test User",
    "address": "123 Test St"
}'
```

### Login
```shell
curl --location 'localhost:8090/api/public/login' \
--header 'Content-Type: application/json' \
--data '{
    "username": "testuser",
    "password": "password123"
}'
```

### Get User Profile
```shell
curl --location 'localhost:8090/api/private/self' \
--header 'Authorization: Bearer YOUR_JWT_TOKEN'
```

### Upload Image
```shell
curl --location 'localhost:8090/api/private/upload' \
--header 'Authorization: Bearer YOUR_JWT_TOKEN' \
--form 'image=@"/path/to/your/image.jpg"'
```

## Original Tasks
TODO #1: Restructure to 3-layers (storage, usecase, controller) ✅  
TODO #2: Write private API to upload image and save image info ✅  
TODO #3: Build a Docker image ✅  
TODO #4: Write unit tests for that private API ❌

