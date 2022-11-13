package authdto

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Token     string `json:"token"`
	Image     string `json:"image"`
	Greeting  string `json:"greeting"`
	BestArt   string `json:"bestArt"`
	Following string `json:"following"`
}

type CheckAuthResponse struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Token     string `json:"token"`
	Image     string `json:"image"`
	Greeting  string `json:"greeting"`
	BestArt   string `json:"bestArt"`
	Following string `json:"following"`
}
