package db

import (
	"math/rand"
	model "shanyraq/models"
	"testing"
	"time"
)

func TestDbConnection(t *testing.T) {
	_, _ = MongoDbConnect()

}

func TestSaveUser(t *testing.T) {
	testUser := &model.User{Username: "TestUsername", Email: "TestEmail", Name: "TestName", Surname: "TestSurname", Telephone: "87776665544", Password: "TestPassword"}

	_, err := SaveUser(*testUser)

	if err != nil {
		if err.Error() != "username is not unique" && err.Error() != "email is not unique" && err.Error() != "telephone is not unique" {
			t.Errorf("Failed creating user: %s", err)
		}
	}

	deleted := DeleteUserByUsername(testUser.Username)

	if !deleted {
		t.Errorf("User is not deleted!")
	}
}

func TestIsUniqueUser(t *testing.T) {

	testUser := RandomUser()
	is_unique, err := IsUniqueUser(*testUser)

	if !is_unique {
		t.Errorf("user should be unique: %s", err)
	}

	_, err = SaveUser(*testUser)

	if err != nil {
		t.Errorf("Failed creating user: %s", err)
	}

	is_unique, err = IsUniqueUser(*testUser)

	if is_unique {
		t.Errorf("user shouldn't be unique: %s", err)
	}

	deleted := DeleteUserByUsername(testUser.Username)

	if !deleted {
		t.Errorf("User is not deleted!")
	}
}

func TestGetUser(t *testing.T) {

	testUser := RandomUser()
	_, err := SaveUser(*testUser)

	if err != nil {
		if err.Error() != "username is not unique" && err.Error() != "email is not unique" && err.Error() != "telephone is not unique" {
			t.Errorf("Failed creating user: %s", err)
		}
	}

	retrievedUser := GetUserByUsername(testUser.Username)

	if retrievedUser.Username != testUser.Username {
		t.Errorf("Usernames should be same, testUsername is %s while retrieved on is %s", testUser.Username, retrievedUser.Username)
	}

	anotherUser := GetUserById(retrievedUser.ID)

	if anotherUser.Username != testUser.Username {
		t.Errorf("Smth wrong with GetUserById. Usernames should be same, testUsername is %s while retrieved on is %s", testUser.Username, anotherUser.Username)
	}

	deleted := DeleteUserByUsername(testUser.Username)

	if !deleted {
		t.Errorf("User is not deleted!")
	}

}

func TestUpdateUser(t *testing.T) {
	testUser := RandomUser()

	_, err := SaveUser(*testUser)

	if err != nil {
		if err.Error() != "username is not unique" && err.Error() != "email is not unique" && err.Error() != "telephone is not unique" {
			t.Errorf("Failed creating user: %s", err)
		}
	}

	testUser = GetUserByUsername(testUser.Username)

	testUser.Name = "NewTestName"

	err = UpdateUser(*testUser)

	if err != nil {
		t.Errorf("Failed updating user: %s", err)
	}

	dbUser := GetUserByUsername(testUser.Username)
	if dbUser.Name != "NewTestName" {
		t.Errorf("Failed updating username. %s should be %s", dbUser.Name, "NewTestName")
	}
	deleted := DeleteUserByUsername(testUser.Username)

	if !deleted {
		t.Errorf("User is not deleted!")
	}

}

func TestIsValidCredentials(t *testing.T) {
	testUser := RandomUser()

	_, err := SaveUser(*testUser)

	if err != nil {
		if err.Error() != "username is not unique" && err.Error() != "email is not unique" && err.Error() != "telephone is not unique" {
			t.Errorf("Failed creating user: %s", err)
		}
	}

	is_valid := IsValidCredentials(*testUser)

	if !is_valid {
		t.Errorf("Credentials should be valid!")
	}

	testUser.Password = "some bullshit"

	is_valid = IsValidCredentials(*testUser)

	if is_valid {
		t.Errorf("Credentials shouldn't be valid!")
	}

	deleted := DeleteUserByUsername(testUser.Username)

	if !deleted {
		t.Errorf("User is not deleted!")
	}
}

func TestValidateUser(t *testing.T) {
	testUser := RandomUser()

	_, err := SaveUser(*testUser)
	if err != nil {
		if err.Error() != "username is not unique" && err.Error() != "email is not unique" && err.Error() != "telephone is not unique" {
			t.Errorf("Failed creating user: %s", err)
		}
	}

	err = ValidateUser(testUser.Username)

	if err != nil {
		t.Errorf("VALIDATING USER FAIL: %s", err.Error())
	}

	dbUser := GetUserByUsername(testUser.Username)
	if !dbUser.IsValidated {
		t.Errorf("User should be validated! ")
	}

	deleted := DeleteUserByUsername(testUser.Username)

	if !deleted {
		t.Errorf("User is not deleted!")
	}
}

func RandomUser() *model.User {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var seededRand *rand.Rand = rand.New(
		rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, 5)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	RandomUsername := string(b) + "Test"
	randomEmail := string(b) + "Email"
	randomTel := string(b) + "Tel"
	testUser := &model.User{Username: RandomUsername, Email: randomEmail, Telephone: randomTel, Password: "TestPassword", Name: "TestName", Surname: "TestSurname"}

	return testUser
}
