-- +goose Up
-- +goose StatementBegin
ALTER TABLE media
ADD recomended BOOLEAN DEFAULT false;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE media
DROP COLUMN recomended;
-- +goose StatementEnd
