package radpixclient

import (
	"context"
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/common"
)

func TestClient(t *testing.T) {
	cl := NewClient(&Config{
		HostURL: "https://kovan.infura.io",
		Address: "0x6d382af479cc7d5f3337e7224261f6e289fddeb1",
	})

	header, err := cl.client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		t.Error(err)
	}

	t.Log(header.Number)
}

func TestQuery(t *testing.T) {
	cl := NewClient(&Config{
		HostURL: "https://kovan.infura.io",
		Address: "0xa74e7fea1db19f0f41d054854f4d950f1c6ff513",
	})

	events, err := cl.Query(&QueryConfig{
		LogTopics: [][]common.Hash{
			[]common.Hash{LogBuyPixelTopic},
		},
	})

	if err != nil {
		t.Error(err)
	}

	spew.Dump(events)
}

func TestGridSize(t *testing.T) {
	cl := NewClient(&Config{
		HostURL: "https://kovan.infura.io",
		Address: "0x6d382af479cc7d5f3337e7224261f6e289fddeb1",
	})

	x, y, err := cl.GridSize()
	if err != nil {
		t.Error(err)
	}

	if x == 0 {
		t.Error("expected greater than 0")
	}
	if y == 0 {
		t.Error("expected greater than 0")
	}

	t.Log(x, y)
}
