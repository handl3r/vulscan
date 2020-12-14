package services

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
	"vulscan/src/adapter/repositories"
	"vulscan/src/enums"
	"vulscan/src/packages"
)

// Login, registration, validate permission to read data. ex: user query resource belongs to another user

type AuthenticationService struct {
	userRepository *repositories.UserRepository
	accessTokenExp time.Duration
	authSecretKey  string
}

func NewAuthenticationService(
	userRepository *repositories.UserRepository, accessTokenExpTime time.Duration, authSecretKey string,
) *AuthenticationService {
	return &AuthenticationService{
		userRepository: userRepository,
		accessTokenExp: accessTokenExpTime,
		authSecretKey:  authSecretKey,
	}
}

func (a *AuthenticationService) Authenticate(authPack packages.AuthenticationPack) (*packages.AuthenticationResponsePack, enums.Error) {
	user, err := a.userRepository.FindByEmail(authPack.Email)
	if err == enums.ErrEntityNotFound {
		return nil, enums.ErrUnauthenticated
	}
	if err != nil {
		log.Printf("Can not find user by email with error %s", err)
		return nil, enums.ErrSystem
	}
	if !a.validatePassword(user.Password, authPack.Password) {
		return nil, enums.ErrUnauthenticated
	}
	tokenInfoPack := packages.TokenInformationPack{
		UserID:  user.ID,
		ExpTime: a.accessTokenExp,
	}
	accessToken, err := a.generateAccessToken(tokenInfoPack)
	if err != nil {
		log.Printf("Can not generate access token for user %s with error %s", user.ID, err)
		return nil, enums.ErrSystem
	}
	return &packages.AuthenticationResponsePack{
		AccessToken: accessToken,
		User:        user,
	}, nil
}

func (a *AuthenticationService) validatePassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func (a *AuthenticationService) generateAccessToken(tokenInfoPack packages.TokenInformationPack) (string, error) {
	claims := jwt.MapClaims{
		"uid": tokenInfoPack.UserID,
		"exp":    time.Now().Add(tokenInfoPack.ExpTime).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(a.authSecretKey))
}
