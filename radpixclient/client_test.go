package radpixclient

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/common"
)

func TestClient(t *testing.T) {
	cl := NewClient(&Config{
		HostURL: "https://rinkeby.infura.io",
		Address: "0x0000000000000000000000000000000000000001",
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
