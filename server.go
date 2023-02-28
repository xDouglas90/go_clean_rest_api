package main

import (
	"fmt"
	"net/http"

	controller "github.com/xdouglas90/gomux-rest-api/controllers"
	router "github.com/xdouglas90/gomux-rest-api/http"
)

var (
	postController controller.PostController = controller.NewPostController()
	httpRouter     router.Router             = router.NewMuxRouter()
)

func main() {
	const port string = ":8080"
	// s := router.PathPrefix("/api/v1").Subrouter()
	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Up and running...")
	})
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)

	httpRouter.SERVE(port)
}
