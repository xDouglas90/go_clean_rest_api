package main

import (
	"fmt"
	"net/http"

	controller "github.com/xdouglas90/gomux-rest-api/controllers"
	router "github.com/xdouglas90/gomux-rest-api/http/router"
	"github.com/xdouglas90/gomux-rest-api/repository"
	"github.com/xdouglas90/gomux-rest-api/service"
)

var (
	postRepository repository.PostRepository = repository.NewFirestoreRepository()
	postService    service.PostService       = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService)
	httpRouter     router.Router             = router.NewChiRouter()
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
