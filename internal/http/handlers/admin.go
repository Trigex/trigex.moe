package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"html"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/labstack/echo/v5"

	"github.com/trigex/trigex.moe/internal/content"
	hrender "github.com/trigex/trigex.moe/internal/http/render"
	"github.com/trigex/trigex.moe/internal/views"
)

func (h *PageHandlers) ServeAdminPage(c *echo.Context) error {
	postsRows, err := h.store.ListBlogPosts(c.Request().Context())
	if err != nil {
		return err
	}

	posts := make([]views.BlogPost, 0, len(postsRows))
	for _, row := range postsRows {
		posts = append(posts, views.BlogPost{
			Slug:        row.Slug,
			Title:       row.Title,
			Excerpt:     row.Excerpt,
			Body:        row.Body,
			PublishedAt: row.PublishedAt,
		})
	}

	tracksRows, err := h.store.ListMusicTracks(c.Request().Context())
	if err != nil {
		return err
	}
	tracks := make([]views.AdminTrack, 0, len(tracksRows))
	for _, row := range tracksRows {
		tracks = append(tracks, views.AdminTrack{
			ID:          row.ID,
			Title:       row.Title,
			Genre:       row.Genre,
			ReleaseDate: row.ReleaseDate,
		})
	}

	projectsRows, err := h.store.ListProjects(c.Request().Context())
	if err != nil {
		return err
	}
	projects := make([]views.AdminProject, 0, len(projectsRows))
	for _, row := range projectsRows {
		projects = append(projects, views.AdminProject{
			ID:      row.ID,
			Name:    row.Name,
			RepoURL: row.RepoUrl,
		})
	}

	linksRows, err := h.store.ListSocialLinks(c.Request().Context())
	if err != nil {
		return err
	}

	links := make([]views.Link, 0, len(linksRows))
	for _, row := range linksRows {
		links = append(links, views.Link{
			ID:   row.ID,
			Name: row.Name,
			URL:  row.Url,
		})
	}

	return hrender.Templ(c, http.StatusOK, views.Layout("trigex.moe | Admin", views.AdminPage(posts, tracks, projects, links)))
}

func (h *PageHandlers) ServeEditBlogPage(c *echo.Context) error {
	row, err := h.store.GetBlogPostBySlug(c.Request().Context(), c.Param("slug"))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return echo.NewHTTPError(http.StatusNotFound, "blog post not found")
		}
		return err
	}

	bodyHTML, err := hrender.MarkdownHTML(row.Body)
	if err != nil {
		return err
	}

	post := views.BlogPost{
		Slug:        row.Slug,
		Title:       row.Title,
		Excerpt:     row.Excerpt,
		Body:        row.Body,
		BodyHTML:    bodyHTML,
		PublishedAt: row.PublishedAt,
	}

	return hrender.Templ(c, http.StatusOK, views.Layout("trigex.moe | Edit Blog Post", views.BlogEditPage(post)))
}

func (h *PageHandlers) ServeBlogPreview(c *echo.Context) error {
	post, err := blogPreviewFromRequest(c)
	if err != nil {
		return err
	}

	bodyHTML, err := hrender.MarkdownHTML(post.Body)
	if err != nil {
		return err
	}
	post.BodyHTML = bodyHTML

	return hrender.Templ(c, http.StatusOK, views.BlogPreview(post))
}

func (h *PageHandlers) UploadBlogImage(c *echo.Context) error {
	imageURL, err := h.saveUploadedImage(c, "image", "blog")
	if err != nil {
		return c.HTML(http.StatusOK, fmt.Sprintf(
			`<p class="text-error">Upload failed: %s</p>`,
			html.EscapeString(err.Error()),
		))
	}
	if imageURL == "" {
		return c.HTML(http.StatusOK, `<p class="text-error">Upload failed: image is required.</p>`)
	}

	markdown := fmt.Sprintf("![image](%s)", imageURL)
	resp := fmt.Sprintf(
		"<div class=\"space-y-2\"><p class=\"text-success\">Uploaded: <a class=\"link\" href=\"%s\" target=\"_blank\">%s</a></p><p class=\"text-xs text-base-content/70\">Markdown</p><pre class=\"rounded-box bg-base-300 p-2 text-xs\">%s</pre><img src=\"%s\" alt=\"Uploaded blog image\" class=\"max-h-48 rounded-box\"/></div>",
		html.EscapeString(imageURL),
		html.EscapeString(imageURL),
		html.EscapeString(markdown),
		html.EscapeString(imageURL),
	)

	return c.HTML(http.StatusOK, resp)
}

func (h *PageHandlers) CreateBlogPost(c *echo.Context) error {
	post, err := blogPostFromRequest(c)
	if err != nil {
		return err
	}

	bodyHTML, err := hrender.MarkdownHTML(post.Body)
	if err != nil {
		return err
	}
	post.BodyHTML = bodyHTML

	if err := h.store.InsertBlogPost(c.Request().Context(), content.InsertBlogPostParams{
		ID:          slugOrDefault(post.Slug, post.Title),
		Slug:        slugOrDefault(post.Slug, post.Title),
		Title:       post.Title,
		Excerpt:     post.Excerpt,
		Body:        post.Body,
		PublishedAt: post.PublishedAt,
	}); err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, "/admin/")
}

