package main

import (
	"fmt"
	"log"
	"net/http"

	ipfsutil "github.com/RadicalPixels/ipfsutil"
)

func main() {
	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		file, _, err := r.FormFile("file")
		if err != nil {
			log.Print(err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		defer file.Close()

		hash, err := ipfsutil.AddFile(file)
		if err != nil {
			log.Print(err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		fmt.Printf("hash: %s\n", hash)

		if _, err = w.Write([]byte(hash)); err != nil {
			log.Print(err)
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
	})

	port := 8000
	log.Printf("listening on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
