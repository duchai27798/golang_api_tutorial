package repository

import (
	"github.com/duchai27798/golang_api_tutorial/src/data/entity"
	"github.com/duchai27798/golang_api_tutorial/src/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type IUserRepository interface {
	InsertUser(user entity.User) entity.User
	UpdateUser(user entity.User) entity.User
	VerifyCredential(email string, password string) interface{}
	FindByEmail(email string) entity.User
	IsDuplicateEmail(email string) (tx *gorm.DB)
	FindByID(userID string) entity.User
}

type UserRepository struct {
	connection *gorm.DB
}

func (userRepository UserRepository) InsertUser(user entity.User) entity.User {
	user.Password = hashAndSalt([]byte(user.Password))
	userRepository.connection.Save(&user)
	return user
}

func (userRepository UserRepository) UpdateUser(user entity.User) entity.User {
	if user.Password == "" {
		user.Password = hashAndSalt([]byte(user.Password))
	} else {
		var userTmp entity.User
		userRepository.connection.Find(&userTmp, user.ID)
		user.Password = userTmp.Password
	}
	userRepository.connection.Save(&user)
	return user
}

func (userRepository UserRepository) VerifyCredential(email string, password string) interface{} {
	var user entity.User
	if err := userRepository.connection.Where("email = ?", email).First(&user).Error; err != nil {
		utils.LogError(err)
	}
	return user
}

func (userRepository UserRepository) FindByEmail(email string) entity.User {
	var user entity.User
	userRepository.connection.Where("email = ?", email).Take(&user)
	return user
}

func (userRepository UserRepository) FindByID(userID string) entity.User {
	var user entity.User
	userRepository.connection.Find(&user, userID)
	return user
}

func (userRepository UserRepository) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user entity.User
	return userRepository.connection.Where("email = ?", email).Take(&user)
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		connection: db,
	}
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		utils.LogError(err)
	}
	return string(hash)
}
