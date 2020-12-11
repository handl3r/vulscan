package packages

import "vulscan/src/models"

type AuthenticationPack struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegistrationPack struct {
	Email                string `json:"email"`
	Password             string `json:"password"`
	PasswordConfirmation string `json:"password_confirmation"`
}

type AuthenticationResponsePack struct {
	AccessToken string       `json:"token"`
	User        *models.User `json:"user"`
}
