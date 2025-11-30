package controllers

import (
	"encoding/json"
	"net/http"

	"example/moviecrud/middleware"
	"example/moviecrud/models"
	"example/moviecrud/services"
	"example/moviecrud/utils"

	"github.com/gorilla/mux"
)

type UserController struct {
	Service *services.UserService
}

func NewUserController(s *services.UserService) *UserController {
	return &UserController{Service: s}
}

// تبدیل models.User به PublicUser (بدون پسورد و با id درست)
func toPublicUser(u *models.User) models.PublicUser {
	return models.PublicUser{
		ID:       u.UserID, // این مهمه! از UserID استفاده می‌کنه نه ID
		FullName: u.FullName,
		Email:    u.Email,
		Level:    u.Level,
	}
}

func toPublicUsers(users []*models.User) []models.PublicUser {
	result := make([]models.PublicUser, len(users))
	for i, u := range users {
		result[i] = toPublicUser(u)
	}
	return result
}

func (c *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := c.Service.GetAll()
	if err != nil {
		utils.JSONError(w, "failed to fetch users", http.StatusInternalServerError)
		return
	}

	publicUsers := toPublicUsers(users)
	utils.JSONResponse(w, publicUsers, http.StatusOK)
}

func (c *UserController) GetUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	user, err := c.Service.GetByID(id)
	if err != nil || user == nil {
		utils.JSONError(w, "user not found", http.StatusNotFound)
		return
	}

	publicUser := toPublicUser(user)
	utils.JSONResponse(w, publicUser, http.StatusOK)
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	// این رو بعداً غیرفعال می‌کنیم چون فقط از /auth/register باید ثبت نام کنه
	utils.JSONError(w, "use /auth/register endpoint", http.StatusForbidden)
}

func (c *UserController) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]

	var input struct {
		FullName string           `json:"fullname" validate:"omitempty,min=3"`
		Level    models.UserLevel `json:"level" validate:"omitempty,dive"`
	}

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.JSONError(w, "invalid body", http.StatusBadRequest)
		return
	}

	// اینجا باید از سرویس آپدیت استفاده کنی (بعداً پیاده می‌کنیم)
	user, err := c.Service.GetByID(id)
	if err != nil || user == nil {
		utils.JSONError(w, "not found", http.StatusNotFound)
		return
	}

	if input.FullName != "" {
		user.FullName = input.FullName
	}
	if input.Level.LevelID != "" {
		user.Level = input.Level
	}

	updatedUser, err := c.Service.Update(id, user)
	if err != nil {
		utils.JSONError(w, err.Error(), http.StatusBadRequest)
		return
	}

	utils.JSONResponse(w, toPublicUser(updatedUser), http.StatusOK)
}

func (c *UserController) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if err := c.Service.Delete(id); err != nil {
		utils.JSONError(w, "not found or error", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (c *UserController) GetMe(w http.ResponseWriter, r *http.Request) {
	claims := middleware.GetCurrentUser(r)
	if claims == nil {
		utils.JSONError(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	user, err := c.Service.GetByUserID(claims.UserID)
	if err != nil || user == nil {
		utils.JSONError(w, "user not found", http.StatusNotFound)
		return
	}

	utils.JSONResponse(w, toPublicUser(user), http.StatusOK)
}

func (c *UserController) UpdateMe(w http.ResponseWriter, r *http.Request) {
	claims := middleware.GetCurrentUser(r)
	if claims == nil {
		utils.JSONError(w, "unauthorized", http.StatusUnauthorized)
		return
	}

	var input struct {
		FullName string `json:"fullname" validate:"omitempty,min=3"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		utils.JSONError(w, "invalid body", http.StatusBadRequest)
		return
	}

	user, _ := c.Service.GetByUserID(claims.UserID)
	if user.FullName = input.FullName; input.FullName != "" {
		updated, err := c.Service.Update(claims.UserID, user)
		if err != nil {
			utils.JSONError(w, "update failed", http.StatusBadRequest)
			return
		}

		utils.JSONResponse(w, toPublicUser(updated), http.StatusOK)
	}
}
