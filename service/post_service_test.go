package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/xdouglas90/gomux-rest-api/entity"
)

type MockPostRepository struct {
	mock.Mock
}

func (mock *MockPostRepository) Save(post *entity.Post) (*entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)

	return result.(*entity.Post), args.Error(1)
}

func (mock *MockPostRepository) FindAll() ([]entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)

	return result.([]entity.Post), args.Error(1)
}

func TestFindAll(t *testing.T) {
	mockRepo := new(MockPostRepository)
	var identifier int64 = 1
	post := entity.Post{ID: identifier, Title: "Test", Content: "Hello World"}

	// Setup expectations
	mockRepo.On("FindAll").Return([]entity.Post{post}, nil)

	testService := NewPostService(mockRepo)

	result, _ := testService.FindAll()

	// Mock Assertion: Behavioral
	mockRepo.AssertExpectations(t)

	// Data Assertion
	assert.Equal(t, identifier, result[0].ID)
	assert.Equal(t, "Test", result[0].Title)
	assert.Equal(t, "Hello World", result[0].Content)
}

func TestCreate(t *testing.T) {
	mockRepo := new(MockPostRepository)
	post := entity.Post{Title: "Test", Content: "Hello World"}

	// Setup expectations
	mockRepo.On("Save").Return(&post, nil)

	testService := NewPostService(mockRepo)

	result, err := testService.Create(&post)

	// Mock Assertion: Behavioral
	mockRepo.AssertExpectations(t)

	// Data Assertion
	assert.NotNil(t, result.ID)
	assert.Equal(t, "Test", result.Title)
	assert.Equal(t, "Hello World", result.Content)
	assert.Nil(t, err)
}

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
