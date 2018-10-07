package radpixclient

import (
	"context"
	"log"
	"math/big"
	"strings"

	contract "github.com/RadicalPixels/server/contract"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Client ...
type Client struct {
	client     *ethclient.Client
	startBlock *big.Int
	address    common.Address
}

// Config ...
type Config struct {
	HostURL string
	Address string
}

// NewClient ...
func NewClient(cfg *Config) *Client {
	client, err := ethclient.Dial(cfg.HostURL)
	if err != nil {
		log.Fatal(err)
	}

	return &Client{
		client:     client,
		address:    common.HexToAddress(cfg.Address),
		startBlock: big.NewInt(8989021),
		//3116194
	}
}

// LogBuyPixel ...
type LogBuyPixel struct {
	ID          [32]byte
	Seller      common.Address
	Buyer       common.Address
	X           *big.Int
	Y           *big.Int
	Price       *big.Int
	ContentData [32]byte
}

// LogSetPixelPrice ...
type LogSetPixelPrice struct {
	ID     []byte
	seller common.Address
	X      *big.Int
	Y      *big.Int
	Price  *big.Int
}

// LogAddFunds ...
type LogAddFunds struct {
	Owner      common.Address
	AddedFunds *big.Int
}

// BeginDutchAuction ...
type BeginDutchAuction struct {
	PixelID   [32]byte
	TokenID   *big.Int
	AuctionID [32]byte
	Initiator common.Address
	X         *big.Int
	Y         *big.Int
	StartTime *big.Int
	EndTime   *big.Int
}

// UpdateAuctionBid ...
type UpdateAuctionBid struct {
	pixelID   [32]byte
	TokenID   *big.Int
	AuctionID [32]byte
	Bidder    common.Address
	AmountBet *big.Int
	timeBet   *big.Int
}

// EndDutchAuction ...
type EndDutchAuction struct {
	pixelID [32]byte
	TokenID *big.Int
	Claimer common.Address
	X       *big.Int
	Y       *big.Int
}

// QueryConfig ....
type QueryConfig struct {
	LogTopics [][]common.Hash
}

// LogBuyPixelTopic ...
var LogBuyPixelTopic common.Hash

// Query ...
func (s *Client) Query(cfg *QueryConfig) ([]interface{}, error) {
	header, err := s.client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		if err.Error() != "missing required field 'mixHash' for Header" {
			return nil, err
		}
	}

	query := ethereum.FilterQuery{
		FromBlock: s.startBlock,
		ToBlock:   header.Number,
		//ToBlock: new(big.Int).Add(s.startBlock, big.NewInt(1000000)),
		Addresses: []common.Address{
			s.address,
		},
		Topics: cfg.LogTopics,
	}

	logs, err := s.client.FilterLogs(context.Background(), query)
	if err != nil {
		return nil, err
	}

	contractAbi, err := abi.JSON(strings.NewReader(string(contract.ContractABI)))
	if err != nil {
		return nil, err
	}

	var eventLogs []interface{}
	for _, vLog := range logs {
		switch vLog.Topics[0].Hex() {
		case LogBuyPixelTopic.Hex():
			var event LogBuyPixel
			err := contractAbi.Unpack(&event, "BuyPixel", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}

			event.ID = vLog.Topics[1]
			event.Seller = common.HexToAddress(vLog.Topics[2].Hex())
			event.Buyer = common.HexToAddress(vLog.Topics[3].Hex())

			eventLogs = append(eventLogs, event)
		}
	}

	return eventLogs, nil
}

// GridSize ...
func (s *Client) GridSize() (int, int, error) {
	caller, err := contract.NewContractCaller(s.address, s.client)
	if err != nil {
		return 0, 0, err
	}
	xMax, err := caller.XMax(&bind.CallOpts{})
	if err != nil {
		return 0, 0, err
	}
	yMax, err := caller.XMax(&bind.CallOpts{})
	if err != nil {
		return 0, 0, err
	}

	_ = yMax
	_ = xMax
	//return int(xMax.Uint64()), int(yMax.Uint64()), nil
	return 25, 25, nil
}

// LogTopicHash ...
func LogTopicHash(fnsig string) common.Hash {
	eventSignature := []byte(fnsig)
	hash := crypto.Keccak256Hash(eventSignature)
	return hash
}

func init() {
	LogBuyPixelTopic = LogTopicHash("BuyPixel(bytes32,address,address,uint256,uint256,uint256,bytes32)")
}
