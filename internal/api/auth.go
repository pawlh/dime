package api

import (
	"crypto/rand"
	"dime/internal/database"
	"dime/internal/models"
	"github.com/labstack/echo/v4"
	"net/http"
)

type getMeResponse struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func GetMe(c echo.Context) error {
	user := models.User{
		FirstName: "John",
		LastName:  "Doe",
	}

	return c.JSON(200, getMeResponse{
		FirstName: user.FirstName,
		LastName:  user.LastName,
	})
}

type LoginRequest struct {
	Id       string `json:"id"`
	Password string `json:"password"`
}

func Login(c echo.Context) error {
	var loginRequest LoginRequest
	err := c.Bind(&loginRequest)
	if err != nil {
		return mustSendError(c, http.StatusBadRequest, "invalid request body", err)
	}

	if loginRequest.Id == "" || loginRequest.Password == "" {
		return mustSendError(c, http.StatusBadRequest, "missing userId and/or password", nil)
	}

	userDao, err := database.DB.UserDAO()
	if err != nil {
		return mustSendError(c, http.StatusInternalServerError, "internal server error", err)
	}

	if match, err := userDao.GetUser(loginRequest.Id); err != nil {
		return mustSendError(c, http.StatusInternalServerError, "internal server error", err)
	} else if match == nil {
		return mustSendError(c, http.StatusUnauthorized, "invalid userId and/or password", nil)
	} else if match.Password != loginRequest.Password {
		return mustSendError(c, http.StatusUnauthorized, "invalid userId and/or password", nil)
	}

	if cookie, err := generateCookie(); err != nil {
		return mustSendError(c, http.StatusInternalServerError, "internal server error", err)
	} else {
		c.SetCookie(cookie)
	}

	return c.JSON(200, nil)
}

type RegisterRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Password  string `json:"password"`
}

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

	if _, err := userDao.AddUser(models.User{
		FirstName: registerRequest.FirstName,
		LastName:  registerRequest.LastName,
		Password:  registerRequest.Password,
	}); err != nil {
		return mustSendError(c, http.StatusInternalServerError, "internal server error", err)
	}

	return c.JSON(200, nil)
}

func GetUsers(c echo.Context) error {
	userDao, err := database.DB.UserDAO()
	if err != nil {
		return mustSendError(c, http.StatusInternalServerError, "internal server error", err)
	}

	users, err := userDao.GetUsers()
	if err != nil {
		return mustSendError(c, http.StatusInternalServerError, "internal server error", err)
	}

	// strip out passwords
	// TODO: this is a hack, there should probably be a separate struct for this
	for i := range users {
		users[i].Password = ""
	}

	return c.JSON(200, users)
}

func generateCookie() (*http.Cookie, error) {
	token, err := generateRandomBytes(32)
	if err != nil {
		return nil, err
	}

	return &http.Cookie{
		Name:  "token",
		Value: string(token),
	}, nil
}

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}
