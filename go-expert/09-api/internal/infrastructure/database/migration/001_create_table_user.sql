-- +goose Up
-- +goose StatementBegin
CREATE TABLE user
(
    id       CHAR(36) PRIMARY KEY,
    name     VARCHAR(255),
    email    VARCHAR(255),
    password VARCHAR(255)

);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user;
-- +goose StatementEnd
