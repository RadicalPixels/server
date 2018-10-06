package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	ipfsutil "github.com/RadicalPixels/server/ipfsutil"
)

// Run ...
func Run() {
	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		file, _, err := r.FormFile("file")
		if err != nil {
			handleError(w, err)
		}
		defer file.Close()

		hash, err := ipfsutil.Add(file)
		if err != nil {
			handleError(w, err)
		}

		fmt.Printf("hash: %s\n", hash)
		ipns, err := ipfsutil.Publish(hash)
		if err != nil {
			handleError(w, err)
		}

		resp := struct {
			IPNS string `json:"ipns"`
			Hash string `json:"hash"`
		}{
			IPNS: ipns,
			Hash: hash,
		}

		fmt.Printf("ipns: %s\n", ipns)
		jsonBytes, err := json.Marshal(resp)
		w.Header().Set("Content-Type", "application/json")
		if _, err = w.Write(jsonBytes); err != nil {
			handleError(w, err)
		}
	})

	port := 8000
	log.Printf("listening on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}

func handleError(w http.ResponseWriter, err error) {
	log.Print(err)
	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
}
