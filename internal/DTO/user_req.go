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