package app

import (
  "log"
  "net/http"
  "os"
  "path/filepath"

  "github.com/lightbliss/clean_architecture_blog/internal/adapters/postrepo/filesystem"
  "github.com/lightbliss/clean_architecture_blog/internal/adapters/renderer/goldmark"
  "github.com/lightbliss/clean_architecture_blog/internal/core/blog"
  "github.com/lightbliss/clean_architecture_blog/internal/ui/web"
  "github.com/lightbliss/clean_architecture_blog/pkg/env"
)

type Context struct {
	Port int
	TemplatePath string
	StaticPath string
	PostPath string
	BaseURL string
}

func NewContext() *Context {
	return &Context{
		Port: env.GetInt("PORT", 3000),
		TemplatePath: env.GetString("TEMPLATE_PATH", filepath.Join("web", "template")),
		StaticPath: env.GetString("STATIC_PATH", filepath.Join("web", "static")),
		PostPath: env.GetString("POST_PATH", filepath.Join("posts")),
		BaseURL: env.GetString("BASE_URL", "http://localho.st:3000"),
	}
}

func (c *Context) WebServer() *web.Server {
	return web.NewServer(c.Port, c.Router, c.Logger)
}

func (c *Context) Router() *web.Router {
	return web.NewRouter(c.TemplatePath, c.StaticPath, c.Usecases(), c.BaseURL)
}

func (c *Context) Usecases() *web.Usecases {
	return &web.Usecases{
		ViewPost: c.ViewPostUsecase()
	}
}

func (c *Context) ViewPostUsecase() *blog.ViewPostUsecase {
	return blog.NewViewPostUsecase(c.PostRepo(), c.Renderer())
}

func (c *Context) PostRepo() *filesystem.PostRepo {
	return filesystem.NewPostRepo(c.PostPath)
}

func (c *Context) Renderer() *goldmark.Renderer {
	return goldmark.NewRenderer()
}

func (c *Context) Logger() *log.Logger {
	return log.New(os.Stdout, "web: ", log.Ldate|log.Ltime|log.LUTC)
}
