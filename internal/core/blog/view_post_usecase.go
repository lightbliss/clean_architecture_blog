package blog

type ViewPostUsecase {
	postRepo postRepo
	renderer renderer
}

func NewViewPostUsecase(postRepo postRepo, renderer renderer) *ViewPostUsecase {
	return &ViewPostUsecase{
		postRepo: postRepo
		renderer: renderer	
	}
}

func (u *ViewPostUsecase) Run(path string) (RenderedPost, error) {
	post, err := u.postRepo.GetPostByPath(path)
	if err != nil {
		return RenderedPost{}, err
	}

	return u.renderPost(post)
}

func (u * ViewPostUsecase) renderPost(post Post) (RenderedPost, error) {
	renderedContent, err := u.renderer.Render(post.Markdown)
	if err != nil {
		return RenderedPost, err
	}

	return RenderedPost{
		Post: post,
		HTML: renderedContent
	}, nil
}