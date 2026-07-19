-- name: ListSocialLinks :many
SELECT id, name, url
FROM social_links
ORDER BY id ASC;

-- name: CreateSocialLink :exec
INSERT OR IGNORE INTO social_links (
    id, name, url
) VALUES (
    ?, ?, ?
);
