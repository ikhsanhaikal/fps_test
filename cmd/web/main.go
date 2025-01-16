package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"ikhsanhaikal.com/fastprint-test/pgdb"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	dbUrl := os.Getenv("GOOSE_DBSTRING")

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, dbUrl)
	if err != nil {
		panic(err.Error())
	}

	defer conn.Close(ctx)

	if err := conn.Ping(ctx); err != nil {
		panic(err.Error())
	}

	queries := pgdb.New(conn)

	// initialize(queries) TODO: check if already populated

	r := gin.Default()

	r.Use(func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Next()
	})

	r.GET("/welcome", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Greetings",
		})
	})

	r.GET("/produk", func(ctx *gin.Context) {
		q1 := ctx.DefaultQuery("perPage", "5")
		q2 := ctx.DefaultQuery("page", "0")

		perPage, err := strconv.ParseInt(q1, 10, 64)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"errors": err.Error(),
			})
			return
		}

		page, err := strconv.ParseInt(q2, 10, 64)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"errors": err.Error(),
			})
			return
		}

		fmt.Printf("page: %d, perPage: %d\n", page, perPage)

		produk_plural, err := queries.ListProduk(ctx, pgdb.ListProdukParams{
			Limit:  int32(perPage),
			Offset: int32((page - 1) * perPage),
		})

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"errors": err.Error(),
				"data":   nil,
			})
		}

		total, _ := queries.TotalProduk(ctx)

		ctx.JSON(http.StatusOK, gin.H{
			"errors": nil,
			"data":   produk_plural,
			"total":  total,
		})
	})

	r.POST("/produk", create_produk_handler(queries))
	r.DELETE("/produk/:id", delete_produk_handler(queries))
	r.PUT("/produk/:id", update_produk_handler(queries))

	r.Run("localhost:3000")
}

// 4. Buat halaman untuk menampilkan data yang sudah anda simpan
// 5. Lalu tampilkan data yang hanya memiliki status " bisa dijual " this should be done on react
// 7. Untuk fitur tambah dan edit gunakan form validasi (inputan nama harus diisi, dan harga harus berupa inputan angka)
