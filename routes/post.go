package routes

import (
	"template/controller"
	router "template/http"
	"template/repository"

	"template/service"
)

var (
	dbRepository   repository.PostRepository = repository.NewDatabaseRepository()
	postService    service.PostService       = service.NewPostService(dbRepository)
	postController controller.PostController = controller.NewPostController(postService)
)

type PostRoutes struct{}

func (c *PostRoutes) Routing(httpRouter router.Router) {
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)
}
