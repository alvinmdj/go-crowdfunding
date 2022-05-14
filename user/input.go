package user

// RegisterUserInput is the input required for registering a user
type RegisterUserInput struct {
	Name       string `json:"name" binding:"required"`
	Occupation string `json:"occupation" binding:"required"`
	Email      string `json:"email" binding:"required,email"`
	Password   string `json:"password" binding:"required,gte=8"`
}

// LoginUserInput is the input required for logging in a user
type LoginInput struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// CheckEmailInput is the input required for checking if an email is available
type CheckEmailInput struct {
	Email string `json:"email" binding:"required,email"`
}
