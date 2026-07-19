-- name: ListProjects :many
SELECT id, name, description, repo_url, tech_stack
FROM projects
ORDER BY id ASC;

-- name: CreateProject :exec
INSERT OR IGNORE INTO projects (
    id, name, description, repo_url, tech_stack
) VALUES (
    ?, ?, ?, ?, ?
);
