package controllers

import (
	"bytes"

	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"

	"testing"

	"UserService/internal/models"
	"UserService/internal/repositories"
	"UserService/internal/services"
)

func TestUserController_CreateUserAlreadyExists(t *testing.T) {
	r := gin.New()

    expectedUser := models.User{Email: "email", Password: "password"}
    expectedUsers := map[string]models.User{expectedUser.Email: expectedUser}
    userRepository := repositories.NewUserRepositoryImplWithUsers(expectedUsers)
    userService := services.NewUserServiceImpl(userRepository)
    userController := NewUserControllerImpl(userService)


    r.POST("/api/v1/users", userController.CreateUser)
	// Create a test HTTP request
    jsonBody := []byte(`{"email": "email", "password": "password"}`)
    bodyReader := bytes.NewReader(jsonBody)
	req, err := http.NewRequest("POST", "/api/v1/users", bodyReader)
	if err != nil {
		t.Fatal(err)
	}
	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	// Serve the HTTP request
	r.ServeHTTP(rr, req)
	// Check the status code
	if status := rr.Code; status != http.StatusConflict {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusConflict)
	}
}

func TestUserController_CreateUser(t *testing.T) {
	r := gin.New()

    userRepository := repositories.NewUserRepositoryImpl()
    userService := services.NewUserServiceImpl(userRepository)
    userController := NewUserControllerImpl(userService)


    r.POST("/api/v1/users", userController.CreateUser)
	// Create a test HTTP request
    jsonBody := []byte(`{"email": "email", "password": "password"}`)
    bodyReader := bytes.NewReader(jsonBody)
	req, err := http.NewRequest("POST", "/api/v1/users", bodyReader)
	if err != nil {
		t.Fatal(err)
	}
	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	// Serve the HTTP request
	r.ServeHTTP(rr, req)
	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
