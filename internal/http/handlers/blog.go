package handlers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/labstack/echo/v5"

	"github.com/trigex/trigex.moe/internal/http/render"
	"github.com/trigex/trigex.moe/internal/views"
)

func (h *PageHandlers) ServeBlogIndexPage(c *echo.Context) error {
	rows, err := h.store.ListBlogPosts(c.Request().Context())
	if err != nil {
		return err
	}

	posts := make([]views.BlogPost, 0, len(rows))
	for _, row := range rows {
		posts = append(posts, views.BlogPost{
			Slug:        row.Slug,
			Title:       row.Title,
			Excerpt:     row.Excerpt,
			Body:        row.Body,
			PublishedAt: row.PublishedAt,
		})
	}

	return render.Templ(c, http.StatusOK, views.Layout("trigex.moe | Blog", views.BlogPage(posts)))
}

func (h *PageHandlers) ServeBlogPostPage(c *echo.Context) error {
	post, err := h.store.GetBlogPostBySlug(c.Request().Context(), c.Param("slug"))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return echo.NewHTTPError(http.StatusNotFound, "blog post not found")
		}
		return err
	}

	bodyHTML, err := render.MarkdownHTML(post.Body)
	if err != nil {
		return err
	}

	return render.Templ(c, http.StatusOK, views.Layout("trigex.moe | "+post.Title, views.BlogPostPage(views.BlogPost{
		Slug:        post.Slug,
		Title:       post.Title,
		Excerpt:     post.Excerpt,
		Body:        post.Body,
		BodyHTML:    bodyHTML,
		PublishedAt: post.PublishedAt,
	})))
}
