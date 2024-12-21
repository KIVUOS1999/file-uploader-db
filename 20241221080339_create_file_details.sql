-- +goose Up
-- +goose StatementBegin
CREATE TABLE file_details (
    file_id CHARACTER VARYING(255) PRIMARY KEY,
    file_name CHARACTER VARYING(255) NOT NULL,
    file_size BIGINT,
    total_chunks BIGINT,
    created_at BIGINT,
    user_id CHARACTER VARYING(255) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE file_details;
-- +goose StatementEnd
