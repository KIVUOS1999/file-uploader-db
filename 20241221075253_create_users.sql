-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id CHARACTER VARYING(255) PRIMARY KEY,
    name CHARACTER VARYING(255),
    email CHARACTER VARYING(255) NOT NULL UNIQUE,
    picture TEXT,
    created_at BIGINT NOT NULL,
    alloted_size BIGINT DEFAULT 10737418240
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
