package auth

type RegisterRequest struct {
	Name     string `form:"name" binding:"required,min=3,max=123"`
	Email    string `form:"email" binding:"required,email,min=4,max=123"`
	Password string `form:"password" binding:"required,min=4,max=123"`
}
