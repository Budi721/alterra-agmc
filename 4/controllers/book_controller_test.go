package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDeleteBookController(t *testing.T) {
	type args struct {
		c  echo.Context
		id uint
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DeleteBookController(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("DeleteBookController() error = %v, wantErr %v", err, tt.wantErr)
			}

			e := echo.New()
			req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/v1/book/%v", tt.args.id), nil)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			rec := httptest.NewRecorder()

			e.ServeHTTP(rec, req)
		})
	}
}

func TestGetBookController(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetBookController(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("GetBookController() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetBooksController(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GetBooksController(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("GetBooksController() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostBookController(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PostBookController(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("PostBookController() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPutBookController(t *testing.T) {
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PutBookController(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("PutBookController() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
