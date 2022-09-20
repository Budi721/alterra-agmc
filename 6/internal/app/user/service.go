package user

import (
	"github.com/Budi721/alterra-agmc/v6/internal/dto"
	"github.com/Budi721/alterra-agmc/v6/internal/factory"
	"github.com/Budi721/alterra-agmc/v6/internal/pkg/util"
	"github.com/Budi721/alterra-agmc/v6/internal/repository"
	"github.com/Budi721/alterra-agmc/v6/pkg/constant"
	pkgutil "github.com/Budi721/alterra-agmc/v6/pkg/util"
	res "github.com/Budi721/alterra-agmc/v6/pkg/util/response"
	"github.com/pkg/errors"
)

type service struct {
	UserRepository repository.User
}

type Service interface {
	LoginByEmailAndPassword(payload *dto.UserLogin) (string, error)
	CreateUser(payload *dto.User) (*dto.User, error)
	UpdateUser(id uint, payload *dto.UserUpdate) (*dto.UserUpdate, error)
	GetUsers() ([]dto.User, error)
	GetUser(id uint) (*dto.User, error)
	DeleteUser(id uint) (*dto.User, error)
}

func NewService(f *factory.Factory) Service {
	return &service{
		UserRepository: f.UserRepository,
	}
}

func (s service) LoginByEmailAndPassword(payload *dto.UserLogin) (string, error) {
	var result dto.UserLogin
	data, err := s.UserRepository.LoginUser(payload.Email, payload.Password)
	if err != nil {
		if err == constant.RecordNotFound {
			return "", res.ErrorBuilder(&res.ErrorConstant.NotFound, err)
		}
		return "", res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	if !(pkgutil.CompareHashPassword(payload.Password, data.Password)) {
		result.Email = data.Email
		result.Password = data.Password
		return "", res.ErrorBuilder(
			&res.ErrorConstant.EmailOrPasswordIncorrect,
			errors.New(res.ErrorConstant.EmailOrPasswordIncorrect.Response.Meta.Message),
		)
	}

	claims := util.CreateJWTClaims(data.Email, data.ID)
	token, err := util.CreateJWTToken(claims)
	if err != nil {
		return "", res.ErrorBuilder(
			&res.ErrorConstant.InternalServerError,
			errors.New("error when generating token"),
		)
	}

	return token, nil
}

func (s service) CreateUser(payload *dto.User) (*dto.User, error) {
	hashed, err := pkgutil.HashPassword(payload.Password)
	if err != nil {
		return &dto.User{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	result, err := s.UserRepository.CreateUser(payload.Name, payload.Email, hashed)

	return &dto.User{
		ID:       result.ID,
		Name:     result.Name,
		Email:    result.Email,
		Password: result.Password,
	}, nil
}

func (s service) UpdateUser(id uint, payload *dto.UserUpdate) (*dto.UserUpdate, error) {
	result, err := s.UserRepository.UpdateUser(id, payload.Name, payload.Email, payload.Password)
	if err != nil {
		return &dto.UserUpdate{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return &dto.UserUpdate{
		Name:     result.Name,
		Email:    result.Email,
		Password: result.Password,
	}, nil
}

func (s service) GetUsers() ([]dto.User, error) {
	users, err := s.UserRepository.GetUsers()
	if err != nil {
		return []dto.User{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	var result []dto.User
	for _, user := range users {
		result = append(result, dto.User{
			ID:       user.ID,
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
		})
	}

	return result, nil
}

func (s service) GetUser(id uint) (*dto.User, error) {
	user, err := s.UserRepository.GetUser(id)
	if err != nil {
		return &dto.User{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return &dto.User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}

func (s service) DeleteUser(id uint) (*dto.User, error) {
	user, err := s.UserRepository.DeleteUser(id)
	if err != nil {
		return &dto.User{}, res.ErrorBuilder(&res.ErrorConstant.InternalServerError, err)
	}

	return &dto.User{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}
