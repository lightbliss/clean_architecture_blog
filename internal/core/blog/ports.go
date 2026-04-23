package blog

import "errors"

const ErrPostNotFound = errors.New("post not found")

type postRepo interface {
	GetPostByPath(path string) (Post, error)
	GetAllPosts() ([]Post, error)
}

type renderer interface {
	Render(content string) (string, error)
}