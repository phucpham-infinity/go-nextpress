-- +goose Up
CREATE TYPE user_status AS ENUM ('active', 'inactive');

CREATE TABLE IF NOT EXISTS users (
    uuid UUID NOT NULL,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    activation_key VARCHAR(50) NULL,
    status user_status DEFAULT 'inactive' NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    registered_at TIMESTAMPTZ NULL,
    deleted_at TIMESTAMPTZ NULL
);

-- +goose Down
DROP TYPE user_status;
DROP TABLE IF EXISTS users;