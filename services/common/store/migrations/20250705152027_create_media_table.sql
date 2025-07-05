-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS media (
	id SERIAL PRIMARY KEY,
	type varchar(15) NOT NULL,
	title varchar(30) NOT NULL,
	year smallint NOT NULL,
	genres text NOT NULL,
	seasons smallint NULL,
	caps smallint NULL,
	description text NOT NULL,
	rating smallint NOT NULL,
	image text,
	fav boolean DEFAULT false,
	viewed boolean DEFAULT false,
	UNIQUE(title)
	);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE media;
-- +goose StatementEnd
