package api

import (
	"dime/internal/models"
	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"net/http"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	user := models.User{
		Username: "John",
		Name:     "Henry",
	}

	token, err := generateToken(user)
	if err != nil {
		t.Fatal(err)
	}

	claims, _ := jwt.ParseWithClaims(token, &jwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if valid := claims.Valid; !valid {
		t.Fatalf(`token.Valid != true, got: %t`, valid)
	}

	customClaim := claims.Claims.(*jwtCustomClaims)
	if customClaim.Username != user.Username {
		t.Fatal(`token.Username != user.Username, got: ` + customClaim.Username + `, want: ` + user.Username)
	}
	if customClaim.Name != user.Name {
		t.Fatal(`token.Name != user.Name, got: ` + customClaim.Name + `, want: ` + user.Name)
	}
}

func TestLogin(t *testing.T) {
	BeforeEach()
	InsertTestUser()

	tests := []struct {
		name         string
		body         *models.User
		expectedCode int
	}{
		{
			name: "Missing Username",
			body: &models.User{
				Password: "password",
				Name:     "name",
			},
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "Missing Password",
			body: &models.User{
				Username: "username",
				Name:     "name",
			},
			expectedCode: http.StatusBadRequest,
		},
		{
			"Valid Login",
			testUser,
			http.StatusOK,
		},
		{
			"Password is incorrect",
			&models.User{
				Username: testUser.Username,
				Password: "fake",
			},
			http.StatusUnauthorized,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, resBody := ServeAndRequest(echo.GET, "/api/login", Login, tt.body)

			if res.Code != tt.expectedCode {
				t.Fatalf("expected status code %d, got %d with message \"%s\"", tt.expectedCode, res.Code, resBody["error"])
			}

			if tt.expectedCode == http.StatusOK {
				if username, _ := resBody["username"]; username != tt.body.Username {
					t.Fatal("username not found in response")
				}
				if name, _ := resBody["name"]; name != tt.body.Name {
					t.Fatal("name not found in response")
				}
			}
		})
	}
}

func TestRegister(t *testing.T) {
	BeforeEach()

	tests := []struct {
		name         string
		body         *models.User
		expectedCode int
	}{
		{
			name: "Missing Username",
			body: &models.User{
				Password: "password",
				Name:     "name",
			},
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "Missing Password",
			body: &models.User{
				Username: "username",
				Name:     "name",
			},
			expectedCode: http.StatusBadRequest,
		},
		{
			name: "Missing Name",
			body: &models.User{
				Username: "username",
				Password: "password",
			},
			expectedCode: http.StatusBadRequest,
		},
		{
			"Valid Register",
			testUser,
			http.StatusOK,
		},
		{
			"User Already Exists",
			testUser,
			http.StatusConflict,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, resBody := ServeAndRequest(echo.GET, "/api/register", Register, tt.body)

			if res.Code != tt.expectedCode {
				t.Fatalf("expected status code %d, got %d with message \"%s\"", tt.expectedCode, res.Code, resBody["error"])
			}

			if tt.expectedCode == http.StatusOK && resBody["token"] == "" {
				t.Fatal("expected token, got empty string")
			}

		})
	}
}
