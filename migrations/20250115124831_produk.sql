-- +goose Up
-- +goose StatementBegin
CREATE TABLE produk (
	id_produk   SERIAL PRIMARY KEY,
	nama_produk TEXT NOT NULL,
	harga   DECIMAL(10, 2) NOT NULL,
	kategori_id  BIGSERIAL NOT NULL,
	status_id  BIGSERIAL NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT NOW(),
	constraint fk_produk_kategori
     foreign key (kategori_id) 
     REFERENCES kategori (id_kategori),
	constraint fk_produk_status
     foreign key (status_id) 
     REFERENCES status (id_status)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE produk;
-- +goose StatementEnd
