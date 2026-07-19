package app

import (
	"context"
	"crypto/subtle"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"

	"github.com/trigex/trigex.moe/assets"
	"github.com/trigex/trigex.moe/internal/content"
	"github.com/trigex/trigex.moe/internal/http/handlers"
	"github.com/trigex/trigex.moe/internal/http/routes"
)

const SiteName = "trigex.moe"
const SiteURL = "https://trigex.moe"

func New(adminUser, adminPass, dbPath, dataDir string) (*echo.Echo, error) {
	store, err := content.Open(context.Background(), dbPath)
	if err != nil {
		return nil, err
	}

	uploadsDir := filepath.Join(dataDir, "uploads")
	if err := os.MkdirAll(uploadsDir, 0o755); err != nil {
		return nil, err
	}

	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	e.StaticFS("/static", assets.FS)
	e.Static("/uploads", uploadsDir)

	adminAuth, err := newAdminAuth(adminUser, adminPass)
	if err != nil {
		return nil, err
	}

	pageHandlers := handlers.NewPageHandlers(SiteName, SiteURL, store)
	e.HTTPErrorHandler = handlers.NewErrorHandler(SiteName)
	routes.Register(e, pageHandlers, adminAuth)

	return e, nil
}

func newAdminAuth(user, pass string) (echo.MiddlewareFunc, error) {
	return middleware.BasicAuth(func(c *echo.Context, username, password string) (bool, error) {
		userMatch := subtle.ConstantTimeCompare([]byte(username), []byte(user)) == 1
		passMatch := subtle.ConstantTimeCompare([]byte(password), []byte(pass)) == 1
		return userMatch && passMatch, nil
	}), nil
}
