package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

func initialize() {
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
}

func main() {
	r := gin.Default()

	initialize()

	r.GET("/welcome", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Greetings",
		})
	})

	// r.Run("localhost:3000")
}
