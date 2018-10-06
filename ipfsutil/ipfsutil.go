package ipfsutil

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"time"

	ipfs "github.com/ipfs/go-ipfs-api"
)

// Publish returns IPNS name
func Publish(hash string) (string, error) {
	shell := ipfs.NewLocalShell()
	keyName := "self"
	resp, err := shell.PublishWithDetails(hash, keyName, time.Hour*100, time.Hour*100, false)
	if err != nil {
		return "", err
	}

	return resp.Name, nil
}

// Add ...
func Add(data []byte) (string, error) {
	// my have ipfs daemon running first
	shell := ipfs.NewLocalShell()

	reader := bytes.NewReader(data)
	hash, err := shell.Add(reader)
	if err != nil {
		log.Fatal(err)
	}

	return hash, nil
}

// AddFile ...
func AddFile(file io.Reader) (string, error) {
	// my have ipfs daemon running first
	shell := ipfs.NewLocalShell()

	var buf bytes.Buffer
	if _, err := io.Copy(&buf, file); err != nil {
		return "", err
	}

	fmt.Printf("adding to ipfs: %s\n", string(buf.Bytes()))

	hash, err := shell.Add(file)
	if err != nil {
		log.Fatal(err)
	}

	return hash, nil
}
