-- name: ListMusicTracks :many
SELECT id, title, genre, flac_url, mp3_url, soundcloud_url, youtube_url, cover_image, release_date
FROM music_tracks
ORDER BY release_date DESC, title ASC;

-- name: CreateMusicTrack :exec
INSERT OR IGNORE INTO music_tracks (
    id, title, genre, flac_url, mp3_url, soundcloud_url, youtube_url, cover_image, release_date
) VALUES (
    ?, ?, ?, ?, ?, ?, ?, ?, ?
);
