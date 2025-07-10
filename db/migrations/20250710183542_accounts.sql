-- +goose Up
CREATE TABLE IF NOT EXISTS accounts (
    account_id BIGINT PRIMARY KEY,
    balance NUMERIC(20,8) NOT NULL
);


-- +goose Down
DROP TABLE IF EXISTS accounts;
