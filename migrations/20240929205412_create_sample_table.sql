-- +goose Up
-- +goose StatementBegin
CREATE TABLE "samples"
(
    "id"         bigserial,
    "name"       varchar(255) NOT NULL,
    "created_at" timestamptz  DEFAULT NOW(),
    "updated_at" timestamptz  DEFAULT NOW(),
    PRIMARY KEY ("id")
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "samples";
-- +goose StatementEnd
