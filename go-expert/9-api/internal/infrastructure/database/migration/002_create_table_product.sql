-- +goose Up
-- +goose StatementBegin
CREATE TABLE product
(
    id         CHAR(36) PRIMARY KEY,
    name       VARCHAR(255),
    price      FLOAT,
    created_at DATETIME(6)

);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE product;
-- +goose StatementEnd
