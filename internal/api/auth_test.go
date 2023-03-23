package api

import (
	"dime/internal/models"
	"github.com/golang-jwt/jwt/v4"
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
