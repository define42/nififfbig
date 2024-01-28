package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
	"bytes"
	"github.com/define42/goflowfilev3"
)

func bigfile100mb(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	start := time.Now()

	fileContent := "Hello, NiFi!asdasdasdahsdkjasdhasdjkasdh asdhaskdj asdhaksjdhasdka sdasdas dhaksjdashfa skjasdh asdjaksd askjdkahs dakjs dh"
	packager := goflowfilev3.NewFlowFilePackagerV3()
	attributes := map[string]string{
		"Author":   "John Doe",
		"filename": "image.jpg",
		"Type":     "Example",
	}
	const size10MB = 1000 * 1024 * 1024

	for i := 0; i <= 1000000; i++ {
		err := packager.PackageFlowFile(bytes.NewReader([]byte(fileContent)), w, attributes, int64(len(fileContent)))
		if err != nil {
			fmt.Println("Error creating flow file:", err)
			return
		}
	}
	elapsed := time.Since(start)
	fmt.Println(" elapsed:", elapsed, " size:", int64(len(fileContent))*10000000, "IP:", r.RemoteAddr)

}

func main() {
	server := &http.Server{
		Addr: ":8080",
	}

	server.SetKeepAlivesEnabled(false)

	http.HandleFunc("/", bigfile100mb)
	log.Fatal(server.ListenAndServe())
}
