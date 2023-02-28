package repository

import (
	"context"
	"log"
	"os"

	"github.com/xdouglas90/gomux-rest-api/entity"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

type repo struct{}

func NewFirestoreRepository() PostRepository {
	return &repo{}
}

const (
	projectId     string = "gomux-rest-api"
	collecionName string = "posts"
)

func (*repo) Save(post *entity.Post) (*entity.Post, error) {
	ctx := context.Background()
	client, err := getClient(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v", err)
		return nil, err
	}

	defer client.Close()
	client.Collection(collecionName).Add(ctx, map[string]interface{}{
		"ID":      post.ID,
		"Title":   post.Title,
		"Content": post.Content,
	})
	if err != nil {
		log.Fatalf("failed adding a new post: %v", err)
		return nil, err
	}

	return post, nil
}

func (*repo) FindAll() ([]entity.Post, error) {
	ctx := context.Background()
	client, err := getClient(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v", err)
		return nil, err
	}

	defer client.Close()

	var posts []entity.Post
	iterator := client.Collection(collecionName).Documents(ctx)
	for {
		doc, err := iterator.Next()
		if err != nil {
			log.Fatalf("failed to iterate the list of posts: %v", err)
			return nil, err
		}
		var post entity.Post
		doc.DataTo(&post)
		posts = append(posts, post)
	}

	return posts, nil
}

func getClient(ctx context.Context) (*firestore.Client, error) {
	opt := option.WithCredentialsFile(os.Getenv("GOOGLE_APP_CREDENTIALS"))
	client, err := firestore.NewClient(ctx, projectId, opt)
	if err != nil {
		log.Fatalf("error creating firestore client: %v", err)
		return nil, err
	}

	return client, nil
}
