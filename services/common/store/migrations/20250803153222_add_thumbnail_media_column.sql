-- +goose Up
-- +goose StatementBegin
ALTER TABLE media
ADD thumbnail TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE media
DROP COLUMN thumbnail;
-- +goose StatementEnd
