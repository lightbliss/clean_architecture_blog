package filesystem

import (
	"io/ioutil"
	"path/filepath"

	"github.com/lightbliss/clean_architecture_blog/internal/core/blog"
)

type PostRepo struct {
	BasePath string
}

func NewPostRepo(basePath string) *PostRepo {
	return &PostRepo{
		BasePath: basePath
	}
}

func (r PostRepo) GetPostByPath(path string) (blog.Post, error) {
	content, err := ioutil.ReadFile(filepath.Join(r.BasePath, path+".md"))
	if err != nil {
		return blog.Post{}, blog.ErrPostNotFound
	}

	post, err := ParseFileContent(string(content))
	if err != nil {
		return blog.Post{}, err
	}

	post.Path = path

	return post, nil
}


