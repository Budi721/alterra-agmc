package controllers

import (
	"github.com/labstack/echo/v4"
	"testing"
)

func TestDeleteUserController(t *testing.T) {
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
			if err := DeleteUserController(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("DeleteUserController() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetUserController(t *testing.T) {
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
			if err := GetUserController(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("GetUserController() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetUsersController(t *testing.T) {
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
			if err := GetUsersController(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("GetUsersController() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLoginUserController(t *testing.T) {
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
			if err := LoginUserController(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("LoginUserController() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPostUserController(t *testing.T) {
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
			if err := PostUserController(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("PostUserController() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPutUserController(t *testing.T) {
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
			if err := PutUserController(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("PutUserController() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
