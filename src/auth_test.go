package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var baseURL = "http://api:8080" // Asegúrate de que esta URL es correcta para tu entorno de pruebas

func makeRequest(method, url string, body interface{}) (*http.Response, error) {
	bodyBuffer := new(bytes.Buffer)
	if err := json.NewEncoder(bodyBuffer).Encode(body); err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, url, bodyBuffer)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	return client.Do(req)
}

func TestSignupUserSuccess(t *testing.T) {
	body := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		IsAdmin  bool   `json:"is_admin"`
	}{
		Email:    "user@example.com",
		Password: "password123",
		IsAdmin:  false,
	}

	resp, err := makeRequest(http.MethodPost, baseURL+"/auth/signup", body)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200 but got %v", resp.StatusCode)
}

func TestSignupAdminSuccess(t *testing.T) {
	body := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		IsAdmin  bool   `json:"is_admin"`
	}{
		Email:    "admin@example.com",
		Password: "admin123",
		IsAdmin:  true,
	}

	resp, err := makeRequest(http.MethodPost, baseURL+"/auth/signup", body)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200 but got %v", resp.StatusCode)
}

func TestSignupUserConflict(t *testing.T) {
	// Primero, realiza un registro exitoso
	_, err := makeRequest(http.MethodPost, baseURL+"/auth/signup", struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		IsAdmin  bool   `json:"is_admin"`
	}{
		Email:    "user@example.com",
		Password: "password123",
		IsAdmin:  false,
	})
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}

	// Luego, intenta registrar el mismo usuario de nuevo
	resp, err := makeRequest(http.MethodPost, baseURL+"/auth/signup", struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		IsAdmin  bool   `json:"is_admin"`
	}{
		Email:    "user@example.com",
		Password: "password123",
		IsAdmin:  false,
	})
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusConflict, resp.StatusCode, "Expected status code 409 but got %v", resp.StatusCode)
}

func TestSignupAdminConflict(t *testing.T) {
	// Primero, realiza un registro exitoso
	_, err := makeRequest(http.MethodPost, baseURL+"/auth/signup", struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		IsAdmin  bool   `json:"is_admin"`
	}{
		Email:    "admin@example.com",
		Password: "admin123",
		IsAdmin:  true,
	})
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}

	// Luego, intenta registrar el mismo admin de nuevo
	resp, err := makeRequest(http.MethodPost, baseURL+"/auth/signup", struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		IsAdmin  bool   `json:"is_admin"`
	}{
		Email:    "admin@example.com",
		Password: "admin123",
		IsAdmin:  true,
	})
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusConflict, resp.StatusCode, "Expected status code 409 but got %v", resp.StatusCode)
}

func TestSigninUserSuccess(t *testing.T) {
	body := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
        IsAdmin  bool   `json:"is_admin"`
	}{
		Email:    "user@example.com",
		Password: "password123",
        IsAdmin:  false,
	}

	resp, err := makeRequest(http.MethodPost, baseURL+"/auth/signin", body)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200 but got %v", resp.StatusCode)
}

func TestSigninAdminSuccess(t *testing.T) {
	body := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
        IsAdmin  bool   `json:"is_admin"`
	}{
		Email:    "admin@example.com",
		Password: "admin123",
        IsAdmin:  true,
	}

	resp, err := makeRequest(http.MethodPost, baseURL+"/auth/signin", body)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode, "Expected status code 200 but got %v", resp.StatusCode)
}

func TestSigninUserNotRegistered(t *testing.T) {
	body := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
        IsAdmin  bool   `json:"is_admin"`
	}{
		Email:    "nonexistent@example.com",
		Password: "password123",
        IsAdmin:  false,
	}

	resp, err := makeRequest(http.MethodPost, baseURL+"/auth/signin", body)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode, "Expected status code 401 but got %v", resp.StatusCode)
}

func TestSigninAdminNotRegistered(t *testing.T) {
	body := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
        IsAdmin  bool   `json:"is_admin"`
	}{
		Email:    "nonexistentadmin@example.com",
		Password: "admin123",
        IsAdmin:  true,
	}

	resp, err := makeRequest(http.MethodPost, baseURL+"/auth/signin", body)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode, "Expected status code 401 but got %v", resp.StatusCode)
}

func TestSigninUserIncorrectPassword(t *testing.T) {
	// Primero, realiza un registro exitoso
	_, err := makeRequest(http.MethodPost, baseURL+"/auth/signup", struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		IsAdmin  bool   `json:"is_admin"`
	}{
		Email:    "user@example.com",
		Password: "password123",
		IsAdmin:  false,
	})
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}

	// Luego, intenta iniciar sesión con una contraseña incorrecta
	body := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
        IsAdmin  bool   `json:"is_admin"`
	}{
		Email:    "user@example.com",
		Password: "wrongpassword",
        IsAdmin:  false,
	}

	resp, err := makeRequest(http.MethodPost, baseURL+"/auth/signin", body)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode, "Expected status code 401 but got %v", resp.StatusCode)
}

func TestSigninAdminIncorrectPassword(t *testing.T) {
	// Primero, realiza un registro exitoso
	_, err := makeRequest(http.MethodPost, baseURL+"/auth/signup", struct {
		Email    string `json:"email"`
		Password string `json:"password"`
		IsAdmin  bool   `json:"is_admin"`
	}{
		Email:    "admin@example.com",
		Password: "admin123",
		IsAdmin:  true,
	})
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}

	// Luego, intenta iniciar sesión con una contraseña incorrecta
	body := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
        IsAdmin  bool   `json:"is_admin"`
	}{
		Email:    "admin@example.com",
		Password: "wrongpassword",
        IsAdmin:  true,
	}

	resp, err := makeRequest(http.MethodPost, baseURL+"/auth/signin", body)
	if err != nil {
		t.Fatalf("Failed to make request: %v", err)
	}
	defer resp.Body.Close()

	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode, "Expected status code 401 but got %v", resp.StatusCode)
}