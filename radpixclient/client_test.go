package radpixclient

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/common"
)

func TestClient(t *testing.T) {
	cl := NewClient(&Config{
		//HostURL: "http://localhost:8545",
		HostURL: "https://rinkeby.infura.io",
		Address: "0x6d382af479cc7d5f3337e7224261f6e289fddeb1",
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
