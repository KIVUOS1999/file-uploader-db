-- +goose Up
-- +goose StatementBegin
CREATE TABLE chunk_details (
    chunk_id CHARACTER VARYING(255) PRIMARY KEY,
    file_id CHARACTER VARYING(255) NOT NULL,
    check_sum CHARACTER VARYING(255),
    chunk_order INTEGER NOT NULL,
    created_at BIGINT
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE chunk_details;
-- +goose StatementEnd
