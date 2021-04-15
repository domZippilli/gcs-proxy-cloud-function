// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

// init is run once, on a function "cold start"
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

// ProxyGCS is the entry point for the cloud function, providing a proxy that
// permits HTTP protocol usage of a GCS bucket's contents.
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

// GET handles GET requests.
func GET(ctx context.Context, output http.ResponseWriter, input *http.Request) {
	// Do the request to get response content stream
	objectName := convertURLtoObject(input.URL.String())
	responseContent, err := GCS.Bucket(BUCKET).Object(objectName).NewReader(ctx)
	if err != nil {
		if err == storage.ErrObjectNotExist {
			http.Error(output, "404 - Not Found", http.StatusNotFound)
			return
		} else {
			log.Fatal(err)
		}
	}
	defer responseContent.Close()

	// Copy with a 16MB buffer (aligns well with optimal chunk size for GCS)
	_, err = io.CopyBuffer(output, responseContent, make([]byte, 16*1024^2))
	if err != nil {
		log.Fatal(err)
	}
	return
}

// convertURLtoObject converts a URL to an appropriate object path. This also
// includes redirecting root requests "/" to index.html.
func convertURLtoObject(url string) (object string) {
	switch url {
	case "/":
		return "index.html"
	default:
		return url[1:]
	}
}
