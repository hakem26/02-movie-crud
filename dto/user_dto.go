package dto

type UpdateProfileRequest struct {
	FullName string `json:"fullname" validate:"omitempty,min=3"`
	Avatar   string `json:"avatar" validate:"omitempty,url"`
}

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" validate:"required,min=8"`
	NewPassword string `json:"new_password" validate:"required,min=8"`
}