package mongodb

import (
	"dime/internal/models"
	"testing"
)

const mongoUri = "mongodb://localhost:27018"

func TestUserDAO_AddUser(t *testing.T) {
	beforeEach(t)

	db := Init(mongoUri)
	defer db.Disconnect()

	userDao, err := db.UserDAO()
	if err != nil {
		t.Errorf("Error getting user dao: %v", err)
	}

	user := models.User{
		FirstName: "John",
		LastName:  "Doe",
		Password:  "password",
	}

	uuid, err := userDao.AddUser(user)
	if err != nil {
		t.Errorf("Error adding user: %v", err)
	}

	if uuid == "" {
		t.Errorf("uuid is empty")
	}
}

func TestUserDAO_GetUser(t *testing.T) {
	beforeEach(t)

	db := Init(mongoUri)
	defer db.Disconnect()

	userDao, err := db.UserDAO()
	if err != nil {
		t.Errorf("Error getting user dao: %v", err)
	}
	/* Set up */
	testUser := models.User{
		FirstName: "John",
		LastName:  "Doe",
		Password:  "password",
	}

	uuid, err := userDao.AddUser(testUser)
	if err != nil {
		t.Errorf("Error adding user: %v", err)
	}
	/* End set up */

	foundUser, err := userDao.GetUser(uuid)
	if err != nil {
		t.Errorf("Error getting user: %v", err)
	}

	if foundUser == nil {
		t.Errorf("User wasn't found")
	}

	if foundUser.FirstName != testUser.FirstName {
		t.Errorf("User first names don't match. Expected %v, got %v", testUser.FirstName, foundUser.FirstName)
	}

	if foundUser.LastName != testUser.LastName {
		t.Errorf("User last names don't match. Expected %v, got %v", testUser.LastName, foundUser.LastName)
	}

	if foundUser.Password != testUser.Password {
		t.Errorf("User passwords don't match. Expected %v, got %v", testUser.Password, foundUser.Password)
	}

}
