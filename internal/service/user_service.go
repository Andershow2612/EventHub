package service

import (
	"errors"
	"time"

	"eventHub.com/internal/entity"
	"eventHub.com/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService{
	return &UserService{Repo: &repo}
}

var jwtSecret = []byte("46070d4bf934fb0d4b06d9e2c46e346944e322444900a435d7d9a95e6d7435f5")

func (s *UserService) ListUser() ([]entity.User, error){
	return s.Repo.List()
}

func (s *UserService) ListAUser(id int) (*entity.User, error){
	return s.Repo.ListById(id)
}

func (s *UserService) CreateUser(user *entity.User) (*entity.User, error){
	
	if user.Age < 18 || user.Age > 110{
		return nil, errors.New("user must be over 18")
	}

	_, err := s.Repo.ListByEmail(user.Email)
	if err != nil{
		return nil, errors.New("email already used please try again")
	}
	
	hashedpassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	if err != nil{
		return nil, err
	}

	createdUser := &entity.User{
		Name: user.Name,
		Age: user.Age,
		Email: user.Email,
		Password: string(hashedpassword),
		RoleID: 3,
		ProfileImg: nil,
		CreatedAt: time.Now(),
		Active: 1,
		VerificationCode: 0,

	}

	created, err := s.Repo.Create(createdUser)
	if err != nil{
		return nil, err
	}
	return created, nil
}

func (s *UserService) Login(email, password string) (string, error){
	user, err := s.Repo.ListByEmail(email)
	if err != nil{
		return "", errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil{
		return "", errors.New("invalid password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"role": user.Role.Name,
		"exp": time.Now().Add(2 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil{
		return "", errors.New("fail to generate token")
	}

	return tokenString, nil

}
