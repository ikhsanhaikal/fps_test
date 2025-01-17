-- name: TotalProduk :one
SELECT COUNT(*) AS total FROM produk;

-- name: GetKategoriById :one
SELECT * from kategori 
WHERE id_kategori = $1;

-- name: GetKategoriByIds :many
SELECT id_kategori as id , nama_kategori from kategori 
WHERE id_kategori = ANY(@ids::int[]); 

-- name: GetStatusByIds :many
SELECT id_status as id , nama_status from status 
WHERE id_status = ANY(@ids::int[]); 

-- name: ListProduk :many
SELECT id_produk as id, nama_produk, harga, kategori_id, status_id, created_at FROM produk
ORDER BY id LIMIT $1 OFFSET $2;

-- name: ListKategori :many
SELECT id_kategori as id, nama_kategori FROM kategori 
ORDER BY id LIMIT $1 OFFSET $2;

-- name: ListStatus :many
SELECT id_status as id, nama_status FROM status 
ORDER BY id LIMIT $1 OFFSET $2;

-- name: CreateProduk :one
INSERT INTO produk (
  nama_produk, harga, kategori_id, status_id
) VALUES (
  $1, $2, $3, $4
)
RETURNING id_produk as id, nama_produk, harga, kategori_id, status_id; 

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

-- name: DeleteProduk :one
DELETE FROM produk
WHERE id_produk = $1
RETURNING id_produk as id, nama_produk, harga, kategori_id, status_id;

-- name: UpdateProduk :one
UPDATE produk SET
  nama_produk = COALESCE(sqlc.narg('nama_produk'), nama_produk),
  harga = COALESCE(sqlc.narg('harga'), harga),
  kategori_id = COALESCE(sqlc.narg('kategori_id'), kategori_id),
  status_id = COALESCE(sqlc.narg('status_id'), status_id)
WHERE id_produk = $1
RETURNING *;
