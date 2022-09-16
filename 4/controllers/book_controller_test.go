package controllers

import (
	"github.com/Budi721/alterra-agmc/v2/lib/mock"
	"github.com/Budi721/alterra-agmc/v2/middlewares"
	"github.com/Budi721/alterra-agmc/v2/models"
	"github.com/go-playground/validator/v10"
	"net/http"
	"net/http/httptest"
	"strings"
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
		wantError      bool
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
				wantError:      false,
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
				wantError:      true,
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

func TestGetBookController(t *testing.T) {
	type args struct {
		id string
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
			name: "success response 200",
			args: args{
				id: "1",
			},
			expectation: expectation{
				code:           http.StatusOK,
				expectedResult: "{\"status\":\"success\",\"code\":200,\"data\":{\"id\":1,\"title\":\"Anak Singkong\",\"author\":\"Chairil Tanjung\",\"price\":50000}}\n",
				wantError:      false,
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
				wantError:      true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			r := httptest.NewRequest(http.MethodGet, "/v1/books", nil)
			r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			w := httptest.NewRecorder()
			c := e.NewContext(r, w)
			c.SetParamNames("id")
			c.SetParamValues(tt.args.id)

			err := GetBookController(c)

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

func TestGetBooksController(t *testing.T) {
	type expectation struct {
		code           int
		expectedResult any
		wantError      bool
	}

	tests := []struct {
		name string
		expectation
	}{
		{
			name: "success response 200",
			expectation: expectation{
				code:           http.StatusOK,
				expectedResult: "{\"status\":\"success\",\"code\":200,\"data\":[{\"id\":1,\"title\":\"Anak Singkong\",\"author\":\"Chairil Tanjung\",\"price\":50000},{\"id\":2,\"title\":\"Garis Waktu\",\"author\":\"Fiersa Besari\",\"price\":35000}]}\n",
				wantError:      false,
			},
		},
		{
			name: "not found response",
			expectation: expectation{
				code:           http.StatusNotFound,
				expectedResult: models.Response{Status: "not found", Code: 0x194, Data: []models.Book{}},
				wantError:      true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			r := httptest.NewRequest(http.MethodGet, "/v1/books", nil)
			r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			w := httptest.NewRecorder()
			c := e.NewContext(r, w)

			// remove all static data
			var temporary []models.Book
			if tt.expectation.wantError {
				copy(temporary, mock.Books)
				mock.Books = []models.Book{}
			}
			err := GetBooksController(c)
			if tt.expectation.wantError {
				mock.Books = temporary
			}

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

func TestPostBookController(t *testing.T) {
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
			name: "success response 200",
			args: args{
				json: `
				{
				  "id": 5,
				  "title": "Example Test",
				  "author": "Budi Rahmawan",
				  "price": 100000
				}`,
			},
			expectation: expectation{
				code:           http.StatusCreated,
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
			r := httptest.NewRequest(http.MethodPost, "/v1/books", strings.NewReader(tt.args.json))
			r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			w := httptest.NewRecorder()
			c := e.NewContext(r, w)

			err := PostBookController(c)

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

func TestPutBookController(t *testing.T) {
	type args struct {
		id   string
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
			name: "success response 200",
			args: args{
				id: "1",
				json: `
				{
				  "id": 3,
				  "title": "Example Test",
				  "author": "Budi Rahmawan",
				  "price": 100000
				}`,
			},
			expectation: expectation{
				code:           http.StatusCreated,
				expectedResult: "{\"status\":\"success\",\"code\":200,\"data\":{\"id\":3,\"title\":\"Example Test\",\"author\":\"Budi Rahmawan\",\"price\":100000}}\n",
				wantError:      false,
			},
		},
		{
			name: "bad request response",
			args: args{
				id:   "2",
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
			r := httptest.NewRequest(http.MethodPut, "/v1/books", strings.NewReader(tt.args.json))
			r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			w := httptest.NewRecorder()
			c := e.NewContext(r, w)
			c.SetParamNames("id")
			c.SetParamValues(tt.args.id)

			err := PutBookController(c)

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
