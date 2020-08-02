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
	Signature, Time string
}

// Upload Model
type Upload struct {
	Folder, Tags string
}

// Replace Model
type Replace struct {
	PublicID string
}

// generator helper
func generator(part string) Response {
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	hash := sha1.New()
	part = part + fmt.Sprintf("&timestamp=%s%s", timestamp, os.Getenv("APISECRET"))
	io.WriteString(hash, part)
	return Response{
		Signature: hex.EncodeToString(hash.Sum(nil)),
		Time:      strconv.FormatInt(time.Now().Unix(), 10),
	}
}

// SignatureGenerator type Upload
func (u Upload) SignatureGenerator() Response {
	return generator(fmt.Sprintf("folder=%s&tags=%s", u.Folder, u.Tags))
}

// SignatureGenerator type Destroy
func (r Replace) SignatureGenerator() Response {
	return generator(fmt.Sprintf("public_id=%s", r.PublicID))
}

// Print function
func Print(s Signature) {
	log.Println(s.SignatureGenerator())
}

func main() {
	u := Upload{
		Folder: "loop/artist",
		Tags:   "artist",
	}
	Print(u)

	r := Replace{
		PublicID: "loop/artist/sihq0wba7ngacevkjhbs",
	}
	Print(r)
}
