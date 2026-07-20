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

	return render.Templ(c, http.StatusOK, views.Layout(
		h.pageMeta(
			"trigex.moe | Projects",
			"Programming projects by Trigex, including Go, C++, Blazor, and .NET work with links to source repositories.",
			"/projects",
		),
		views.ProjectsPage(data),
	))
}
