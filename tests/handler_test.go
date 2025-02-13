package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"short-url-service/internal/handler"
	"short-url-service/internal/storage"
	"testing"
)

// Мок хранилища
type MockStorage struct {
	storage.Storage
	PostFunc func(url string) (string, error)
	GetFunc  func(shortURL string) (string, error)
}

func (m *MockStorage) Post(url string) (string, error) {
	return m.PostFunc(url)
}

func (m *MockStorage) Get(shortURL string) (string, error) {
	return m.GetFunc(shortURL)
}

func TestPostHandler_Success(t *testing.T) {
	mockStorage := &MockStorage{
		PostFunc: func(url string) (string, error) {
			return "short123", nil
		},
	}

	handler := handler.PostHandler(mockStorage)
	reqBody := map[string]string{"url": "http://example.com"}
	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/api/url", bytes.NewBuffer(body))
	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Errorf("expected status 200 but got %d", recorder.Code)
	}

	var response map[string]string
	_ = json.NewDecoder(recorder.Body).Decode(&response)

	if response["short_url"] != "short123" {
		t.Errorf("expected short_url 'short123' but got %s", response["short_url"])
	}
}

func TestPostHandler_BadRequest_MissingURL(t *testing.T) {
	mockStorage := &MockStorage{
		PostFunc: func(url string) (string, error) {
			return "", nil
		},
	}

	handler := handler.PostHandler(mockStorage)
	reqBody := map[string]string{}
	body, _ := json.Marshal(reqBody)
	req := httptest.NewRequest("POST", "/api/url", bytes.NewBuffer(body))
	recorder := httptest.NewRecorder()

	handler.ServeHTTP(recorder, req)

	if recorder.Code != http.StatusBadRequest {
		t.Errorf("expected status 400 but got %d", recorder.Code)
	}
}
