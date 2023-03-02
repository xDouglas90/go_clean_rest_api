package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xdouglas90/gomux-rest-api/entity"
)

func TestValidateEmptyPost(t *testing.T) {
	testService := NewPostService(nil)

	err := testService.Validate(nil)

	assert.NotNil(t, err)
	assert.Equal(t, "the post is empty", err.Error())
}

func TestValidateEmptyPostTitle(t *testing.T) {
	testService := NewPostService(nil)

	post := entity.Post{ID: 1, Title: "", Content: "Hello World"}

	err := testService.Validate(&post)

	assert.NotNil(t, err)
	assert.Equal(t, "the post title is empty", err.Error())
}
