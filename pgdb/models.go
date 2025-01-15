// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package pgdb

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Kategori struct {
	IDKategori   int32
	NamaKategori string
}

type Produk struct {
	IDProduk   int32
	NamaProduk string
	Harga      pgtype.Numeric
	KategoriID int64
	StatusID   int64
	CreatedAt  pgtype.Timestamp
}

type Status struct {
	IDStatus   int32
	NamaStatus string
}
