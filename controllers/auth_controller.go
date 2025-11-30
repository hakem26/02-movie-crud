package controllers

import (
	"encoding/json"
	"net/http"

	"example/moviecrud/dto"
	"example/moviecrud/services"
	"example/moviecrud/utils"

	"github.com/go-playground/validator/v10"
)

type AuthController struct {
	authService *services.AuthService
	validator   *validator.Validate
}

func NewAuthController(authSrv *services.AuthService) *AuthController {
	return &AuthController{
		authService: authSrv,
		validator:   validator.New(),
	}
}

func (c *AuthController) Register(w http.ResponseWriter, r *http.Request) {
	var req dto.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.JSONError(w, "invalid JSON", http.StatusBadRequest)
		return
	}

	// این خط مهم بود: validate بعد از decode
	if err := c.validator.Struct(req); err != nil {
		// برگردوندن خطای خوانا
		utils.JSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := c.authService.Register(&req)
	if err != nil {
		if err.Error() == "email already registered" {
			utils.JSONError(w, "email already in use", http.StatusConflict)
			return
		}
		utils.JSONError(w, "registration failed", http.StatusInternalServerError)
		return
	}

	utils.JSONResponse(w, user, http.StatusCreated)
}

func (c *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.JSONError(w, "invalid request", http.StatusBadRequest)
		return
	}

	if err := c.validator.Struct(req); err != nil {
		utils.JSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	tokenRes, err := c.authService.Login(req.Email, req.Password)
	if err != nil {
		utils.JSONError(w, "invalid email or password", http.StatusUnauthorized)
		return
	}

	utils.JSONResponse(w, tokenRes, http.StatusOK)
}

func (c *AuthController) RefreshToken(w http.ResponseWriter, r *http.Request) {
	// بعداً پیاده‌سازی می‌کنیم (با ذخیره refresh token در دیتابیس یا redis)
	utils.JSONError(w, "not implemented yet", http.StatusNotImplemented)
}