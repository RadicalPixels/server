package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"

	ipfsutil "github.com/RadicalPixels/server/ipfsutil"
	radpixclient "github.com/RadicalPixels/server/radpixclient"
	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/common"
)

// Pixel ...
type Pixel struct {
	ID     string   `json:"ID"`
	X      int64    `json:"x"`
	Y      int64    `json:"y"`
	Owner  string   `json:"owner"`
	Price  float64  `json:"price"`
	Colors []string `json:"colors"`
}

// Run ...
func Run() {
	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("received request")
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

	http.HandleFunc("/grid", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("received request")
		client := radpixclient.NewClient(&radpixclient.Config{
			HostURL: "https://rinkeby.infura.io",
			Address: "",
		})

		events, err := client.Query(&radpixclient.QueryConfig{
			LogTopics: [][]common.Hash{
				[]common.Hash{radpixclient.LogBuyPixelTopic},
			},
		})

		if err != nil {
			handleError(w, err)
		}

		spew.Dump(events)

		var pixels []*Pixel
		for _, event := range events {
			var colors []string
			p, ok := event.(radpixclient.LogBuyPixel)
			if !ok {
				log.Fatal("no ok")
			}
			pixels = append(pixels, &Pixel{
				ID:     string(p.ID[:]),
				X:      p.X.Int64(),
				Y:      p.Y.Int64(),
				Owner:  string(p.Buyer[:]),
				Price:  float64(10),
				Colors: colors,
			})
		}

		jsonBytes, err := json.Marshal(pixels)
		w.Header().Set("Content-Type", "application/json")
		if _, err = w.Write(jsonBytes); err != nil {
			handleError(w, err)
		}
	})

	http.HandleFunc("/content", func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var colors []string
		if err := decoder.Decode(&colors); err != nil {
			handleError(w, err)
			return
		}

		if len(colors) == 0 {
			err := errors.New("expected more than 0")
			handleError(w, err)
			return
		}

		contentBytes, err := json.Marshal(colors)
		hash, err := ipfsutil.Add(contentBytes)
		if err != nil {
			handleError(w, err)
			return
		}

		resp := struct {
			//IPNS string `json:"ipns"`
			Hash string `json:"hash"`
		}{
			//IPNS: "",
			Hash: fmt.Sprintf("/ipfs/%s", hash),
		}

		jsonBytes, err := json.Marshal(resp)
		w.Header().Set("Content-Type", "application/json")
		if _, err = w.Write(jsonBytes); err != nil {
			handleError(w, err)
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("index")
	})

	port := 8000
	log.Printf("listening on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), nil))
}

func handleError(w http.ResponseWriter, err error) {
	log.Print(err)
	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
}
