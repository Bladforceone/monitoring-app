-- +goose Up
-- +goose StatementBegin
CREATE TABLE results (
    id SERIAL PRIMARY KEY,
    url TEXT NOT NULL,
    status_code INT NOT NULL,
    duration_ms INT NOT NULL,
    error TEXT,
    created_at TIMESTAMP DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE results;
-- +goose StatementEnd
