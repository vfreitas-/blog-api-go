
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE users (
    id bigserial PRIMARY key,
    login varchar(50) NOT NULL,
    password varchar(255) NOT NULL,
    name varchar(50) NOT NULL,
    created_at TIMESTAMP DEFAULT NULL,
    updated_at TIMESTAMP DEFAULT NULL,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE users;
