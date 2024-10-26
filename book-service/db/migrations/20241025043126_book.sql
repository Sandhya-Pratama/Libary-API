-- +goose Up
CREATE TABLE public.books (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    isbn VARCHAR(20) UNIQUE NOT NULL,
    author_id INT NOT NULL,
    category_id INT NOT NULL,
    stock INT DEFAULT 0
);

-- +goose Down
DROP TABLE IF EXISTS public.books;