func (h *PageHandlers) UpdateBlogPost(c *echo.Context) error {
	oldSlug := c.Param("slug")
	post, err := blogPostFromRequest(c)
	if err != nil {
		return err
	}

	newSlug := slugOrDefault(post.Slug, post.Title)
	rowsAffected, err := h.store.UpdateBlogPostBySlug(c.Request().Context(), content.UpdateBlogPostBySlugParams{
		Slug:        newSlug,
		Title:       post.Title,
		Excerpt:     post.Excerpt,
		Body:        post.Body,
		PublishedAt: post.PublishedAt,
		Slug_2:      oldSlug,
	})
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "blog post not found")
	}

	return c.Redirect(http.StatusSeeOther, "/admin/blog/"+newSlug+"/edit")
}

func (h *PageHandlers) DeleteBlogPost(c *echo.Context) error {
	rows, err := h.store.DeleteBlogPostBySlug(c.Request().Context(), c.Param("slug"))
	if err != nil {
		return err
	}
	if rows == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "blog post not found")
	}
	return c.Redirect(http.StatusSeeOther, "/admin/")
}

func (h *PageHandlers) CreateSocialLink(c *echo.Context) error {
	name := strings.TrimSpace(c.Request().FormValue("name"))
	if name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "name is required")
	}
	url := strings.TrimSpace(c.Request().FormValue("url"))
	if url == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "url is required")
	}

	id := strings.TrimSpace(c.Request().FormValue("id"))
	if id == "" {
		id = slugify(name)
	}

	if err := h.store.CreateSocialLink(c.Request().Context(), content.CreateSocialLinkParams{
		ID:   id,
		Name: name,
		Url:  url,
	}); err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, "/admin/")
}

func (h *PageHandlers) UpdateSocialLink(c *echo.Context) error {
	name := strings.TrimSpace(c.Request().FormValue("name"))
	if name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "name is required")
	}
	url := strings.TrimSpace(c.Request().FormValue("url"))
	if url == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "url is required")
	}

	rows, err := h.store.UpdateSocialLinkByID(c.Request().Context(), content.UpdateSocialLinkByIDParams{
		Name: name,
		Url:  url,
		ID:   c.Param("id"),
	})
	if err != nil {
		return err
	}
	if rows == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "social link not found")
	}

	return c.Redirect(http.StatusSeeOther, "/admin/")
}

func (h *PageHandlers) DeleteSocialLink(c *echo.Context) error {
	rows, err := h.store.DeleteSocialLinkByID(c.Request().Context(), c.Param("id"))
	if err != nil {
		return err
	}
	if rows == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "social link not found")
	}
	return c.Redirect(http.StatusSeeOther, "/admin/")
}

func (h *PageHandlers) CreateMusicTrack(c *echo.Context) error {
	title := strings.TrimSpace(c.Request().FormValue("title"))
	if title == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "title is required")
	}

	releaseDate, err := mustDate(c.Request().FormValue("release_date"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	coverImage := strings.TrimSpace(c.Request().FormValue("cover_image"))
	uploadedCover, err := h.saveUploadedImage(c, "cover_file", "covers")
	if err != nil {
		return err
	}
	if uploadedCover != "" {
		coverImage = uploadedCover
	}

	if err := h.store.InsertMusicTrack(c.Request().Context(), content.InsertMusicTrackParams{
		ID:            slugify(title),
		Title:         title,
		Genre:         strings.TrimSpace(c.Request().FormValue("genre")),
		FlacUrl:       strings.TrimSpace(c.Request().FormValue("flac_url")),
		Mp3Url:        strings.TrimSpace(c.Request().FormValue("mp3_url")),
		SoundcloudUrl: strings.TrimSpace(c.Request().FormValue("soundcloud_url")),
		YoutubeUrl:    strings.TrimSpace(c.Request().FormValue("youtube_url")),
		CoverImage:    coverImage,
		ReleaseDate:   releaseDate,
	}); err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, "/admin/")
}

func (h *PageHandlers) DeleteMusicTrack(c *echo.Context) error {
	rows, err := h.store.DeleteMusicTrackByID(c.Request().Context(), c.Param("id"))
	if err != nil {
		return err
	}
	if rows == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "music track not found")
	}
	return c.Redirect(http.StatusSeeOther, "/admin/")
}

func (h *PageHandlers) CreateProject(c *echo.Context) error {
	name := strings.TrimSpace(c.Request().FormValue("name"))
	if name == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "name is required")
	}

	if err := h.store.InsertProject(c.Request().Context(), content.InsertProjectParams{
		ID:          slugify(name),
		Name:        name,
		Description: strings.TrimSpace(c.Request().FormValue("description")),
		RepoUrl:     strings.TrimSpace(c.Request().FormValue("repo_url")),
		TechStack:   strings.TrimSpace(c.Request().FormValue("tech_stack")),
	}); err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, "/admin/")
}

