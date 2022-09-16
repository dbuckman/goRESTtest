package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGet(t *testing.T) {
	t.Run("Get", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/get", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(Get)
		handler.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}
		// expected := `{"Method":"GET","Headers":{}}`
		// if rr.Body.String() != expected {
		// 	t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
		// }
	})
}

func TestDelete(t *testing.T) {
	t.Run("Get", func(t *testing.T) {
		req, err := http.NewRequest("DELETE", "/delete", nil)
		if err != nil {
			t.Fatal(err)
		}
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(Delete)
		handler.ServeHTTP(rr, req)
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}
		// expected := `{"Method":"GET","Headers":{}}`
		// if rr.Body.String() != expected {
		// 	t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
		// }
	})
}
