package routes

import (
	"fmt"
	"net/http"
	router "template/http"
)

var (
	httpRouterAuth router.Router = router.NewMuxRouter()
)

type AuthhRoutes struct{}

func (c *AuthhRoutes) Routing() {
	httpRouterAuth.GET("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Up n run")
	})
}
