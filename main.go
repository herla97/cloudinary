package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error getting env, not comming through %v", err)
	}
}

// Signature interface
type Signature interface {
	SignatureGenerator() Response
}

// Response Model
type Response struct {
	Signature string
	Time      string
}

// Upload Model
type Upload struct {
	APISecret string
	Folder    string
	Tags      string
}

// Replace Model
type Replace struct {
	APISecret string
	PublicID  string
}

// generator helper
func generator(part string) string {
	hash := sha1.New()
	io.WriteString(hash, part)

	return hex.EncodeToString(hash.Sum(nil))
}

// SignatureGenerator type Upload
func (u Upload) SignatureGenerator() Response {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	signature := generator(fmt.Sprintf("folder=%s&tags=%s&timestamp=%s%s", u.Folder, u.Tags, timestamp, u.APISecret))
	return Response{
		Signature: signature,
		Time:      timestamp,
	}
}

// SignatureGenerator type Destroy
func (r Replace) SignatureGenerator() Response {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	signature := generator(fmt.Sprintf("public_id=%s&timestamp=%s%s", r.PublicID, timestamp, r.APISecret))
	return Response{
		Signature: signature,
		Time:      timestamp,
	}
}

func main() {
	// var s Signature = Upload{
	// 	APISecret: os.Getenv("APISECRET"),
	// 	Folder:    "loop/artist",
	// 	Tags:      "artist",
	// }

	var s Signature = Replace{
		APISecret: os.Getenv("APISECRET"),
		PublicID:  "loop/artist/gthcg1xcui6iczsw4zxt",
	}

	log.Println(s.SignatureGenerator())
}
