package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xdouglas90/gomux-rest-api/entity"
	"github.com/xdouglas90/gomux-rest-api/repository"
	"github.com/xdouglas90/gomux-rest-api/service"
)

const (
	TITLE   string = "Title 1"
	CONTENT string = "Content 1"
)

var (
	postRepo       repository.PostRepository = repository.NewSQLiteRepository()
	postSrv        service.PostService       = service.NewPostService(postRepo)
	postController PostController            = NewPostController(postSrv)
)

func TestAddPost(t *testing.T) {
	// Create a new HTTP POST request
	jsonPost := []byte(`{"title": "` + TITLE + `", "content": "` + CONTENT + `"}`)
	req, _ := http.NewRequest("POST", "/posts", bytes.NewBuffer(jsonPost))

	// Assign HTTP Handler function (controller AddPost function)
	handler := http.HandlerFunc(postController.AddPost)

	// Record HTTP Response (httptest)
	res := httptest.NewRecorder()

	// Dispatch HTTP request
	handler.ServeHTTP(res, req)

	// Add Assertions on the HTTP Status code and the response
	status := res.Code
	if status != http.StatusOK {
		t.Errorf("handler returned a wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Decode the HTTP response
	var post entity.Post
	json.NewDecoder(io.Reader(res.Body)).Decode(&post)

	// Assert HTTP response
	assert.NotNil(t, post.ID)
	assert.NotNil(t, TITLE, post.Title)
	assert.NotNil(t, CONTENT, post.Content)

	// Clean up database
	tearDown(post.ID)
}

func setup() {
	postRepo.Save(&entity.Post{Title: TITLE, Content: CONTENT})
}

func tearDown(postID int64) {
	var post entity.Post = entity.Post{
		ID: postID,
	}
	postRepo.Delete(&post)
}

func TestGetPosts(t *testing.T) {
	// Insert new post into database
	setup()

	// Create a new HTTP GET request
	req, _ := http.NewRequest("GET", "/posts", nil)

	// Assign HTTP Handler function (controller GetPosts function)
	handler := http.HandlerFunc(postController.GetPosts)

	// Record HTTP Response (httptest)
	res := httptest.NewRecorder()

	// Dispatch HTTP request
	handler.ServeHTTP(res, req)

	// Add Assertions on the HTTP Status code and the response
	status := res.Code
	if status != http.StatusOK {
		t.Errorf("handler returned a wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Decode the HTTP response
	var posts []entity.Post
	json.NewDecoder(io.Reader(res.Body)).Decode(&posts)

	// Assert HTTP response
	assert.NotNil(t, posts[0].ID)
	assert.NotNil(t, posts[0].Title)
	assert.NotNil(t, posts[0].Content)
	assert.Equal(t, 1, len(posts))

	// Clean up database
	tearDown(posts[0].ID)
}
