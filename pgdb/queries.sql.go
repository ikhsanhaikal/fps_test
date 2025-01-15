// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: queries.sql

package pgdb

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createKategori = `-- name: CreateKategori :one
INSERT INTO kategori (
  nama_kategori 
) VALUES (
  $1 
)
RETURNING id_kategori, nama_kategori
`

func (q *Queries) CreateKategori(ctx context.Context, namaKategori string) (Kategori, error) {
	row := q.db.QueryRow(ctx, createKategori, namaKategori)
	var i Kategori
	err := row.Scan(&i.IDKategori, &i.NamaKategori)
	return i, err
}

const createProduk = `-- name: CreateProduk :one
INSERT INTO produk (
  nama_produk, harga, kategori_id, status_id
) VALUES (
  $1, $2, $3, $4
)
RETURNING id_produk, nama_produk, harga, kategori_id, status_id, created_at
`

type CreateProdukParams struct {
	NamaProduk string
	Harga      pgtype.Numeric
	KategoriID int64
	StatusID   int64
}

func (q *Queries) CreateProduk(ctx context.Context, arg CreateProdukParams) (Produk, error) {
	row := q.db.QueryRow(ctx, createProduk,
		arg.NamaProduk,
		arg.Harga,
		arg.KategoriID,
		arg.StatusID,
	)
	var i Produk
	err := row.Scan(
		&i.IDProduk,
		&i.NamaProduk,
		&i.Harga,
		&i.KategoriID,
		&i.StatusID,
		&i.CreatedAt,
	)
	return i, err
}

const createStatus = `-- name: CreateStatus :one
INSERT INTO status (
  nama_status
) VALUES (
  $1
)
RETURNING id_status, nama_status
`

func (q *Queries) CreateStatus(ctx context.Context, namaStatus string) (Status, error) {
	row := q.db.QueryRow(ctx, createStatus, namaStatus)
	var i Status
	err := row.Scan(&i.IDStatus, &i.NamaStatus)
	return i, err
}

const deleteProduk = `-- name: DeleteProduk :exec
DELETE FROM produk
WHERE id_produk = $1
RETURNING id_produk, nama_produk, harga, kategori_id, status_id, created_at
`

func (q *Queries) DeleteProduk(ctx context.Context, idProduk int32) error {
	_, err := q.db.Exec(ctx, deleteProduk, idProduk)
	return err
}

const listProduk = `-- name: ListProduk :many
SELECT id_produk, nama_produk, harga, kategori_id, status_id, created_at FROM produk
`

func (q *Queries) ListProduk(ctx context.Context) ([]Produk, error) {
	rows, err := q.db.Query(ctx, listProduk)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Produk
	for rows.Next() {
		var i Produk
		if err := rows.Scan(
			&i.IDProduk,
			&i.NamaProduk,
			&i.Harga,
			&i.KategoriID,
			&i.StatusID,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateProduk = `-- name: UpdateProduk :one
UPDATE produk SET
  nama_produk = COALESCE($2, nama_produk),
  harga = COALESCE($3, harga),
  kategori_id = COALESCE($4, kategori_id),
  status_id = COALESCE($5, status_id)
WHERE id_produk = $1
RETURNING id_produk, nama_produk, harga, kategori_id, status_id, created_at
`

type UpdateProdukParams struct {
	IDProduk   int32
	NamaProduk pgtype.Text
	Harga      pgtype.Numeric
	KategoriID pgtype.Int8
	StatusID   pgtype.Int8
}

func (q *Queries) UpdateProduk(ctx context.Context, arg UpdateProdukParams) (Produk, error) {
	row := q.db.QueryRow(ctx, updateProduk,
		arg.IDProduk,
		arg.NamaProduk,
		arg.Harga,
		arg.KategoriID,
		arg.StatusID,
	)
	var i Produk
	err := row.Scan(
		&i.IDProduk,
		&i.NamaProduk,
		&i.Harga,
		&i.KategoriID,
		&i.StatusID,
		&i.CreatedAt,
	)
	return i, err
}
