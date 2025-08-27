package dto

type UserReq struct {
	Name     string `json:"UserName" binding:"required"`
	Age      int    `json:"Age" binding:"required"`
	Email    string `json:"Email" binding:"required"`
	Password string `json:"Password" binding:"required"`
}

type LoginReq struct {
	Email    string `json:"Email" binding:"required"`
	Password string `json:"Password" binding:"required"`
}

type UserUpdateReq struct {
	UserName string `json:"UserName"`
	Age      int    `json:"Age"`
	Email    string `json:"Email"`
}

type UserUpdatePassword struct {
	OldPassword string `json:"Old_Password"`
	NewPassword string `json:"New_Password"`
}