package main

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/tweetyah/lib"
)

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func main() {
	godotenv.Load()
	// Read the entire file into a byte slice
	bytes, err := ioutil.ReadFile("./image.png")
	if err != nil {
		log.Fatal(err)
	}

	var encoded string

	// Determine the content type of the image file
	mimeType := http.DetectContentType(bytes)

	// Prepend the appropriate URI scheme header depending
	// on the MIME type
	switch mimeType {
	case "image/jpeg":
		encoded += "data:image/jpeg;base64,"
	case "image/png":
		encoded += "data:image/png;base64,"
	}

	// Append the base64 encoded output
	encoded += toBase64(bytes)

	// Print the full base64 representation of the image
	// fmt.Println(encoded)

	token := os.Getenv("TOKEN")
	domain := "fosstodon.org"

	err = lib.UploadMediaToMastodon(domain, token, encoded)
	if err != nil {
		log.Fatal(err)
	}
}
