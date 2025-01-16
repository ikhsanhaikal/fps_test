package main

import (
	"fmt"
	"math/big"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
	"ikhsanhaikal.com/fastprint-test/pgdb"
)

type CreateProdukRequest struct {
	Nama       string `form:"nama" json:"nama" xml:"nama"  binding:"required"`
	Harga      int    `form:"harga" json:"harga" xml:"harga" binding:"required"`
	KategoriID int    `form:"kategori_id" json:"kategori_id" xml:"kategori_id" binding:"required"`
	StatusID   int    `form:"status_id" json:"status_id" xml:"status_id" binding:"required"`
}

func create_produk_handler(queries *pgdb.Queries) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var json CreateProdukRequest
		if err := ctx.ShouldBindJSON(&json); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Printf("body request: %+v\n", json)

		produk, err := queries.CreateProduk(ctx, pgdb.CreateProdukParams{
			NamaProduk: json.Nama,
			Harga:      pgtype.Numeric{Exp: 0, Int: big.NewInt(int64(json.Harga)), NaN: false, Valid: true},
			KategoriID: int64(json.KategoriID),
			StatusID:   int64(json.StatusID),
		})

		if err != nil {
			ctx.JSONP(http.StatusInternalServerError, gin.H{
				"errors": err.Error(),
			})
			return
		}
		ctx.JSONP(http.StatusOK, gin.H{
			"errors": nil,
			"data":   produk,
		})
	}
}

type ProdukUri struct {
	ID int32 `uri:"id" binding:"required"`
}

func delete_produk_handler(queries *pgdb.Queries) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var produkUri ProdukUri

		if err := ctx.ShouldBindUri(&produkUri); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
			return
		}

		fmt.Printf("produk_id: %#v\n", produkUri)

		if err := queries.DeleteProduk(ctx, produkUri.ID); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"errors": err.Error(),
			})
			return
		}
		ctx.Status(http.StatusNoContent)
	}
}

func update_produk_handler(queries *pgdb.Queries) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var produkUri ProdukUri

		if err := ctx.ShouldBindUri(&produkUri); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
			return
		}

		body := struct {
			Nama       *string `form:"nama" json:"nama" xml:"nama"  binding:"omitempty"`
			Harga      *int    `form:"harga" json:"harga" xml:"harga" binding:"omitempty"`
			KategoriID *int    `form:"kategori_id" json:"kategori_id" xml:"kategori_id" binding:"omitempty"`
			StatusID   *int    `form:"status_id" json:"status_id" xml:"status_id" binding:"omitempty"`
		}{}

		if err := ctx.ShouldBindJSON(&body); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		fmt.Printf("produk_id: %#v\n", produkUri)

		var namaProduk pgtype.Text
		var harga pgtype.Numeric
		var kategoriId pgtype.Int8
		var statusId pgtype.Int8

		if body.Nama != nil {
			namaProduk.String = *body.Nama
			namaProduk.Valid = true
		}
		if body.Harga != nil {
			harga.Int = big.NewInt(int64(*body.Harga))
			harga.Exp = 0
			harga.NaN = false
			harga.Valid = true
		}
		if body.KategoriID != nil {
			kategoriId.Int64 = int64(*body.KategoriID)
			kategoriId.Valid = true
		}
		if body.StatusID != nil {
			statusId.Int64 = int64(*body.StatusID)
			statusId.Valid = true
		}

		produk, err := queries.UpdateProduk(ctx, pgdb.UpdateProdukParams{
			Id:         produkUri.ID,
			NamaProduk: namaProduk,
			Harga:      harga,
			KategoriID: kategoriId,
			StatusID:   statusId,
		})

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"errors": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"errors": nil,
			"data":   produk,
		})

	}
}
