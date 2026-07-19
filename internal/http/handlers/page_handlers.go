package handlers

import "github.com/trigex/trigex.moe/internal/content"

type PageHandlers struct {
	siteName string
	siteURL  string
	store    *content.Store
}

func NewPageHandlers(siteName, siteURL string, store *content.Store) *PageHandlers {
	return &PageHandlers{
		siteName: siteName,
		siteURL:  siteURL,
		store:    store,
	}
}
