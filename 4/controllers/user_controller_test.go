package controllers

import (
	"github.com/Budi721/alterra-agmc/v2/constants"
	"github.com/Budi721/alterra-agmc/v2/middlewares"
	"github.com/Budi721/alterra-agmc/v2/models"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
)

// userRepositoryMock mocking struct implement UserRepository interface
type userRepositoryMock struct{ Mock mock.Mock }

func (repository *userRepositoryMock) GetUsers() ([]models.User, error) {
	arguments := repository.Mock.Called()
	if arguments.Get(0) == nil {
		return []models.User{
			{
				ID:       1,
				Name:     "Budi",
				Email:    "rahmawanbd@gmail.com",
				Password: "123456",
			},
			{
				ID:       1,
				Name:     "Budi",
				Email:    "rahmawanbd@gmail.com",
				Password: "123456",
			},
		}, nil
	}

	return []models.User{}, gorm.ErrRecordNotFound
}

func (repository *userRepositoryMock) GetUser(id uint) (*models.User, error) {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == "success" {
		return &models.User{
			ID:       1,
			Name:     "Budi",
			Email:    "rahmawanbd@gmail.com",
			Password: "123456",
		}, nil
	} else {
		return &models.User{}, gorm.ErrRecordNotFound
	}
}

func (repository *userRepositoryMock) CreateUser(name string, email string, password string) (*models.User, error) {
	arguments := repository.Mock.Called(name, email, password)
	if arguments.Bool(0) {
		return &models.User{
			ID:       1,
			Name:     name,
			Email:    email,
			Password: password,
		}, nil
	}

	return &models.User{}, gorm.ErrRecordNotFound
}

func (repository *userRepositoryMock) UpdateUser(id uint, name string, email string, password string) (*models.User, error) {
	arguments := repository.Mock.Called(id, name, email, password)
	if arguments.Bool(0) {
		return &models.User{
			ID:       id,
			Name:     name,
			Email:    email,
			Password: password,
		}, nil
	}

	return &models.User{}, gorm.ErrRecordNotFound
}

func (repository *userRepositoryMock) DeleteUser(id uint) (*models.User, error) {
	arguments := repository.Mock.Called(id)
	if arguments.Get(0) == "success" {
		return &models.User{ID: id}, nil
	}

	return &models.User{}, gorm.ErrRecordNotFound
}

func (repository *userRepositoryMock) LoginUser(email string, password string) (string, error) {
	arguments := repository.Mock.Called(email, password)
	if arguments.Get(0) == "raharjobd@gmail.com" && arguments.Get(1) == "123456" {
		return "ini token jwt", nil
	} else {
		return "", gorm.ErrRecordNotFound
	}
}

