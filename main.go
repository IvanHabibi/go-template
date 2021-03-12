package main

import (
	"fmt"
	"net/http"
	"template/auth"
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

	migrateRoutes := routes.MigrateRoutes{}
	migrateRoutes.Routing(httpRouter)

	postRoutes := routes.PostRoutes{}
	postRoutes.Routing(httpRouter)

	// authRoutes := routes.AuthRoutes{}
	// authRoutes.Routing(httpRouter)

	auth.Authorization() // yang coba2 login keycloak dan bedah tokenya di golang (g bakal kepake)
	// auth.InstrospectToken()

	// httpRouter.GET("/migrate", func(w http.ResponseWriter, r *http.Request) {
	// 	migration.Migrate(w, r)
	// })

	httpRouter.SERVE()

}
