package mapper

import (
	"time"

	dto "eventHub.com/internal/DTO"
	"eventHub.com/internal/entity"
)

//user response
func ToUserResponse(user *entity.User) dto.UserResponse{
	return dto.UserResponse{
		ID: user.ID,
		Name: user.Name,
		Age: user.Age,
		Email: user.Email,
		Created_at: user.Created_at,
		Updated_at: user.Updated_at,
		Deleted_at: user.Deleted_at,
		Active: user.Active,
	}
}

func ToUserResponseList(users []entity.User) []dto.UserResponse{
	var responses []dto.UserResponse
	
	for _, user := range users {
		responses = append(responses, ToUserResponse(&user))
	}
	return responses
}

//user request
func ToUserEntity(req dto.UserReq) entity.User{
	return entity.User{
		Name: req.Name,
		Email: req.Email,
		Password: req.Password,
		Age: req.Age,
		Role_id: 1,
		Profile_img: nil,
		Active: 1,
		Created_at: time.Now(),
	}
}