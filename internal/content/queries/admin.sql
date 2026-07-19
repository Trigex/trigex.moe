-- name: InsertMusicTrack :exec
INSERT INTO music_tracks (
    id, title, genre, flac_url, mp3_url, soundcloud_url, youtube_url, cover_image, release_date
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?
);

-- name: InsertProject :exec
INSERT INTO projects (
    id, name, description, repo_url, tech_stack
) VALUES (
    ?, ?, ?, ?, ?
);

-- name: InsertBlogPost :exec
INSERT INTO blog_posts (
    id, slug, title, excerpt, body, published_at
) VALUES (
    ?, ?, ?, ?, ?, ?
);

-- name: UpdateBlogPostBySlug :execrows
UPDATE blog_posts
SET
    slug = ?,
    title = ?,
    excerpt = ?,
    body = ?,
    published_at = ?
WHERE slug = ?;

-- name: DeleteBlogPostBySlug :execrows
DELETE FROM blog_posts
WHERE slug = ?;

-- name: UpdateSocialLinkByID :execrows
UPDATE social_links
SET
    name = ?,
    url = ?
WHERE id = ?;

-- name: DeleteSocialLinkByID :execrows
DELETE FROM social_links
WHERE id = ?;

-- name: DeleteMusicTrackByID :execrows
DELETE FROM music_tracks
WHERE id = ?;

-- name: DeleteProjectByID :execrows
DELETE FROM projects
WHERE id = ?;
