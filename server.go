package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	controller "github.com/xdouglas90/gomux-rest-api/controller"
	router "github.com/xdouglas90/gomux-rest-api/http/router"
	"github.com/xdouglas90/gomux-rest-api/repository"
	"github.com/xdouglas90/gomux-rest-api/service"
)

var (
	postRepository repository.PostRepository = repository.NewSQLiteRepository()
	postService    service.PostService       = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService)
	httpRouter     router.Router             = router.NewMuxRouter()
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)
	httpRouter.SERVE(os.Getenv("PORT"))
}