func (h *PageHandlers) DeleteProject(c *echo.Context) error {
	rows, err := h.store.DeleteProjectByID(c.Request().Context(), c.Param("id"))
	if err != nil {
		return err
	}
	if rows == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "project not found")
	}
	return c.Redirect(http.StatusSeeOther, "/admin/")
}

func blogPostFromRequest(c *echo.Context) (views.BlogPost, error) {
	title := strings.TrimSpace(c.Request().FormValue("title"))
	if title == "" {
		return views.BlogPost{}, echo.NewHTTPError(http.StatusBadRequest, "title is required")
	}

	publishedAt, err := mustDate(c.Request().FormValue("published_at"))
	if err != nil {
		return views.BlogPost{}, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return views.BlogPost{
		Slug:        strings.TrimSpace(c.Request().FormValue("slug")),
		Title:       title,
		Excerpt:     strings.TrimSpace(c.Request().FormValue("excerpt")),
		Body:        c.Request().FormValue("body"),
		PublishedAt: publishedAt,
		BodyHTML:    "",
	}, nil
}

func blogPreviewFromRequest(c *echo.Context) (views.BlogPost, error) {
	publishedAt, err := mustDate(c.Request().FormValue("published_at"))
	if err != nil {
		return views.BlogPost{}, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	title := strings.TrimSpace(c.Request().FormValue("title"))
	if title == "" {
		title = "Untitled draft"
	}

	return views.BlogPost{
		Slug:        strings.TrimSpace(c.Request().FormValue("slug")),
		Title:       title,
		Excerpt:     strings.TrimSpace(c.Request().FormValue("excerpt")),
		Body:        c.Request().FormValue("body"),
		PublishedAt: publishedAt,
	}, nil
}

func slugOrDefault(slug, fallback string) string {
	slug = strings.TrimSpace(slug)
	if slug != "" {
		return slugify(slug)
	}
	return slugify(fallback)
}

func slugify(value string) string {
	value = strings.ToLower(strings.TrimSpace(value))
	var b strings.Builder
	b.Grow(len(value))

	lastDash := false
	for _, r := range value {
		switch {
		case r >= 'a' && r <= 'z':
			b.WriteRune(r)
			lastDash = false
		case r >= '0' && r <= '9':
			b.WriteRune(r)
			lastDash = false
		default:
			if !lastDash {
				b.WriteByte('-')
				lastDash = true
			}
		}
	}

	out := strings.Trim(b.String(), "-")
	if out == "" {
		return "entry"
	}
	return out
}

func mustDate(value string) (time.Time, error) {
	value = strings.TrimSpace(value)
	if value == "" {
		return time.Now().UTC(), nil
	}
	d, err := time.Parse("2006-01-02", value)
	if err != nil {
		return time.Time{}, fmt.Errorf("invalid date %q: %w", value, err)
	}
	return d.UTC(), nil
}

func (h *PageHandlers) saveUploadedImage(c *echo.Context, fieldName, subdir string) (string, error) {
	file, fileHeader, err := c.Request().FormFile(fieldName)
	if errors.Is(err, http.ErrMissingFile) {
		return "", nil
	}
	if err != nil {
		return "", echo.NewHTTPError(http.StatusBadRequest, "invalid uploaded file")
	}
	defer file.Close()

	ext, err := validatedImageExtension(file, fileHeader.Filename)
	if err != nil {
		return "", echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	base := slugify(strings.TrimSuffix(fileHeader.Filename, filepath.Ext(fileHeader.Filename)))
	if base == "" {
		base = "image"
	}
	filename := fmt.Sprintf("%s-%d%s", base, time.Now().UnixNano(), ext)

	uploadDir := filepath.Join(h.uploadsDir, subdir)
	if err := os.MkdirAll(uploadDir, 0o755); err != nil {
		return "", err
	}

	dstPath := filepath.Join(uploadDir, filename)
	dst, err := os.OpenFile(dstPath, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0o644)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err := file.Seek(0, io.SeekStart); err != nil {
		return "", err
	}
	if _, err := io.Copy(dst, file); err != nil {
		return "", err
	}

	return "/uploads/" + subdir + "/" + filename, nil
}

func validatedImageExtension(file multipart.File, originalFilename string) (string, error) {
	buf := make([]byte, 512)
	n, err := file.Read(buf)
	if err != nil && !errors.Is(err, io.EOF) {
		return "", err
	}
	contentType := http.DetectContentType(buf[:n])

	switch contentType {
	case "image/jpeg":
		return ".jpg", nil
	case "image/jpg":
		return ".jpg", nil
	case "image/png":
		return ".png", nil
	case "image/gif":
		return ".gif", nil
	case "image/webp":
		return ".webp", nil
	default:
		switch strings.ToLower(filepath.Ext(originalFilename)) {
		case ".jpg", ".jpeg":
			return ".jpg", nil
		case ".png":
			return ".png", nil
		case ".gif":
			return ".gif", nil
		case ".webp":
			return ".webp", nil
		default:
			return "", fmt.Errorf("unsupported image type: %s", contentType)
		}
	}
}
