package api

import (
	"dime/internal/dao"
	"dime/internal/dbs"
	"dime/internal/models"
	"errors"
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
		return mustSendError(c, http.StatusBadRequest, "missing username and/or password", nil)
	}

	if match, err := dbs.DB.UserDao().FindByUsername(user.Username); err != nil {
		return mustSendError(c, http.StatusInternalServerError, "error finding user", nil)
	} else if match.Password != user.Password {
		return mustSendError(c, http.StatusUnauthorized, "bad credentials", nil)
	}

	token, err := generateToken(user)
	if err != nil {
		return mustSendError(c, http.StatusInternalServerError, "error generating token", nil)
	}

	if err != nil {
		return mustSendError(c, http.StatusInternalServerError, "error signing token", nil)
	}

	if err = c.JSON(http.StatusOK, echo.Map{
		"token": token,
	}); err != nil {
		return mustSendError(c, http.StatusInternalServerError, "error sending token", nil)
	}

	return nil
}

func Register(c echo.Context) error {
	var user models.User

	if err := c.Bind(&user); err != nil {
		return mustSendError(c, http.StatusBadRequest, "missing username and/or password", err)
	}

	if user.Name == "" || user.Username == "" || user.Password == "" {
		return mustSendError(c, http.StatusBadRequest, "missing required fields", nil)
	}

	if match, err := dbs.DB.UserDao().FindByUsername(user.Username); err != nil {
		if !errors.As(err, &dao.UserNotFound{}) {
			return mustSendError(c, http.StatusInternalServerError, "error searching existing users", err)
		}
	} else if match.Username != "" {
		return mustSendError(c, http.StatusConflict, "username already exists", nil)
	}

	if err := dbs.DB.UserDao().Create(&user); err != nil {
		return mustSendError(c, http.StatusInternalServerError, "error creating user", err)
	}

	token, err := generateToken(user)
	if err != nil {
		return mustSendError(c, http.StatusInternalServerError, "error generating token", err)
	}

	if err := c.JSON(http.StatusOK, echo.Map{
		"token": token,
	}); err != nil {
		return mustSendError(c, http.StatusInternalServerError, "error sending message", err)
	}

	return nil
}

func generateToken(user models.User) (string, error) {
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

	return signedToken, err
}

func validateToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Get("user").(*jwt.Token)
		claims := token.Claims.(jwt.MapClaims)

		if claims.Valid() != nil {
			return mustSendError(c, http.StatusUnauthorized, "bad token", nil)
		}

		if !claims.VerifyExpiresAt(time.Now().Add(time.Hour).Unix(), true) {
			return mustSendError(c, http.StatusUnauthorized, "token expired", nil)
		}

		c.Set("username", claims["usr"])

		return next(c)
	}
}
