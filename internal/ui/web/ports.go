package web

import "github.com/lightbliss/blog/internal/core/blog"

type Usecases interface {
	ViewPost ViewPostUsecase
}

type ViewPostUsecase interface {
	Run(path string) (blog.RenderedPost, error)
}