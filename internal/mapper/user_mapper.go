package mapper

import (
	"time"

	dto "eventHub.com/internal/DTO"
	"eventHub.com/internal/entity"
)

//user response
func ToUserResponse(user *entity.User) dto.UserResponse {
    return dto.UserResponse{
        ID:         user.ID,
        Name:       user.Name,
        Age:        user.Age,
        Email:      user.Email,
        Created_at: user.CreatedAt,
        Updated_at: user.UpdatedAt,
        Deleted_at: user.DeletedAt,
        Active:     user.Active,
        Role:       ToRoleResponse(user.Role),
    }
}

func ToUserResponseList(users []entity.User) []dto.UserResponse{
	var responses []dto.UserResponse
	
	for _, user := range users {
		responses = append(responses, ToUserResponse(&user))
	}
	return responses
}

func ToUserCreated(user *entity.User) dto.UserCreatedResponse{
	return dto.UserCreatedResponse{
		Name: user.Name,
		Email: user.Email,
		Age: user.Age,
	}
}

//user request
func ToUserEntity(req dto.UserReq) *entity.User{
	return &entity.User{
		Name: req.Name,
		Email: req.Email,
		Password: req.Password,
		Age: req.Age,
		RoleID: 3,
		ProfileImg: nil,
		Active: 1,
		CreatedAt: time.Now(),
	}
}

func ToUserUpdateEntity(req dto.UserUpdateReq) *entity.User {
    return &entity.User{
        Name:  req.UserName,
        Age:   req.Age,
        Email: req.Email,
    }
}