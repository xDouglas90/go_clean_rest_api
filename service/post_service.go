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

var repo repository.PostRepository = repository.NewFirestoreRepository()

func NewPostService() PostService {
	return &service{}
}

func (*service) Create(post *entity.Post) (*entity.Post, error) {
	post.ID = rand.Int63()
	return repo.Save(post)
}

func (*service) FindAll() ([]entity.Post, error) {
	return repo.FindAll()
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