var (
	urMock = &userRepositoryMock{
		Mock: mock.Mock{},
	}

	userController = NewUserController(urMock)
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
				expectedResult: "{\"status\":\"success\",\"code\":200,\"data\":{\"token\":\"ini token jwt\"}}\n",
				wantError:      false,
			},
		},
		{
			name: "not found login response 404",
			args: args{
				json: `
				{
				  "email": "rahmawan@dummy.com",
				  "password": "123456"
				}`,
			},
			expectation: expectation{
				code:           http.StatusNotFound,
				expectedResult: models.Response{Status: "not found", Code: 0x194, Data: interface{}(nil)},
				wantError:      true,
			},
		},
		{
			name: "bad request response",
			args: args{
				json: "{}",
			},
			expectation: expectation{
				code:           http.StatusBadRequest,
				expectedResult: models.Response{Status: "bad request: Key: 'UserLogin.Email' Error:Field validation for 'Email' failed on the 'required' tag\nKey: 'UserLogin.Password' Error:Field validation for 'Password' failed on the 'required' tag", Code: 0x190, Data: interface{}(nil)},
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

			urMock.Mock.On("LoginUser", "raharjobd@gmail.com", "123456").Return("raharjobd@gmail.com", "123456")
			urMock.Mock.On("LoginUser", "rahmawan@dummy.com", "123456").Return("rahmawan@dummy.com", "123456")
			err := userController.LoginUserController(c)

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
			name: "success login response 200",
			args: args{
				id: "1",
			},
			expectation: expectation{
				code:           http.StatusOK,
				expectedResult: "{\"status\":\"success\",\"code\":200,\"data\":{\"id\":1,\"name\":\"\",\"email\":\"\",\"password\":\"\"}}\n",
				wantError:      false,
			},
		},
		{
			name: "not found login response 404",
			args: args{
				id: "99",
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
			e.Validator = &middlewares.CustomValidator{Validator: validator.New()}
			e.Use(middlewares.JWTMiddleware)

			userId, _ := strconv.Atoi(tt.args.id)
			got, _ := middlewares.GenerateToken(uint(userId))
			token, _ := jwt.Parse(got, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					t.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}

				return []byte(constants.SecretJwt), nil
			})

			r := httptest.NewRequest(http.MethodDelete, "/v1/users", nil)
			r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			w := httptest.NewRecorder()

			c := e.NewContext(r, w)
			c.Set("user", token)
			c.SetParamNames("id")
			c.SetParamValues(tt.args.id)

			urMock.Mock.On("DeleteUser", uint(1)).Return("success")
			urMock.Mock.On("DeleteUser", uint(99)).Return(nil)
			err := userController.DeleteUserController(c)

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

func TestGetUserController(t *testing.T) {
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
			name: "success login response 200",
			args: args{
				id: `1`,
			},
			expectation: expectation{
				code:           http.StatusOK,
				expectedResult: "{\"status\":\"success\",\"code\":200,\"data\":{\"id\":1,\"name\":\"Budi\",\"email\":\"rahmawanbd@gmail.com\",\"password\":\"123456\"}}\n",
				wantError:      false,
			},
		},
		{
			name: "not found login response 404",
			args: args{
				id: `99`,
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
			e.Validator = &middlewares.CustomValidator{Validator: validator.New()}

			r := httptest.NewRequest(http.MethodGet, "/v1/users", nil)
			r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			w := httptest.NewRecorder()

			c := e.NewContext(r, w)
			c.SetParamNames("id")
			c.SetParamValues(tt.args.id)

			urMock.Mock.On("GetUser", uint(1)).Return("success")
			urMock.Mock.On("GetUser", uint(99)).Return("not_found")
			err := userController.GetUserController(c)

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

func TestGetUsersController(t *testing.T) {
	type expectation struct {
		code           int
		expectedResult any
	}

	tests := []struct {
		name string
		expectation
	}{
		{
			name: "success login response 200",
			expectation: expectation{
				code:           http.StatusOK,
				expectedResult: "{\"status\":\"success\",\"code\":200,\"data\":[{\"id\":1,\"name\":\"Budi\",\"email\":\"rahmawanbd@gmail.com\",\"password\":\"123456\"},{\"id\":1,\"name\":\"Budi\",\"email\":\"rahmawanbd@gmail.com\",\"password\":\"123456\"}]}\n",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			e.Validator = &middlewares.CustomValidator{Validator: validator.New()}

			r := httptest.NewRequest(http.MethodGet, "/v1/users", nil)
			r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			w := httptest.NewRecorder()

			c := e.NewContext(r, w)

			urMock.Mock.On("GetUsers").Return(nil)
			err := userController.GetUsersController(c)

			if assert.NoError(t, err) {
				assert.Equal(t, tt.expectation.code, w.Code)
				assert.Equal(t, tt.expectation.expectedResult, w.Body.String())
			}
		})
	}
}

func TestPostUserController(t *testing.T) {
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
				  "name": "Budi Raharjo",
				  "email": "raharjobd@gmail.com",
				  "password": "123456"
				}`,
			},
			expectation: expectation{
				code:           http.StatusCreated,
				expectedResult: "{\"status\":\"success\",\"code\":200,\"data\":{\"id\":1,\"name\":\"Budi Raharjo\",\"email\":\"raharjobd@gmail.com\",\"password\":\"123456\"}}\n",
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
				expectedResult: models.Response{Status: "bad request: Key: 'User.Name' Error:Field validation for 'Name' failed on the 'required' tag\nKey: 'User.Email' Error:Field validation for 'Email' failed on the 'required' tag\nKey: 'User.Password' Error:Field validation for 'Password' failed on the 'required' tag", Code: 0x190, Data: interface{}(nil)},
				wantError:      true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			e.Validator = &middlewares.CustomValidator{Validator: validator.New()}
			r := httptest.NewRequest(http.MethodPost, "/v1/users", strings.NewReader(tt.args.json))
			r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			w := httptest.NewRecorder()
			c := e.NewContext(r, w)

			urMock.Mock.On("CreateUser", "Budi Raharjo", "raharjobd@gmail.com", "123456").Return(true)
			err := userController.PostUserController(c)

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

func TestPutUserController(t *testing.T) {
	type args struct {
		id   uint
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
				id: 1,
				json: `
				{
				  "name": "Budi Raharjo",
				  "email": "raharjobd@gmail.com",
				  "password": "123456"
				}`,
			},
			expectation: expectation{
				code:           http.StatusOK,
				expectedResult: "{\"status\":\"success\",\"code\":200,\"data\":{\"id\":1,\"name\":\"Budi Raharjo\",\"email\":\"raharjobd@gmail.com\",\"password\":\"123456\"}}\n",
				wantError:      false,
			},
		},
		{
			name: "bad request response",
			args: args{
				id:   2,
				json: "{}",
			},
			expectation: expectation{
				code:           http.StatusBadRequest,
				expectedResult: models.Response{Status: "bad request: Key: 'UserUpdate.Name' Error:Field validation for 'Name' failed on the 'required' tag\nKey: 'UserUpdate.Email' Error:Field validation for 'Email' failed on the 'required' tag\nKey: 'UserUpdate.Password' Error:Field validation for 'Password' failed on the 'required' tag", Code: 0x190, Data: interface{}(nil)},
				wantError:      true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := echo.New()
			e.Validator = &middlewares.CustomValidator{Validator: validator.New()}
			got, _ := middlewares.GenerateToken(tt.args.id)
			token, _ := jwt.Parse(got, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					t.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}

				return []byte(constants.SecretJwt), nil
			})

			r := httptest.NewRequest(http.MethodPut, "/v1/users", strings.NewReader(tt.args.json))
			r.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			w := httptest.NewRecorder()
			c := e.NewContext(r, w)
			c.Set("user", token)
			c.SetParamNames("id")
			c.SetParamValues(strconv.Itoa(int(tt.args.id)))

			urMock.Mock.On("UpdateUser", tt.args.id, "Budi Raharjo", "raharjobd@gmail.com", "123456").Return(true)
			err := userController.PutUserController(c)

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
