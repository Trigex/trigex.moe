package views

import "strings"

type PageMeta struct {
	Title         string
	Description   string
	CanonicalURL  string
	Robots        string
	OGType        string
	OGImageURL    string
	TwitterCard   string
	PublishedTime string
	JSONLD        []string
}

func MetaWithDefaults(meta PageMeta) PageMeta {
	if strings.TrimSpace(meta.Robots) == "" {
		meta.Robots = "index,follow"
	}
	if strings.TrimSpace(meta.OGType) == "" {
		meta.OGType = "website"
	}
	if strings.TrimSpace(meta.TwitterCard) == "" {
		if strings.TrimSpace(meta.OGImageURL) != "" {
			meta.TwitterCard = "summary_large_image"
		} else {
			meta.TwitterCard = "summary"
		}
	}
	return meta
}
