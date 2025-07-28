-- +goose Up
-- +goose StatementBegin
ALTER TABLE media
ADD runtime SMALLINT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE media
DROP COLUMN runtime;
-- +goose StatementEnd
