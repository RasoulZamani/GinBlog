package auth

type LoginRequest struct {
	Email    string `form:"email" binding:"required,email,min=4,max=123"`
	Password string `form:"password" binding:"required,min=4,max=123"`
}
