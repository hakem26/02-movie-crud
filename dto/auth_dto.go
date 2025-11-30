package dto

type RegisterRequest struct {
	FullName string       `json:"fullname" validate:"required,min=3"`
	Email    string       `json:"email" validate:"required,email"`
	Password string       `json:"password" validate:"required,min=8"`
	Level    LevelRequest `json:"level" validate:"required"` // فقط required کافیه!
}

type LevelRequest struct {
	LevelID   string `json:"level_id" validate:"required"`
	LevelName string `json:"level_name" validate:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type TokenResponse struct {
	// این رو بعداً از utils می‌آریم، ولی فعلاً خالی باشه
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
}