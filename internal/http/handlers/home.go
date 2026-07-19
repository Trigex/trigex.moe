package handlers

import (
	"net/http"

	"github.com/labstack/echo/v5"

	"github.com/trigex/trigex.moe/internal/http/render"
	"github.com/trigex/trigex.moe/internal/views"
)

func (h *PageHandlers) ServeHomePage(c *echo.Context) error {
	rows, err := h.store.ListSocialLinks(c.Request().Context())
	if err != nil {
		return err
	}

	links := make([]views.Link, 0, len(rows))
	for _, row := range rows {
		links = append(links, views.Link{
			ID:   row.ID,
			Name: row.Name,
			URL:  row.Url,
		})
	}

	data := views.PageData{
		Title: h.siteName,
		Name:  h.siteName,
		Bio:   "Hi I'm Trigex! Welcome to my corner of the internet! I'm a software developer, music producer, DJ, and sysadmin. (BSDs preferred ;)). I unfortunately reside in California, but I hope that'll change one day. Always looking for work!",
		Links: links,
	}

	return render.Templ(c, http.StatusOK, views.Layout("trigex.moe | Home", views.HomePage(data)))
}
