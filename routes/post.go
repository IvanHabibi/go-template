package routes

import (
	"template/controller"
	router "template/http"
	"template/repository"

	"template/service"
)

var (
	dbRepository   repository.PostRepository = repository.NewDatabaseRepository() //query ke DB
	postService    service.PostService       = service.NewPostService(dbRepository) //apa yang diolah (business logic)
	postController controller.PostController = controller.NewPostController(postService) //apa yang akan dikirimkan ke user (scope)
)

type PostRoutes struct{}

func (c *PostRoutes) Routing(httpRouter router.Router) {
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/post", postController.AddPost)
}
