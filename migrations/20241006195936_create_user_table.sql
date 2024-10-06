-- +goose Up
-- +goose StatementBegin
CREATE TABLE "users"
(
    "id"         bigserial,
    "username"   varchar(255) UNIQUE NOT NULL ,
    "password"   varchar(255) NOT NULL,
    "created_at" timestamptz  DEFAULT NOW(),
    "updated_at" timestamptz  DEFAULT NOW(),
    PRIMARY KEY ("id")
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "users";
-- +goose StatementEnd
