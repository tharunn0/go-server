package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func init() {
	// Set gin to TestMode to suppress debug logging during tests
	gin.SetMode(gin.TestMode)
}

func TestHealthEndpoint(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)

	// Assert HTTP response status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Parse JSON response
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to parse JSON response: %v", err)
	}

	// Assert fields
	if status, ok := response["status"].(string); !ok || status != "healthy" {
		t.Errorf("Expected status 'healthy', got '%v'", response["status"])
	}

	if version, ok := response["version"].(string); !ok || version != Version {
		t.Errorf("Expected version '%s', got '%v'", Version, response["version"])
	}

	if timestamp, ok := response["timestamp"].(string); !ok || timestamp == "" {
		t.Errorf("Expected non-empty timestamp, got '%v'", response["timestamp"])
	}
}

func TestMessagesEndpoint(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/messages", nil)
	router.ServeHTTP(w, req)

	// Assert HTTP response status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Parse JSON response
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to parse JSON response: %v", err)
	}

	if message, ok := response["message"].(string); !ok || message == "" {
		t.Errorf("Expected non-empty message, got '%v'", response["message"])
	}

	if success, ok := response["success"].(bool); !ok || !success {
		t.Errorf("Expected success true, got '%v'", response["success"])
	}
}

func TestUsersEndpoint(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/users", nil)
	router.ServeHTTP(w, req)

	// Assert HTTP response status code
	if w.Code != http.StatusOK {
		t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
	}

	// Parse JSON response
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to parse JSON response: %v", err)
	}

	if total, ok := response["total"].(float64); !ok || total != 3 {
		t.Errorf("Expected total 3, got '%v'", response["total"])
	}

	if success, ok := response["success"].(bool); !ok || !success {
		t.Errorf("Expected success true, got '%v'", response["success"])
	}

	data, ok := response["data"].([]interface{})
	if !ok || len(data) != 3 {
		t.Errorf("Expected data slice with length 3, got %T of length %d", response["data"], len(data))
	}
}
