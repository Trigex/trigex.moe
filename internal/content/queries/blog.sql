-- name: ListBlogPosts :many
SELECT id, slug, title, excerpt, body, published_at
FROM blog_posts
ORDER BY published_at DESC, title ASC;

-- name: GetBlogPostBySlug :one
SELECT id, slug, title, excerpt, body, published_at
FROM blog_posts
WHERE slug = ?
LIMIT 1;

-- name: CreateBlogPost :exec
INSERT OR IGNORE INTO blog_posts (
    id, slug, title, excerpt, body, published_at
) VALUES (
    ?, ?, ?, ?, ?, ?
);
