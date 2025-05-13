#!/bin/bash

# Register a user
echo "Registering a user..."
curl --location 'http://localhost:8090/api/public/register' \
--header 'Content-Type: application/json' \
--data '{
    "username": "testuser",
    "password": "password123",
    "full_name": "Test User",
    "address": "Test Address"
}'

echo -e "\n\nLogging in..."
# Login to get a token
TOKEN=$(curl --location 'http://localhost:8090/api/public/login' \
--header 'Content-Type: application/json' \
--data '{
    "username": "testuser",
    "password": "password123"
}' | jq -r '.token')

echo "Token: $TOKEN"

echo -e "\nUploading image..."
# Create a test image
echo "This is a test image" > test_image.txt

# Upload the image using curl and the token we got
curl --location 'http://localhost:8090/api/private/upload' \
--header "Authorization: Bearer $TOKEN" \
--form 'image=@"./test_image.txt"'

echo -e "\n\nChecking if the file was saved..."
ls -la /Users/soanguyen/Development/golang/ct-backend-course-baonguyen-session-4-part-2-homework-init/static/images/
