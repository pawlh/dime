package api

import (
	"dime/internal/database"
	"dime/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

// secret for jwt signing
// TODO: move this to a config file or make it dynamically generated
var secret = []byte("secret")

type getMeResponse struct {
	UserId    string `json:"userId"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// GetMe Get the currently logged-in user
func GetMe(c echo.Context) error {

	userId := c.Get("user").(*jwt.Token).Claims.(jwt.MapClaims)["userId"].(string)

	userDao, err := database.DB.UserDAO()
	if err != nil {
		return mustSendError(c, http.StatusInternalServerError, "internal server error", err)
	}
	user, err := userDao.GetUser(userId)
	if err != nil {
		return mustSendError(c, http.StatusInternalServerError, "internal server error", err)
	}
	if user == nil {
		// TODO: this should never happen, but probably should be handled better
		return mustSendError(c, http.StatusInternalServerError, "internal server error", err)
	}

	return c.JSON(200, getMeResponse{
		UserId:    user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	})
}

type LoginRequest struct {
	UserId   string `json:"userId"`
	Password string `json:"password"`
}

// Login Log in a user
func Login(c echo.Context) error {
	var loginRequest LoginRequest
	err := c.Bind(&loginRequest)
	if err != nil {
		return mustSendError(c, http.StatusBadRequest, "invalid request body", err)
	}

	if loginRequest.UserId == "" || loginRequest.Password == "" {
		return mustSendError(c, http.StatusBadRequest, "missing userId and/or password", nil)
	}

	userDao, err := database.DB.UserDAO()
	if err != nil {
		return mustSendError(c, http.StatusInternalServerError, "internal server error", err)
	}

	if match, err := userDao.GetUser(loginRequest.UserId); err != nil {
		return mustSendError(c, http.StatusInternalServerError, "internal server error", err)
	} else if match == nil {
		return mustSendError(c, http.StatusUnauthorized, "invalid userId and/or password", nil)
	} else if match.Password != loginRequest.Password {
		return mustSendError(c, http.StatusUnauthorized, "invalid userId and/or password", nil)
	}

	if cookie, err := generateTokenCookie(loginRequest.UserId); err != nil {
		return mustSendError(c, http.StatusInternalServerError, "internal server error", err)
	} else {
		c.SetCookie(cookie)
	}

	return c.JSON(200, nil)
}

func Logout(c echo.Context) error {
	cookie := &http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}
	c.SetCookie(cookie)

	return c.JSON(200, nil)
}

type RegisterRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
}

// Register Register a new user
func Register(c echo.Context) error {
	var registerRequest RegisterRequest
	err := c.Bind(&registerRequest)
	if err != nil {
		return mustSendError(c, http.StatusBadRequest, "invalid request body", err)
	}

	if registerRequest.FirstName == "" || registerRequest.LastName == "" || registerRequest.Password == "" {
		return mustSendError(c, http.StatusBadRequest, "missing firstName, lastName and/or password", nil)
	}

	userDao, err := database.DB.UserDAO()
	if err != nil {
		return mustSendError(c, http.StatusInternalServerError, "internal server error", err)
	}

	id, err := userDao.AddUser(models.User{
		FirstName: registerRequest.FirstName,
		LastName:  registerRequest.LastName,
		Password:  registerRequest.Password,
	})
	if err != nil {
		return mustSendError(c, http.StatusInternalServerError, "internal server error", err)
	}

	if cookie, err := generateTokenCookie(id); err != nil {
		return mustSendError(c, http.StatusInternalServerError, "internal server error", err)
	} else {
		c.SetCookie(cookie)
	}

	return c.JSON(200, nil)
}

// GetUsers Get the name and id of all users
// TODO: add a setting to toggle returning names
func GetUsers(c echo.Context) error {
	userDao, err := database.DB.UserDAO()
	if err != nil {
		return mustSendError(c, http.StatusInternalServerError, "internal server error", err)
	}

	users, err := userDao.GetUsers()
	if err != nil {
		return mustSendError(c, http.StatusInternalServerError, "internal server error", err)
	}

	// TODO: this is a hack, there should probably be a separate struct for this
	for i := range users {
		users[i].Password = ""
	}

	return c.JSON(200, users)
}

func RenewTokenMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token).Claims.(jwt.MapClaims)
		userId := user["userId"].(string)
		if cookie, err := generateTokenCookie(userId); err != nil {
			return mustSendError(c, http.StatusInternalServerError, "internal server error", err)
		} else {
			c.SetCookie(cookie)
		}

		return next(c)
	}
}

type jwtCustomClaims struct {
	Id string `json:"userId"`
	jwt.RegisteredClaims
}

func generateTokenCookie(id string) (*http.Cookie, error) {
	claims := &jwtCustomClaims{
		id,
		jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 4)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secret)
	if err != nil {
		return nil, err
	}

	return &http.Cookie{
		Name:     "token",
		Value:    signedToken,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 4),
		HttpOnly: true,
	}, nil
}
