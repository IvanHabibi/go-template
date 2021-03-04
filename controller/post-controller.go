package controller

import (
	"encoding/json"
	"net/http"

	"template/entity"
	"template/function"
	"template/service"
)

type PostController interface {
	GetPosts(resp http.ResponseWriter, req *http.Request)
	AddPost(resp http.ResponseWriter, req *http.Request)
}

type controller struct{}

var (
	postService service.PostService
)

// injection for stability
func NewPostController(service service.PostService) PostController {
	postService = service
	return &controller{}

}

func (*controller) GetPosts(resp http.ResponseWriter, req *http.Request) {
	Posts, err := postService.FindAll()
	if err != nil {
		_ = function.SendResponse(resp, http.StatusInternalServerError, "something happen", nil)
		return
	}
	_ = function.SendResponse(resp, http.StatusOK, "success", Posts)

}

func (*controller) AddPost(resp http.ResponseWriter, req *http.Request) {

	var post entity.Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		_ = function.SendResponse(resp, http.StatusInternalServerError, "something happen", nil)
		return
	}

	err = postService.Validate(&post)
	if err != nil {
		_ = function.SendResponse(resp, http.StatusInternalServerError, "something happen", nil)
		return
	}

	postResult, err := postService.Create(&post)
	if err != nil {
		_ = function.SendResponse(resp, http.StatusInternalServerError, "something happen", nil)
		return
	}

	_ = function.SendResponse(resp, http.StatusOK, "success", postResult)

}
