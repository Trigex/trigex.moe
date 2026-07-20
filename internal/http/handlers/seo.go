package handlers

import (
	"encoding/xml"
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
)

type urlSet struct {
	XMLName xml.Name     `xml:"urlset"`
	Xmlns   string       `xml:"xmlns,attr"`
	URLs    []sitemapURL `xml:"url"`
}

type sitemapURL struct {
	Loc     string `xml:"loc"`
	LastMod string `xml:"lastmod,omitempty"`
}

func (h *PageHandlers) ServeRobotsTXT(c *echo.Context) error {
	body := "User-agent: *\n" +
		"Allow: /\n" +
		"Disallow: /admin\n" +
		"Sitemap: " + h.absoluteURL("/sitemap.xml") + "\n"

	c.Response().Header().Set(echo.HeaderContentType, "text/plain; charset=utf-8")
	c.Response().WriteHeader(http.StatusOK)
	_, err := c.Response().Write([]byte(body))
	return err
}

func (h *PageHandlers) ServeSitemapXML(c *echo.Context) error {
	posts, err := h.store.ListBlogPosts(c.Request().Context())
	if err != nil {
		return err
	}

	urls := []sitemapURL{
		{Loc: h.absoluteURL("/")},
		{Loc: h.absoluteURL("/music")},
		{Loc: h.absoluteURL("/projects")},
		{Loc: h.absoluteURL("/blog")},
	}
	for _, post := range posts {
		urls = append(urls, sitemapURL{
			Loc:     h.absoluteURL("/blog/" + post.Slug),
			LastMod: post.PublishedAt.UTC().Format(time.DateOnly),
		})
	}

	doc := urlSet{
		Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9",
		URLs:  urls,
	}

	c.Response().Header().Set(echo.HeaderContentType, "application/xml; charset=utf-8")
	c.Response().WriteHeader(http.StatusOK)

	enc := xml.NewEncoder(c.Response())
	enc.Indent("", "  ")
	if _, err := c.Response().Write([]byte(xml.Header)); err != nil {
		return err
	}
	if err := enc.Encode(doc); err != nil {
		return err
	}
	return enc.Flush()
}
