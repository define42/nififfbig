package main

import (
	"bytes"
	"fmt"
	"github.com/define42/goflowfilev3"
	"io"
	"log"
	"net/http"
	"time"
)

// Define a type that will implement io.Reader interface
type DataReader struct {
	size int
	read int
}

// Implement the Read method for DataReader
func (dr *DataReader) Read(p []byte) (int, error) {
	// Check if the read size has reached the total size
	if dr.read >= dr.size {
		return 0, io.EOF // End of file (or data stream in this case)
	}

	// Calculate the remaining data size
	remaining := dr.size - dr.read
	if remaining < len(p) {
		p = p[:remaining]
	}

	// Fill the byte slice with some data, e.g., 'A'
	for i := range p {
		p[i] = 'A'
	}

	// Update the read count
	dr.read += len(p)

	// Return the number of bytes read
	return len(p), nil
}

func bigfile10gb(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	const filesize = 10*1024 * 1024 * 1024
	reader := &DataReader{size: filesize}
	_, err := io.Copy(w, reader)
	if err != nil {
		log.Printf("error copying data to response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
func bigfile100gb(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	const filesize = 100*1024 * 1024 * 1024
	reader := &DataReader{size: filesize}
	_, err := io.Copy(w, reader)
	if err != nil {
		log.Printf("error copying data to response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
func bigfile300gb(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	const filesize = 300*1024 * 1024 * 1024
	reader := &DataReader{size: filesize}
	_, err := io.Copy(w, reader)
	if err != nil {
		log.Printf("error copying data to response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
func bigfile500gb(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	const filesize = 500*1024 * 1024 * 1024
	reader := &DataReader{size: filesize}
	_, err := io.Copy(w, reader)
	if err != nil {
		log.Printf("error copying data to response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}


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
	http.HandleFunc("/10gb", bigfile10gb)
	http.HandleFunc("/100gb", bigfile100gb)
	http.HandleFunc("/300gb", bigfile300gb)
	http.HandleFunc("/500gb", bigfile500gb)
	log.Fatal(server.ListenAndServe())
}
