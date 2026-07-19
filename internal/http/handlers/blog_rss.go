package handlers

import (
	"encoding/xml"
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
)

type rssDate time.Time

func (d rssDate) MarshalText() ([]byte, error) {
	return []byte(time.Time(d).UTC().Format(time.RFC1123Z)), nil
}

type rssFeed struct {
	XMLName xml.Name   `xml:"rss"`
	Version string     `xml:"version,attr"`
	Channel rssChannel `xml:"channel"`
}

type rssChannel struct {
	Title         string    `xml:"title"`
	Link          string    `xml:"link"`
	Description   string    `xml:"description"`
	Language      string    `xml:"language"`
	LastBuildDate rssDate   `xml:"lastBuildDate"`
	Items         []rssItem `xml:"item"`
}

type rssItem struct {
	Title       string  `xml:"title"`
	Link        string  `xml:"link"`
	GUID        string  `xml:"guid"`
	PubDate     rssDate `xml:"pubDate"`
	Description string  `xml:"description"`
}

func (h *PageHandlers) ServeBlogRSSFeed(c *echo.Context) error {
	posts, err := h.store.ListBlogPosts(c.Request().Context())
	if err != nil {
		return err
	}

	items := make([]rssItem, 0, len(posts))
	for _, post := range posts {
		link := h.siteURL + "/blog/" + post.Slug
		items = append(items, rssItem{
			Title:       post.Title,
			Link:        link,
			GUID:        link,
			PubDate:     rssDate(post.PublishedAt),
			Description: post.Excerpt,
		})
	}

	feed := rssFeed{
		Version: "2.0",
		Channel: rssChannel{
			Title:         h.siteName + " Blog",
			Link:          h.siteURL + "/blog",
			Description:   "Posts from " + h.siteName,
			Language:      "en-us",
			LastBuildDate: rssDate(time.Now()),
			Items:         items,
		},
	}

	c.Response().Header().Set(echo.HeaderContentType, "application/rss+xml; charset=utf-8")
	c.Response().WriteHeader(http.StatusOK)

	enc := xml.NewEncoder(c.Response())
	enc.Indent("", "  ")
	if _, err := c.Response().Write([]byte(xml.Header)); err != nil {
		return err
	}
	if err := enc.Encode(feed); err != nil {
		return err
	}
	return enc.Flush()
}
