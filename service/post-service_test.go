package service

import (
	"template/entity"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) Save(post *entity.Post) (*entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Post), args.Error(1)
}

func (mock *MockRepository) FindAll() ([]entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Post), args.Error(1)
}

func TestFindAll(t *testing.T) {
	mockRepo := new(MockRepository)

	var identifier string = "C" // check the data type

	post := entity.Post{Title: "A", Text: identifier}

	mockRepo.On("FindAll").Return([]entity.Post{post}, nil)

	testService := NewPostService(mockRepo)

	result, _ := testService.FindAll()

	mockRepo.AssertExpectations(t)

	assert.Equal(t, "A", result[0].Title)

	assert.Equal(t, identifier, result[0].Text) // check the response of find all on

	assert.Equal(t, "C", result[0].Text) //will be error because is have to be C

}

func TestCreate(t *testing.T) {
	mockRepo := new(MockRepository)

	post := entity.Post{Title: "A", Text: "B"}

	mockRepo.On("Save").Return(&post, nil)

	testService := NewPostService(mockRepo)

	result, err := testService.Create(&post)

	mockRepo.AssertExpectations(t)

	assert.Equal(t, "A", result.Title)

	assert.Equal(t, "B", result.Text) //will be error because is have to be C

	assert.Nil(t, err)
}

func TestValidateEmptyPost(t *testing.T) {
	testService := NewPostService(nil)

	err := testService.Validate(nil)

	assert.NotNil(t, err) //check the err not nil
	//          expected value   |   actual value
	assert.Equal(t, "Post are empty", err.Error()) //check the error have the message same as intended

}

func TestValidateEmptyPostTitle(t *testing.T) {
	post := entity.Post{Title: "", Text: "yas"}

	testService := NewPostService(nil)

	err := testService.Validate(&post)

	assert.NotNil(t, err) //check the err not nil
	//          expected value   |   actual value
	assert.Equal(t, "Title are empty", err.Error()) //check the error have the message same as intended

}
