package repository

import (
	"template/entity"
)

type PostRepository interface {
	FindAll() ([]entity.Post, error)
	Save(post *entity.Post) (*entity.Post, error)
	// Update(post *entity.Post, id *int) (*entity.Post, error)
	// Delete(id *int) (*entity.Post, error)
}
