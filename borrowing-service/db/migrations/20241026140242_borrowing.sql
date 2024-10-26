-- +goose Up
CREATE TABLE public.borrowing_books (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    book_id BIGINT NOT NULL,
    borrow_date TIMESTAMP NOT NULL,
    return_date TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- +goose Down
DROP TABLE IF EXISTS public.borrowing_books;
