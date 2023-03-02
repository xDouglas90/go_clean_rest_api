package service

import (
	"errors"
	"math/rand"

	"github.com/xdouglas90/gomux-rest-api/entity"
	"github.com/xdouglas90/gomux-rest-api/repository"
)

type PostService interface {
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
	Validate(post *entity.Post) error
}

type service struct{}

var postRepo repository.PostRepository

func NewPostService(repo repository.PostRepository) PostService {
	postRepo = repo
	return &service{}
}

func (*service) Create(post *entity.Post) (*entity.Post, error) {
	post.ID = rand.Int63()
	return postRepo.Save(post)
}

func (*service) FindAll() ([]entity.Post, error) {
	return postRepo.FindAll()
}

func (*service) Validate(post *entity.Post) error {
	if post == nil {
		err := errors.New("the post is empty")
		return err
	}
	if post.Title == "" {
		err := errors.New("the post title is empty")
		return err
	}
	return nil
}
