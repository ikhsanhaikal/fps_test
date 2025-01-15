package main

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/joho/godotenv"
	"ikhsanhaikal.com/fastprint-test/pgdb"
)

type FastPrintData struct {
	No          string `json:"no"`
	ID_Produk   string `json:"id_produk"`
	Nama_Produk string `json:"nama_produk"`
	Kategori    string `json:"kategori"`
	Harga       string `json:"harga"`
	Status      string `json:"status"`
}
type SuccessResponse struct {
	Error   int             `json:"error"`
	Version string          `json:"version"`
	Data    []FastPrintData `json:"data"`
}
type Memo struct {
	Exist bool
	Id    int32
}

func initialize(queries *pgdb.Queries) {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	postUrl := os.Getenv("FASTPRINT_TEST_API")
	pass := os.Getenv("FASTPRINT_TEST_API_PASS")

	resp, err := http.Head(postUrl)

	if err != nil {
		panic(err.Error())
	}

	uh := resp.Header.Get("x-credentials-username")
	username := strings.Trim(strings.Split(uh, "(")[0], " ")

	y, m, d := time.Now().Date()
	raw := fmt.Sprintf("%s-%02d-%02d-%02d", pass, d, m, y%100)

	hash := md5.New()
	hash.Write([]byte(raw))
	hashStr := hex.EncodeToString(hash.Sum(nil))

	resp, err = http.PostForm(postUrl, url.Values{
		"username": []string{username},
		"password": []string{hashStr},
	})

	fmt.Printf("username: %s,\npassword: %s\n, raw: %s\n", username, hashStr, raw)

	if err != nil {
		panic(err.Error())
	}

	if resp.StatusCode != 200 {
		msg, _ := io.ReadAll(resp.Body)
		fmt.Printf("%#v\n", string(msg))
	}

	b, _ := io.ReadAll(resp.Body)

	var successResp SuccessResponse

	if err := json.Unmarshal(b, &successResp); err != nil {
		panic(err.Error())
	}

	fmt.Printf("%+v", successResp)

	var inserted map[string]Memo = map[string]Memo{}

	for _, data := range successResp.Data {
		if !inserted[data.Kategori].Exist {
			k, err := queries.CreateKategori(context.Background(), data.Kategori)
			if err != nil {
				panic(err.Error())
			}
			inserted[data.Kategori] = Memo{Exist: true, Id: k.IDKategori}
		}
		if !inserted[data.Status].Exist {
			s, err := queries.CreateStatus(context.Background(), data.Status)
			if err != nil {
				panic(err.Error())
			}
			inserted[data.Status] = Memo{Exist: true, Id: s.IDStatus}
		}

		harga, _ := strconv.ParseInt(data.Harga, 10, 64)

		_, err := queries.CreateProduk(context.Background(), pgdb.CreateProdukParams{
			NamaProduk: data.Nama_Produk,
			Harga: pgtype.Numeric{
				Int:   big.NewInt(harga),
				Exp:   0,
				NaN:   false,
				Valid: true,
			},
			KategoriID: int64(inserted[data.Kategori].Id),
			StatusID:   int64(inserted[data.Status].Id),
		})

		if err != nil {
			panic(err.Error())
		}
	}
}
