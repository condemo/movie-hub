-- +goose Up
-- +goose StatementBegin
ALTER TABLE media
ADD first_air SMALLINT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE media
DROP COLUMN first_air;
-- +goose StatementEnd
