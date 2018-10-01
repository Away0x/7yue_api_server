package user

type UserBody struct {
	UserName string `json:"username" binding:"required"`
}