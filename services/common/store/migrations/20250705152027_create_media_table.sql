-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS media (
    id SERIAL PRIMARY KEY,
    media_type VARCHAR(15) NOT NULL,
    title VARCHAR(80) NOT NULL,
    release_year SMALLINT NOT NULL,
    genres TEXT NOT NULL,
    seasons SMALLINT NULL,
    caps SMALLINT NULL,
    description TEXT NOT NULL,
    rating SMALLINT NOT NULL,
    image TEXT,
    fav BOOLEAN DEFAULT false,
    viewed BOOLEAN DEFAULT false,
    UNIQUE (title)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE media;
-- +goose StatementEnd
