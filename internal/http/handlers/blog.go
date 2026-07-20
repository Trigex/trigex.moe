package handlers

import (
	"database/sql"
	"errors"
	"net/http"
	"time"

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

	return render.Templ(c, http.StatusOK, views.Layout(
		h.pageMeta(
			"trigex.moe | Blog",
			"Blog posts, site updates, and longer thoughts from Trigex.",
			"/blog",
		),
		views.BlogPage(posts),
	))
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

	viewPost := views.BlogPost{
		Slug:        post.Slug,
		Title:       post.Title,
		Excerpt:     post.Excerpt,
		Body:        post.Body,
		BodyHTML:    bodyHTML,
		PublishedAt: post.PublishedAt,
	}

	meta := h.articleMeta(
		"trigex.moe | "+post.Title,
		post.Excerpt,
		"/blog/"+post.Slug,
		post.PublishedAt,
	)
	meta.JSONLD = []string{
		mustJSONLD(map[string]any{
			"@context":         "https://schema.org",
			"@type":            "BlogPosting",
			"headline":         post.Title,
			"description":      meta.Description,
			"url":              h.absoluteURL("/blog/" + post.Slug),
			"mainEntityOfPage": h.absoluteURL("/blog/" + post.Slug),
			"datePublished":    post.PublishedAt.UTC().Format(time.RFC3339),
			"dateModified":     post.PublishedAt.UTC().Format(time.RFC3339),
			"author": map[string]any{
				"@type": "Person",
				"name":  "Trigex",
			},
			"publisher": map[string]any{
				"@type": "Person",
				"name":  "Trigex",
			},
			"image": meta.OGImageURL,
		}),
	}

	return render.Templ(c, http.StatusOK, views.Layout(
		meta,
		views.BlogPostPage(viewPost),
	))
}
