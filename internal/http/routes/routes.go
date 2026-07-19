package routes

import (
	"github.com/labstack/echo/v5"

	"github.com/trigex/trigex.moe/internal/http/handlers"
)

func Register(e *echo.Echo, pageHandlers *handlers.PageHandlers, adminAuth echo.MiddlewareFunc) {
	e.GET("/", pageHandlers.ServeHomePage)
	e.GET("/music", pageHandlers.ServeMusicPage)
	e.GET("/projects", pageHandlers.ServeProjectsPage)
	e.GET("/blog", pageHandlers.ServeBlogIndexPage)
	e.GET("/blog/rss.xml", pageHandlers.ServeBlogRSSFeed)
	e.GET("/blog/:slug", pageHandlers.ServeBlogPostPage)

	admin := e.Group("/admin", adminAuth)
	admin.GET("/", pageHandlers.ServeAdminPage)
	admin.GET("/blog/:slug/edit", pageHandlers.ServeEditBlogPage)
	admin.POST("/blog/preview", pageHandlers.ServeBlogPreview)
	admin.POST("/blog/upload-image", pageHandlers.UploadBlogImage)
	admin.POST("/blog", pageHandlers.CreateBlogPost)
	admin.POST("/blog/:slug", pageHandlers.UpdateBlogPost)
	admin.POST("/blog/:slug/delete", pageHandlers.DeleteBlogPost)
	admin.POST("/links", pageHandlers.CreateSocialLink)
	admin.POST("/links/:id", pageHandlers.UpdateSocialLink)
	admin.POST("/links/:id/delete", pageHandlers.DeleteSocialLink)
	admin.POST("/music", pageHandlers.CreateMusicTrack)
	admin.POST("/music/:id/delete", pageHandlers.DeleteMusicTrack)
	admin.POST("/projects", pageHandlers.CreateProject)
	admin.POST("/projects/:id/delete", pageHandlers.DeleteProject)
}
