CREATE TABLE IF NOT EXISTS music_tracks (
    id TEXT PRIMARY KEY,
    title TEXT NOT NULL,
    genre TEXT NOT NULL DEFAULT '',
    flac_url TEXT NOT NULL,
    mp3_url TEXT NOT NULL,
    soundcloud_url TEXT NOT NULL,
    youtube_url TEXT NOT NULL,
    cover_image TEXT NOT NULL,
    release_date TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS projects (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    repo_url TEXT NOT NULL,
    tech_stack TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS blog_posts (
    id TEXT PRIMARY KEY,
    slug TEXT NOT NULL UNIQUE,
    title TEXT NOT NULL,
    excerpt TEXT NOT NULL,
    body TEXT NOT NULL,
    published_at TIMESTAMP NOT NULL
);

CREATE TABLE IF NOT EXISTS social_links (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    url TEXT NOT NULL
);
