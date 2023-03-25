package mongodb

import (
	"dime/internal/models"
	"reflect"
	"testing"
)

func TestUserDao_Create_Find(t *testing.T) {
	testUser := models.User{
		Username: "testUsername",
		Password: "testPassword",
		Name:     "testName",
	}

	userDao := NewUser(client)
	err := userDao.Create(&testUser)
	if err != nil {
		t.Errorf("Error creating a new user: %v", err)
	}

	if match, err := userDao.FindByUsername(testUser.Username); err != nil {
		t.Errorf("Error finding user: %v", err)
	} else if !reflect.DeepEqual(testUser, *match) {
		t.Errorf("Users do not match: %v", err)
	}
}

func TestUserDoesNotExist(t *testing.T) {
	userDao := NewUser(client)
	if match, err := userDao.FindByUsername("nonexistent"); err != nil {
		t.Errorf("Error finding user: %v", err)
	} else if match != nil {
		t.Errorf("User should not exist: %v", err)
	}
}
