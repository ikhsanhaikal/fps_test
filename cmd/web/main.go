package main

import (
	"context"
	"net/http"
	"os"

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

	// initialize(queries)

	r := gin.Default()

	r.GET("/welcome", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Greetings",
		})
	})

	r.POST("/produk", create_produk_handler(queries))
	r.DELETE("/produk/:id", delete_produk_handler(queries))

	r.Run("localhost:3000")
}

// 3. Simpan produk yang sudah anda dapatkan dari url produk
// 4. Buat halaman untuk menampilkan data yang sudah anda simpan
// 5. Lalu tampilkan data yang hanya memiliki status " bisa dijual " this should be done on react
// 6. Buat fitur untuk edit, tambah dan hapus
// 7. Untuk fitur tambah dan edit gunakan form validasi (inputan nama harus diisi, dan harga harus berupa inputan angka)
