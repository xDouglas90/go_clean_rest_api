package repository

import (
	"context"
	"log"
	"os"

	"github.com/xdouglas90/gomux-rest-api/entity"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
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
	it := client.Collection(collecionName).Documents(ctx)
	for {
		doc, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("failed to iterate the list of posts: %v", err)
			return nil, err
		}

		post := entity.Post{
			ID:      doc.Data()["ID"].(int64),
			Title:   doc.Data()["Title"].(string),
			Content: doc.Data()["Text"].(string),
		}
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

// Delete: TODO
func (r *repo) Delete(post *entity.Post) error {
	return nil
}
