-- name: ListVenues :many
SELECT id, name, location, region, link
FROM venues
WHERE (? = '' OR name = ?);

-- name: GetVenueByID :one
SELECT id, name, location, region, link
FROM venues
WHERE id = ? LIMIT 1;

-- name: ListArtists :many
SELECT id, name, dj_name, label, location, region, genre, link
FROM artists
WHERE (? = '' OR name = ?);

-- name: GetArtistByID :one
SELECT id, name, dj_name, label, location, region, genre, link
FROM artists
WHERE id = ? LIMIT 1;

-- name: ListLabels :many
SELECT id, name, location, region, genre, link
FROM labels
WHERE (? = '' OR name = ?);

-- name: GetLabelByID :one
SELECT id, name, location, region, genre, link
FROM labels
WHERE id = ? LIMIT 1;

-- name: ListGenres :many
SELECT id, name
FROM genres
WHERE (? = '' OR name = ?);

-- name: GetGenreByID :one
SELECT id, name
FROM genres
WHERE id = ? LIMIT 1;
