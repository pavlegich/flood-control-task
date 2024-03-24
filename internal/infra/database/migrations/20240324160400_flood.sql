-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE IF NOT EXISTS flood (
    id serial PRIMARY KEY,
    user_id bigint NOT NULL,
    created_at timestamp DEFAULT NOW()
);

-- create indexes
CREATE INDEX IF NOT EXISTS flood_user_id_idx ON flood (user_id);
CREATE INDEX IF NOT EXISTS flood_created_at_idx ON flood (created_at);

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

DROP INDEX flood_user_id_idx;
DROP INDEX flood_created_at_idx;
DROP TABLE flood;
