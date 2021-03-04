package service

import (
	"errors"
	"template/entity"
	"template/repository"
)

type PostService interface {
	Validate(post *entity.Post) error
	Create(post *entity.Post) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}

type service struct{}

var (
	repo repository.PostRepository
)

func NewPostService(repository repository.PostRepository) PostService {
	repo = repository
	return &service{}
}

func (*service) Validate(post *entity.Post) error {
	if post == nil {
		err := errors.New("Post are empty")
		return err
	}

	if post.Title == "" {
		err := errors.New("Title are empty")
		return err
	}

	if post.Text == "" {
		err := errors.New("Text are empty")
		return err
	}

	return nil

}
func (*service) Create(post *entity.Post) (*entity.Post, error) {

	postResult, err := repo.Save(post)
	if err != nil {
		return nil, err
	}

	return postResult, nil
}

func (*service) FindAll() ([]entity.Post, error) {
	posts, err := repo.FindAll()
	return posts, err
}
