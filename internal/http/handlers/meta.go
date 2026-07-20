package handlers

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/trigex/trigex.moe/internal/views"
)

const defaultSiteDescription = "Trigex's personal website with music releases, programming projects, blog posts, and social links."

func (h *PageHandlers) pageMeta(title, description, path string) views.PageMeta {
	return views.MetaWithDefaults(views.PageMeta{
		Title:        title,
		Description:  normalizeDescription(description),
		CanonicalURL: h.absoluteURL(path),
		OGImageURL:   h.absoluteURL("/static/img/android-chrome-512x512.png"),
	})
}

func (h *PageHandlers) articleMeta(title, description, path string, publishedAt time.Time) views.PageMeta {
	meta := h.pageMeta(title, description, path)
	meta.OGType = "article"
	meta.PublishedTime = publishedAt.UTC().Format(time.RFC3339)
	return views.MetaWithDefaults(meta)
}

func noIndexMeta(title, description string) views.PageMeta {
	return views.MetaWithDefaults(views.PageMeta{
		Title:       title,
		Description: normalizeDescription(description),
		Robots:      "noindex,nofollow",
	})
}

func (h *PageHandlers) absoluteURL(path string) string {
	base := strings.TrimRight(h.siteURL, "/")
	if path == "" || path == "/" {
		return base + "/"
	}
	return base + "/" + strings.TrimLeft(path, "/")
}

func normalizeDescription(description string) string {
	description = strings.Join(strings.Fields(strings.TrimSpace(description)), " ")
	if description == "" {
		return defaultSiteDescription
	}
	if len(description) <= 160 {
		return description
	}
	return strings.TrimSpace(description[:157]) + "..."
}

func mustJSONLD(value any) string {
	data, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	return string(data)
}
