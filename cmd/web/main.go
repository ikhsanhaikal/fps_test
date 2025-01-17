package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"ikhsanhaikal.com/fastprint-test/pgdb"
)

type kategoriUri struct {
	ID int32 `uri:"id" binding:"required"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	dbUrl := os.Getenv("GOOSE_DBSTRING")

	ctx := context.Background()

	// conn, err := pgx.Connect(ctx, dbUrl)
	pool, err := pgxpool.New(ctx, dbUrl)

	if err != nil {
		panic(err.Error())
	}

	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		panic(err.Error())
	}

	queries := pgdb.New(pool)

	// initialize(queries) TODO: check if already populated

	r := gin.Default()

	r.Use(func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		ctx.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}

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
			return
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

	r.GET("/kategori", func(ctx *gin.Context) {

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

		arrayOfids := ctx.QueryArray("ids")

		fmt.Printf("arrayOfIds: %+v\n", arrayOfids)

		if len(arrayOfids) >= 1 {
			ids := []int32{}
			for _, v := range arrayOfids {
				id, _ := strconv.Atoi(v)
				ids = append(ids, int32(id))
			}
			kategori_plural, err := queries.GetKategoriByIds(ctx, ids)

			fmt.Printf("kategori_plural: %+v\n", kategori_plural)

			if err != nil {
				ctx.JSON(http.StatusOK, gin.H{
					"errors": err.Error(),
				})
				return
			}

			ctx.JSON(http.StatusOK, gin.H{
				"errors": nil,
				"data":   kategori_plural,
			})
			return
		}

		k, err := queries.ListKategori(ctx, pgdb.ListKategoriParams{
			Limit:  int32(perPage),
			Offset: int32((page - 1) * perPage),
		})

		if err != nil {
			fmt.Printf("errors: %+v\n", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"errors": err.Error(),
			})
			return
		}

		fmt.Printf("k: %+v\n", k)

		ctx.JSON(http.StatusOK, gin.H{
			"errors": nil,
			"data":   k,
			"total":  0,
		})
	})

	r.GET("/status", func(ctx *gin.Context) {
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
		arrayOfids := ctx.QueryArray("ids")

		fmt.Printf("arrayOfIds: %+v\n", arrayOfids)

		if len(arrayOfids) >= 1 {
			ids := []int32{}
			for _, v := range arrayOfids {
				id, _ := strconv.Atoi(v)
				ids = append(ids, int32(id))
			}
			status_plural, err := queries.GetStatusByIds(ctx, ids)

			fmt.Printf("status_plural: %+v\n", status_plural)

			if err != nil {
				ctx.JSON(http.StatusOK, gin.H{
					"errors": err.Error(),
				})
				return
			}

			ctx.JSON(http.StatusOK, gin.H{
				"errors": nil,
				"data":   status_plural,
			})
			return
		}

		s, err := queries.ListStatus(ctx, pgdb.ListStatusParams{
			Limit:  int32(perPage),
			Offset: int32((page - 1) * perPage),
		})

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"errors": err.Error(),
			})
			return
		}

		fmt.Printf("s: %+v\n", s)

		ctx.JSON(http.StatusOK, gin.H{
			"errors": nil,
			"data":   s,
			"total":  0,
		})
	})

	r.GET("/kategori/:id", func(ctx *gin.Context) {
		var kategoriUri kategoriUri

		if err := ctx.ShouldBindUri(&kategoriUri); err != nil {
			fmt.Printf("kategoriUri: %+v\n", kategoriUri)
			ctx.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
			fmt.Printf("error: %+v\n", err)
			return
		}

		k, err := queries.GetKategoriById(ctx, kategoriUri.ID)

		if err != nil {
			fmt.Printf("kategoriUri: %+v\n", kategoriUri)
			fmt.Printf("error: %+v\n", err)
			ctx.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"errors": nil,
			"data":   k,
		})
	})

	r.Run("localhost:3000")
}

// 4. Buat halaman untuk menampilkan data yang sudah anda simpan
// 5. Lalu tampilkan data yang hanya memiliki status " bisa dijual " this should be done on react
// 7. Untuk fitur tambah dan edit gunakan form validasi (inputan nama harus diisi, dan harga harus berupa inputan angka)

// if qId := ctx.Query("id"); qId != "" {
// 	id, err := strconv.Atoi(qId)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{
// 			"errors": err.Error(),
// 		})
// 		return
// 	}
// 	k, err := queries.GetKategoriById(ctx, int32(id))
// 	if err != nil {
// 		ctx.JSON(http.StatusInternalServerError, gin.H{
// 			"errors": err.Error(),
// 		})
// 		return
// 	}

// 	ctx.JSON(http.StatusOK, gin.H{
// 		"errors": nil,
// 		"data":   k,
// 	})
// 	return
// }
