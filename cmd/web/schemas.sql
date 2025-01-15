CREATE TABLE produk (
	id_produk   SERIAL PRIMARY KEY,
	nama_produk TEXT NOT NULL,
	harga   DECIMAL(10, 2) NOT NULL,
	kategori_id  BIGSERIAL NOT NULL,
	status_id  BIGSERIAL NOT NULL,
	created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE kategori (
	id_kategori   SERIAL PRIMARY KEY,
	nama_kategori VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE status (
	id_status SERIAL PRIMARY KEY,
	nama_status VARCHAR(255) UNIQUE NOT NULL
);
