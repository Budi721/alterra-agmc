package controllers

import (
	"github.com/Budi721/alterra-agmc/v2/middlewares"
	"github.com/Budi721/alterra-agmc/v2/models"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestLoginUserController(t *testing.T) {
	type args struct {
		json string
	}

	type expectation struct {
		code           int
		expectedResult any
		wantError      bool
	}

	tests := []struct {
		name string
		args
		expectation
	}{
		{
			name: "success login response 200",
			args: args{
				json: `
				{
				  "email": "raharjobd@gmail.com",
				  "password": "123456"
				}`,
			},
			expectation: expectation{
				code:           http.StatusOK,
				expectedResult: "{\"status\":\"success\",\"code\":201,\"data\":{\"id\":5,\"title\":\"Example Test\",\"author\":\"Budi Rahmawan\",\"price\":100000}}\n",
				wantError:      false,
			},
		},
		{
			name: "bad request response",
			args: args{
				json: "{}",
			},
			expectation: expectation{
				code:           http.StatusBadRequest,
				expectedResult: models.Response{Status: "bad request: Key: 'Book.ID' Error:Field validation for 'ID' failed on the 'required' tag\nKey: 'Book.Title' Error:Field validation for 'Title' failed on the 'required' tag\nKey: 'Book.Author' Error:Field validation for 'Author' failed on the 'required' tag\nKey: 'Book.Price' Error:Field validation for 'Price' failed on the 'required' tag", Code: 0x190, Data: interface{}(nil)},
				wantError:      true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			e.Validator = &middlewares.CustomValidator{Validator: validator.New()}
			r := httptest.NewRequest(http.MethodPost, "/v1/login", strings.NewReader(tt.args.json))
			r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			w := httptest.NewRecorder()
			c := e.NewContext(r, w)

			err := LoginUserController(c)

			if tt.expectation.wantError && assert.Error(t, err) {
				httpError := err.(*echo.HTTPError)
				assert.Equal(t, tt.expectation.code, httpError.Code)
				assert.Equal(t, tt.expectation.expectedResult, httpError.Message)
			}

			if !tt.expectation.wantError && assert.NoError(t, err) {
				assert.Equal(t, tt.expectation.code, w.Code)
				assert.Equal(t, tt.expectation.expectedResult, w.Body.String())
			}
		})
	}
}

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
