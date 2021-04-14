// Package p contains an HTTP Cloud Function.
package p

import (
	"context"
	"io"
	"log"
	"net/http"
	"os"

	storage "cloud.google.com/go/storage"
)

var BUCKET string
var GCS *storage.Client

func init() {
	// set the bucket name from environment variable
	BUCKET = os.Getenv("BUCKET_NAME")

	// initialize the client
	c, err := storage.NewClient(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	GCS = c
}

func ProxyGCS(output http.ResponseWriter, input *http.Request) {
	ctx := context.Background()

	// route HTTP methods to appropriate handlers
	switch input.Method {
	case http.MethodGet:
		GET(ctx, output, input)
	default:
		http.Error(output, "405 - Method Not Allowed", http.StatusMethodNotAllowed)
	}
	return
}

func GET(ctx context.Context, output http.ResponseWriter, input *http.Request) {
	// Do the request to get response content stream
	responseContent, err := GCS.Bucket(BUCKET).Object(input.URL.String()[1:]).NewReader(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer responseContent.Close()

	// Copy with a 16MB buffer (aligns well with optimal chunk size for GCS)
	_, err = io.CopyBuffer(output, responseContent, make([]byte, 16*1024^2))
	if err != nil {
		log.Fatal(err)
	}
	return
}
