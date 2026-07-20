package content

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	_ "modernc.org/sqlite"
)

const DefaultDBPath = "data/trigexmoe.sqlite"

//go:embed schema.sql
var schemaFS embed.FS

type Store struct {
	*Queries
	db *sql.DB
}

func Open(ctx context.Context, path string) (*Store, error) {
	if path == "" {
		path = DefaultDBPath
	}

	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return nil, fmt.Errorf("creating db directory: %w", err)
	}

	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, fmt.Errorf("opening sqlite db: %w", err)
	}

	db.SetMaxOpenConns(1)

	if err := configureSQLite(ctx, db, path); err != nil {
		_ = db.Close()
		return nil, err
	}

	schema, err := schemaFS.ReadFile("schema.sql")
	if err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("reading schema: %w", err)
	}

	if _, err := db.ExecContext(ctx, string(schema)); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("applying schema: %w", err)
	}

	if err := ensureMusicTrackGenreColumn(ctx, db); err != nil {
		_ = db.Close()
		return nil, err
	}

	store := &Store{
		Queries: New(db),
		db:      db,
	}

	if err := store.seed(ctx); err != nil {
		_ = db.Close()
		return nil, err
	}

	return store, nil
}

func (s *Store) Close() error {
	return s.db.Close()
}

func configureSQLite(ctx context.Context, db *sql.DB, path string) error {
	if _, err := db.ExecContext(ctx, "PRAGMA foreign_keys = ON"); err != nil {
		return fmt.Errorf("enabling sqlite foreign keys: %w", err)
	}
	if _, err := db.ExecContext(ctx, "PRAGMA busy_timeout = 5000"); err != nil {
		return fmt.Errorf("setting sqlite busy timeout: %w", err)
	}

	// FreeBSD default uses /var/db/trigexmoe.sqlite. /var/db is often not writable
	// by service users for creating journal side files, so force in-memory journaling.
	cleaned := filepath.Clean(path)
	if runtime.GOOS == "freebsd" && strings.HasPrefix(cleaned, "/var/db/") {
		if _, err := db.ExecContext(ctx, "PRAGMA journal_mode = MEMORY"); err != nil {
			return fmt.Errorf("setting sqlite journal mode: %w", err)
		}
		if _, err := db.ExecContext(ctx, "PRAGMA temp_store = MEMORY"); err != nil {
			return fmt.Errorf("setting sqlite temp store: %w", err)
		}
	}

	return nil
}

func ensureMusicTrackGenreColumn(ctx context.Context, db *sql.DB) error {
	rows, err := db.QueryContext(ctx, "PRAGMA table_info(music_tracks)")
	if err != nil {
		return fmt.Errorf("querying music_tracks columns: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			cid        int
			name       string
			columnType string
			notNull    int
			defaultVal sql.NullString
			pk         int
		)
		if err := rows.Scan(&cid, &name, &columnType, &notNull, &defaultVal, &pk); err != nil {
			return fmt.Errorf("scanning music_tracks column info: %w", err)
		}
		if name == "genre" {
			return nil
		}
	}
	if err := rows.Err(); err != nil {
		return fmt.Errorf("iterating music_tracks columns: %w", err)
	}

	if _, err := db.ExecContext(ctx, `ALTER TABLE music_tracks ADD COLUMN genre TEXT NOT NULL DEFAULT ''`); err != nil {
		return fmt.Errorf("adding music_tracks.genre column: %w", err)
	}
	return nil
}
