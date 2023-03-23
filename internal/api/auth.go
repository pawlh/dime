package api

import (
	"dime/internal/dbs"
	"dime/internal/models"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

// Secret TODO: move this to a config file
var secret = []byte("secret")

type jwtCustomClaims struct {
	Username string `json:"usr"`
	Name     string `json:"nme"`
	jwt.RegisteredClaims
}

func Login(c echo.Context) error {
	var user models.User

	if err := c.Bind(&user); err != nil {
		return mustSendError(c, http.StatusBadRequest, "missing username and/or password")
	}

	if match, err := dbs.DB.UserDao().FindByUsername(user.Username); err != nil {
		return mustSendError(c, http.StatusInternalServerError, "error finding user")
	} else if match.Password != user.Password {
		return mustSendError(c, http.StatusUnauthorized, "bad credentials")
	}

	claims := &jwtCustomClaims{
		user.Username,
		user.Name,
		jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secret)
	if err != nil {
		return mustSendError(c, http.StatusInternalServerError, "error signing token")
	}

	if err = c.JSON(http.StatusOK, echo.Map{
		"token": signedToken,
	}); err != nil {
		return mustSendError(c, http.StatusInternalServerError, "error sending token")
	}

	return nil
}
