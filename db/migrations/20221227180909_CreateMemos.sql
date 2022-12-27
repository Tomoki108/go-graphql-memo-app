-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE memos
(
    id SERIAL,
    title varchar(255) NOT NULL,
    body text NOT NULL,
    PRIMARY KEY(id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE memos;