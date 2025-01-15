-- +goose Up
-- +goose StatementBegin
CREATE TABLE kategori (
	id_kategori   SERIAL PRIMARY KEY,
	nama_kategori VARCHAR(255) UNIQUE NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE kategori;
-- +goose StatementEnd
