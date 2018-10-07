package server

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	ipfsutil "github.com/RadicalPixels/server/ipfsutil"
	radpixclient "github.com/RadicalPixels/server/radpixclient"
	util "github.com/RadicalPixels/server/util"
	"github.com/ethereum/go-ethereum/common"
	gocache "github.com/patrickmn/go-cache"
)

// Pixel ...
type Pixel struct {
	Index    int64    `json:"index"`
	ID       string   `json:"id"`
	X        int64    `json:"x"`
	Y        int64    `json:"y"`
	Owner    string   `json:"owner"`
	Price    float64  `json:"price"`
	Content  string   `json:"content"`
	Colors   []string `json:"colors"`
	Sellable bool     `json:"sellable"`
}

// Server ...
type Server struct {
	cache  *gocache.Cache
	client *radpixclient.Client
}

// Config ...
type Config struct {
	NodeURL         string
	ContractAddress string
}

// NewServer ...
func NewServer(cfg *Config) *Server {
	cache := gocache.New(5*time.Minute, 5*time.Minute)

	client := radpixclient.NewClient(&radpixclient.Config{
		HostURL: cfg.NodeURL,
		Address: cfg.ContractAddress,
	})

	return &Server{
		cache:  cache,
		client: client,
	}
}

// Start ...
func (s *Server) Start() {
	http.HandleFunc("/upload", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("received request")
		file, _, err := r.FormFile("file")
		if err != nil {
			handleError(w, err)
			return
		}
		defer file.Close()

		hash, err := ipfsutil.Add(file)
		if err != nil {
			handleError(w, err)
			return
		}

		fmt.Printf("hash: %s\n", hash)
		ipns, err := ipfsutil.Publish(hash)
		if err != nil {
			handleError(w, err)
			return
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
			return
		}
	})

	http.HandleFunc("/grid", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("received request")

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		pixels, err := s.getPixels()
		if err != nil {
			handleError(w, err)
			return
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
			return
		}
	})

	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("index")
	})

	port := 8000
	log.Printf("listening on port %d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", port), nil))
}

// getPixels ...
func (s *Server) getPixels() ([]Pixel, error) {
	grid, found := s.cache.Get("grid")
	if found {
		fmt.Println("cached")
		pixels := grid.([]Pixel)
		return pixels, nil
	}

	x, y, err := s.client.GridSize()
	if err != nil {
		return nil, err
	}

	events, err := s.client.Query(&radpixclient.QueryConfig{
		LogTopics: [][]common.Hash{
			[]common.Hash{radpixclient.LogBuyPixelTopic},
		},
	})
	if err != nil {
		return nil, err
	}

	pixelsMap := make(map[string]*Pixel)
	for _, event := range events {
		p, ok := event.(radpixclient.LogBuyPixel)
		if !ok {
			return nil, errors.New("could not typecast")
		}

		x := p.X.Int64()
		y := p.Y.Int64()

		pixel := &Pixel{}
		pixel.X = x
		pixel.Y = y
		pixel.ID = hex.EncodeToString(p.ID[:])
		pixel.Owner = "0x" + hex.EncodeToString(p.Buyer[:])
		price, _ := util.ToDecimal(p.Price, 18).Float64()
		pixel.Price = price

		empty := [32]byte{}
		if !bytes.Equal(empty[:], p.ContentData[:]) {
			pixel.Content = hex.EncodeToString(p.ContentData[:])
		}
		if pixel.Content != "" {
			pixel.Colors = util.ParseContinuousColorHexString(pixel.Content)
		}
		pixel.Sellable = true
		pixelsMap[fmt.Sprintf("%d,%d", x, y)] = pixel
	}

	var pixels []Pixel
	index := 0
	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			pix, ok := pixelsMap[fmt.Sprintf("%d,%d", i, j)]
			if ok {
				pixels = append(pixels, Pixel{
					Index:    int64(index),
					X:        int64(i),
					Y:        int64(j),
					ID:       pix.ID,
					Owner:    pix.Owner,
					Price:    pix.Price,
					Content:  pix.Content,
					Colors:   pix.Colors,
					Sellable: pix.Sellable,
				})
			} else {
				pixels = append(pixels, Pixel{
					Index: int64(index),
					X:     int64(i),
					Y:     int64(j),
				})
			}
			index++
		}
	}

	s.cache.Set("grid", pixels, 15*time.Second)
	return pixels, nil
}

func handleError(w http.ResponseWriter, err error) {
	log.Print(err)
	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
}
