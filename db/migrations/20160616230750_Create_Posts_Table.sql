
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE posts (
    id bigserial PRIMARY key,
    title varchar(55) NOT NULL,
    content text NOT NULL,
    banner_image_id integer NOT NULL,
    created_at TIMESTAMP DEFAULT NULL,
    updated_at TIMESTAMP DEFAULT NULL,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE posts;
