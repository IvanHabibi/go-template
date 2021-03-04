package repository

import (
	"template/datasource"
	"template/entity"
)

type repo struct{}

func NewDatabaseRepository() PostRepository {
	return &repo{}
}

func (*repo) FindAll() ([]entity.Post, error) {
	db := datasource.OpenDB()

	var posts []entity.Post

	db.Find(&posts)

	defer datasource.CloseDB()

	return posts, nil

}

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	db := datasource.OpenDB()

	db.Create(&post)

	defer datasource.CloseDB()
	return post, nil
}

// func (*repo) Update(int ID) (*entity.Post, error) {
// 	db := datasource.OpenDB()
// 	defer datasource.CloseDB()

// 	db.Create(&post)

// 	// if err != nil {
// 	// 	log.Fatalf("Failed to create a Firestore Client: %v", err)
// 	// 	return nil, err
// 	// }

// 	return post, nil
// }
