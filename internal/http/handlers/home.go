package handlers

import (
	"net/http"
	"strings"

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

	meta := h.pageMeta("trigex.moe | Home", data.Bio, "/")
	meta.JSONLD = []string{
		mustJSONLD(map[string]any{
			"@context":    "https://schema.org",
			"@type":       "WebSite",
			"name":        h.siteName,
			"url":         h.absoluteURL("/"),
			"description": meta.Description,
		}),
		mustJSONLD(map[string]any{
			"@context": "https://schema.org",
			"@type":    "Person",
			"name":     "Trigex",
			"url":      h.absoluteURL("/"),
			"sameAs":   socialLinkURLs(links),
			"jobTitle": "Software Developer, Music Producer, DJ, and Sysadmin",
		}),
	}

	return render.Templ(c, http.StatusOK, views.Layout(meta, views.HomePage(data)))
}

func socialLinkURLs(links []views.Link) []string {
	urls := make([]string, 0, len(links))
	for _, link := range links {
		if strings.HasPrefix(link.URL, "http://") || strings.HasPrefix(link.URL, "https://") {
			urls = append(urls, link.URL)
		}
	}
	return urls
}
