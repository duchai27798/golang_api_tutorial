package service

import (
	"github.com/duchai27798/golang_api_tutorial/src/data/dto"
	"github.com/duchai27798/golang_api_tutorial/src/data/entity"
	"github.com/duchai27798/golang_api_tutorial/src/data/repository"
	"github.com/duchai27798/golang_api_tutorial/src/utils"
	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
)

type IAuthService interface {
	VerifyCredential(email string, password string) interface{}
	CreateUser(user dto.RegisterDTO) entity.User
	FindByEmail(email string) entity.User
	IsDuplicateEmail(email string) bool
}

type AuthService struct {
	userRepository repository.IUserRepository
}

func (authService AuthService) VerifyCredential(email string, password string) interface{} {
	res := authService.userRepository.VerifyCredential(email, password)
	if v, ok := res.(entity.User); ok {
		if v.Email == email && comparePassword(v.Password, []byte(password)) {
			return res
		}
	}
	return false
}

func (authService AuthService) CreateUser(user dto.RegisterDTO) entity.User {
	u := entity.User{}
	err := smapping.FillStruct(&u, smapping.MapFields(&user))
	if err != nil {
		utils.LogError(err)
		return u
	}
	return authService.userRepository.InsertUser(u)
}

func (authService AuthService) FindByEmail(email string) entity.User {
	return authService.userRepository.FindByEmail(email)
}

func (authService AuthService) IsDuplicateEmail(email string) bool {
	res := authService.userRepository.IsDuplicateEmail(email)
	utils.LogObj(res.Error)
	return !(res.Error == nil)
}

func NewAuthService(userRepository repository.IUserRepository) IAuthService {
	return &AuthService{
		userRepository,
	}
}

func comparePassword(hashedPwd string, plainPwd []byte) bool {
	bytePwd := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(bytePwd, plainPwd)
	if err != nil {
		utils.LogError(err)
		return false
	}
	return true
}
