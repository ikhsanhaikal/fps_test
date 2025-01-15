-- name: ListProduk :many
SELECT * FROM produk;

-- name: CreateProduk :one
INSERT INTO produk (
  nama_produk, harga, kategori_id, status_id
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: CreateKategori :one
INSERT INTO kategori (
  nama_kategori 
) VALUES (
  $1 
)
RETURNING *;

-- name: CreateStatus :one
INSERT INTO status (
  nama_status
) VALUES (
  $1
)
RETURNING *;