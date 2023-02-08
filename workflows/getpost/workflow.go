package getpost

import (
	"context"
	"fmt"
	"time"

	gohttpclient "github.com/bozd4g/go-http-client"
	"go.temporal.io/sdk/workflow"
)

type Post struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int    `json:"userId"`
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

func GetUser(ctx context.Context, id int) (User, error) {
	client := gohttpclient.New("https://jsonplaceholder.typicode.com")

	response, err := client.Get(ctx, fmt.Sprintf("/users/%d", id))
	if err != nil {
		return User{}, fmt.Errorf("unable to fetch user: %v", err)
	}

	var user User
	if err := response.Unmarshal(&user); err != nil {
		return User{}, fmt.Errorf("failed to deserialize user: %v", err)
	}

	return user, nil
}

func GetPost(ctx context.Context, id int) (Post, error) {
	client := gohttpclient.New("https://jsonplaceholder.typicode.com")

	response, err := client.Get(ctx, fmt.Sprintf("/posts/%d", id))
	if err != nil {
		return Post{}, fmt.Errorf("unable to fetch post: %v", err)
	}

	var post Post
	if err := response.Unmarshal(&post); err != nil {
		return Post{}, fmt.Errorf("failed to deserialize post: %v", err)
	}

	return post, nil
}

type PostWithUserOutput struct {
	Author User
	Post   Post
	Hello  int
}

func GetPostWithUser(ctx workflow.Context, postId int) (*PostWithUserOutput, error) {
	ctx = workflow.WithActivityOptions(ctx, workflow.ActivityOptions{
		StartToCloseTimeout: time.Second * 5,
	})

	var post Post
	err := workflow.ExecuteActivity(ctx, GetPost, postId).Get(ctx, &post)
	if err != nil {
		return nil, err
	}

	workflow.Sleep(ctx, 10*time.Second)

	var user User
	err = workflow.ExecuteActivity(ctx, GetUser, post.UserID).Get(ctx, &user)
	if err != nil {
		return nil, err
	}

	result := PostWithUserOutput{
		Post:   post,
		Author: user,
		Hello:  123,
	}

	return &result, nil

}
