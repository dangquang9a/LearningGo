package modeltests

import (
	"log"
	"testing"

	"github.com/dangquang9a/LearningGo/api/models"
	"github.com/dangquang9a/LearningGo/api/utils/randomizer"
	"gopkg.in/go-playground/assert.v1"
)

func TestFindAllPost(t *testing.T) {

	err := refreshUserAndPostTable()
	if err != nil {
		log.Fatalf("Error refreshing user and post table %v\n", err)
	}

	_, _, err = seedUsersAndPosts()
	if err != nil {
		log.Fatalf("Error seeding user and post table %v\n", err)
	}

	posts, err := postInstance.FindAllPosts(server.DB)
	if err != nil {
		t.Errorf("this is the error getting the posts: %v\n", err)
		return
	}
	assert.Equal(t, len(*posts), 2)

}

func TestSavePost(t *testing.T) {

	err := refreshUserAndPostTable()
	if err != nil {
		log.Fatalf("Error refreshing user and post table %v\n", err)
	}

	user, err := seedOneUser()
	if err != nil {
		log.Fatalf("Cannot seed user %v\n", err)
	}

	newPost := models.Post{
		ID:       1,
		Title:    "This is the title",
		Content:  "This is the content",
		AuthorID: user.ID,
	}
	savedPost, err := newPost.SavePost(server.DB)
	if err != nil {
		t.Errorf("this is the error saving the post: %v\n", err)
	}
	assert.Equal(t, newPost.ID, savedPost.ID)
	assert.Equal(t, newPost.Title, savedPost.Title)
	assert.Equal(t, newPost.Content, savedPost.Content)
	assert.Equal(t, newPost.AuthorID, savedPost.AuthorID)

}

func TestGetPostByID(t *testing.T) {

	err := refreshUserAndPostTable()
	if err != nil {
		log.Fatalf("Error refreshing user and post table %v\n", err)
	}

	post, err := seedOneUserAndOnePost()
	if err != nil {
		log.Fatalf("Error seeding user and post table %v\n", err)
	}

	foundPost, err := postInstance.FindPostByID(server.DB, post.ID)
	if err != nil {
		t.Errorf("this is the error finding the post: %v\n", err)
		return
	}

	assert.Equal(t, foundPost.ID, post.ID)
	assert.Equal(t, foundPost.Title, post.Title)
	assert.Equal(t, foundPost.Content, post.Content)

}

func TestUpdatePost(t *testing.T) {

	err := refreshUserAndPostTable()
	if err != nil {
		log.Fatalf("Error refreshing user and post table %v\n", err)
	}

	post, err := seedOneUserAndOnePost()
	if err != nil {
		log.Fatalf("Error seeding user and post table %v\n", err)
	}

	postUpdate := models.Post{
		ID:       1,
		AuthorID: post.AuthorID,
		Title:    randomizer.GenerateName(),
		Content:  "Modified content",
	}

	updatedPost, err := postUpdate.UpdatePost(server.DB)
	if err != nil {
		t.Errorf("this is the error updating the post: %v\n", err)
		return
	}
	assert.Equal(t, updatedPost.ID, postUpdate.ID)
	assert.Equal(t, updatedPost.Title, postUpdate.Title)
	assert.Equal(t, updatedPost.Content, postUpdate.Content)
	assert.Equal(t, updatedPost.AuthorID, postUpdate.AuthorID)

}

func TestDeleteAPost(t *testing.T) {

	err := refreshUserAndPostTable()
	if err != nil {
		log.Fatalf("Error refreshing user and post table %v\n", err)
	}

	post, err := seedOneUserAndOnePost()
	if err != nil {
		log.Fatalf("Error seeding user and post table %v\n", err)
	}
	isDeleted, err := postInstance.DeletePost(server.DB, post.ID, post.AuthorID)
	if err != nil {
		t.Errorf("this is the error updating the user: %v\n", err)
	}
	assert.Equal(t, int(isDeleted), 1)

}
