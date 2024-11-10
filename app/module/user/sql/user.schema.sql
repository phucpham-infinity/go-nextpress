CREATE TABLE IF NOT EXISTS users (
    id UUID DEFAULT gen_random_uuid(),
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    activation_key VARCHAR(50) NOT NULL,
    status user_status DEFAULT 'inactive' NOT NULL,
    created_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    updated_at TIMESTAMPTZ DEFAULT now() NOT NULL,
    registered_at TIMESTAMPTZ NULL,
    deleted_at TIMESTAMPTZ NULL
);