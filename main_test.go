package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestHomeHandler tests the home page handler
func TestHomeHandler(t *testing.T) {
	// Create a new HTTP request
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()

	// Create the handler to test
	handler := http.HandlerFunc(homeHandler)

	// Serve the HTTP request
	handler.ServeHTTP(rr, req)

	// Check if the status code is 200 OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check if Content-Type is correct
	expectedContentType := "text/html; charset=utf-8"
	if contentType := rr.Header().Get("Content-Type"); contentType != expectedContentType {
		t.Errorf("handler returned unexpected content type: got %v want %v",
			contentType, expectedContentType)
	}
}
