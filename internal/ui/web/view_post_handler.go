package web

import (
	"fmt"
  	"html/template"
  	"net/http"
  	"path"

	"github.com/lightbliss/clean_architecture_blog/internal/core/blog"
)

type ViewPostHandler struct {
	usecase ViewPostUsecase
	template *TemplateRenderer
}

func NewViewPostHandler(usecase ViewPostUsecase, template *TemplateRenderer) *ViewPostHandler {
	return &ViewPostHandler{
		usecase: usecase,
		template: template
	}
}

func(h *ViewPostHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	path := path.Base(req.URL.Path)
	renderedPost, err := h.usecase.GetPostByPath(path)
	
	switch err {
	case nil:
		res.WriteHeader(http.StatusOK)
		h.template.Render(res, "view_post.html", h.toViewModel(renderedPost))
	case blog.ErrPostNotFound:
		res.WriteHeader(http.StatusNotFound)
		h.template.Render(res, "404.html", nil)
	default:
		res.WriteHeader(http.StatusInternalServerError)
		h.template.Render(res, "500.html", nil)
	}
}

func (h *ViewPostHandler) toViewModel(p blog.RenderedPost) postViewModel {
	return postViewModel{
		Title: p.Post.Title,
		Author: p.Post.Author,
		Date: p.Post.Time.Format(time.DateOnly),
		Description: p.Post.Description,
		ImagePath: p.Post.ImagePath,
		Path: fmt.Sprintf("post/%s", p.Post.Path),
		Content: template.HTML(p.HTML),
	}
}

type postViewModel struct {
	Title string
	Author string
	Date string
	Description string
	ImagePath string
	Path string
	Content template.HTML
}
