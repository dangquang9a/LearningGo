package modeltests

import (
	"log"
	"testing"
	"time"

	"github.com/dangquang9a/LearningGo/api/models"
	"github.com/dangquang9a/LearningGo/api/utils/randomizer"
	"gopkg.in/go-playground/assert.v1"
)

func TestFindAllUsers(t *testing.T) {

	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}

	err = seedUsers()
	if err != nil {
		log.Fatal(err)
	}

	users, err := userInstance.FindAllUsers(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the users: %v\n", err)
		return
	}

	assert.Equal(t, len(*users), 2)

}

func TestSaveUser(t *testing.T) {

	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}
	newUser := models.User{
		ID:       1,
		Nickname: randomizer.GenerateName(),
		Email:    randomizer.GenerateEmail(),
		Password: "password",
	}
	savedUser, err := newUser.SaveUser(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the users: %v\n", err)
		return
	}

	assert.Equal(t, newUser.ID, savedUser.ID)
	assert.Equal(t, newUser.Email, savedUser.Email)
	assert.Equal(t, newUser.Nickname, savedUser.Nickname)

}

func TestGetUserByID(t *testing.T) {

	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}

	user, err := seedOneUser()
	if err != nil {
		log.Fatalf("cannot seed users table :%v", err)
	}

	foundUser, err := userInstance.FindUserByID(server.DB, user.ID)
	if err != nil {
		t.Errorf("this is the error getting the users: %v\n", err)
		return
	}
	assert.Equal(t, foundUser.ID, user.ID)
	assert.Equal(t, foundUser.Email, user.Email)
	assert.Equal(t, foundUser.Nickname, user.Nickname)

}

func TestUpdateUser(t *testing.T) {

	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}

	user, err := seedOneUser()
	if err != nil {
		log.Fatalf("cannot seed users table :%v", err)
	}

	userUpdate := models.User{
		ID:        1,
		Nickname:  randomizer.GenerateName(),
		Email:     randomizer.GenerateEmail(),
		Password:  "password",
		UpdatedAt: time.Time{},
	}

	updatedUser, err := userUpdate.UpdateUser(server.DB, user.ID)
	if err != nil {
		t.Errorf("this is the error updating the users: %v\n", err)
		return
	}
	assert.Equal(t, updatedUser.ID, userUpdate.ID)
	assert.Equal(t, updatedUser.Email, userUpdate.Email)
	assert.Equal(t, updatedUser.Nickname, userUpdate.Nickname)
}

func TestDeleteUser(t *testing.T) {

	err := refreshUserTable()
	if err != nil {
		log.Fatal(err)
	}

	user, err := seedOneUser()
	if err != nil {
		log.Fatalf("cannot seed users table :%v", err)
	}

	isDeleted, err := userInstance.DelteUser(server.DB, user.ID)
	if err != nil {
		t.Errorf("this is the error deleting the users: %v\n", err)
		return
	}

	assert.Equal(t, int(isDeleted), 1)

}
