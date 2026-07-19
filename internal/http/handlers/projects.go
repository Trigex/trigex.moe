package handlers

import (
	"net/http"

	"github.com/labstack/echo/v5"

	"github.com/trigex/trigex.moe/internal/http/render"
	"github.com/trigex/trigex.moe/internal/views"
)

func (h *PageHandlers) ServeProjectsPage(c *echo.Context) error {
	rows, err := h.store.ListProjects(c.Request().Context())
	if err != nil {
		return err
	}

	data := make([]views.Project, 0, len(rows))
	for _, row := range rows {
		data = append(data, views.Project{
			Name:        row.Name,
			Description: row.Description,
			RepoURL:     row.RepoUrl,
			TechStack:   row.TechStack,
		})
	}

	return render.Templ(c, http.StatusOK, views.Layout("trigex.moe | Projects", views.ProjectsPage(data)))
}
