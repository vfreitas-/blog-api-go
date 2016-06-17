
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE categories (
    id bigserial PRIMARY key,
    name varchar(50) NOT NULL,
    parent_id integer NULL,
    created_at TIMESTAMP DEFAULT NULL,
    updated_at TIMESTAMP DEFAULT NULL,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE categories;
