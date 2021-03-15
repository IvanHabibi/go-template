package main

import (
	"fmt"
	"net/http"
	router "template/http"
	"template/routes"
)

var (
	// dbRepository   repository.PostRepository = repository.NewDatabaseRepository()
	// postService    service.PostService       = service.NewPostService(dbRepository)
	// postController controller.PostController = controller.NewPostController(postService)
	httpRouter router.Router = router.NewMuxRouter()
)

func main() {

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Healthy")
	})

	// for migration purpose
	migrateRoutes := routes.MigrateRoutes{}
	migrateRoutes.Routing(httpRouter)

	// for application root
	postRoutes := routes.PostRoutes{}
	postRoutes.Routing(httpRouter)


	httpRouter.SERVE()

}
