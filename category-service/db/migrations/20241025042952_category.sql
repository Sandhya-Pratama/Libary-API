-- +goose Up
CREATE TABLE public.category (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    bio TEXT
);

-- +goose Down
DROP TABLE IF EXISTS public.category;