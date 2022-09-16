package controllers

import (
	"github.com/Budi721/alterra-agmc/v2/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestDeleteBookController(t *testing.T) {
	type args struct {
		id string
	}

	type expectation struct {
		code           int
		expectedResult any
		isError        bool
	}

	tests := []struct {
		name string
		args
		expectation
	}{
		{
			name: "success response 200",
			args: args{
				id: "1",
			},
			expectation: expectation{
				code:           http.StatusOK,
				expectedResult: "{\"status\":\"success\",\"code\":200,\"data\":{\"id\":1,\"title\":\"Anak Singkong\",\"author\":\"Chairil Tanjung\",\"price\":50000}}\n",
				isError:        false,
			},
		},
		{
			name: "not found response",
			args: args{
				id: "3",
			},
			expectation: expectation{
				code:           http.StatusNotFound,
				expectedResult: models.Response{Status: "not found", Code: 0x194, Data: interface{}(nil)},
				isError:        true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			r := httptest.NewRequest(http.MethodDelete, "/v1/books", nil)
			r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			w := httptest.NewRecorder()
			c := e.NewContext(r, w)
			c.SetParamNames("id")
			c.SetParamValues(tt.args.id)

			err := DeleteBookController(c)

			if tt.expectation.isError && assert.Error(t, err) {
				httpError := err.(*echo.HTTPError)
				assert.Equal(t, tt.expectation.code, httpError.Code)
				assert.Equal(t, tt.expectation.expectedResult, httpError.Message)
			}

			if !tt.expectation.isError && assert.NoError(t, err) {
				assert.Equal(t, tt.expectation.code, w.Code)
				assert.Equal(t, tt.expectation.expectedResult, w.Body.String())
			}
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
