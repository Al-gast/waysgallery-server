package userdto

type CreateUserRequest struct {
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UpdateUserRequest struct {
	Name      string `json:"name"`
	Image     string `json:"image"`
	Greeting  string `json:"greeting"`
	BestArt   string `json:"bestArt"`
	Following string `json:"following"`
}

type UserResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email" `
	Password  string `json:"password"`
	Image     string `json:"image" `
	Greeting  string `json:"greeting"`
	BestArt   string `json:"bestArt" `
	Following string `json:"following"`
}
