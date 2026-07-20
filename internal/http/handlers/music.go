package handlers

import (
	"net/http"

	"github.com/labstack/echo/v5"

	"github.com/trigex/trigex.moe/internal/http/render"
	"github.com/trigex/trigex.moe/internal/views"
)

func (h *PageHandlers) ServeMusicPage(c *echo.Context) error {
	rows, err := h.store.ListMusicTracks(c.Request().Context())
	if err != nil {
		return err
	}

	data := make([]views.Track, 0, len(rows))
	for _, row := range rows {
		data = append(data, views.Track{
			Title:         row.Title,
			Genre:         row.Genre,
			FlacURL:       row.FlacUrl,
			Mp3URL:        row.Mp3Url,
			SoundcloudURL: row.SoundcloudUrl,
			YoutubeURL:    row.YoutubeUrl,
			CoverImage:    row.CoverImage,
			ReleaseDate:   row.ReleaseDate,
		})
	}

	return render.Templ(c, http.StatusOK, views.Layout(
		h.pageMeta(
			"trigex.moe | Music",
			"Music releases by Trigex with streaming and download links across hardcore, makina, uptempo, drum and bass, psytrance, and more.",
			"/music",
		),
		views.MusicPage(data),
	))
}
