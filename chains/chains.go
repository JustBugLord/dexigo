package chains

import (
	"fmt"
	"strconv"
	"strings"
)

type Chain int

const (
	Ton       Chain = 607
	Bsc       Chain = 56
	Arbitrum  Chain = 42161
	Cronos    Chain = 25
	IoTex     Chain = 4689
	Tron      Chain = 195
	Aptos     Chain = 637
	Mantle    Chain = 5000
	ApeChain  Chain = 33139
	Zeta      Chain = 7000
	Solana    Chain = 501
	Avalanche Chain = 43114
	Optimism  Chain = 10
	Sonic     Chain = 146
	Mode      Chain = 34443
	Merlin    Chain = 4200
	Conflux   Chain = 1030
	Metis     Chain = 1088
	Polygon   Chain = 137
	ZkSync    Chain = 324
	PolygonZk Chain = 1101
	Sui       Chain = 784
	Manta     Chain = 169
	Ethereum  Chain = 1
	Base      Chain = 8453
	Scroll    Chain = 534352
	Linea     Chain = 59144
	Sei       Chain = 1329
	X         Chain = 196
	Fantom    Chain = 250
)

func (c *Chain) UnmarshalJSON(data []byte) error {
	cleaned := strings.Trim(string(data), `"`)
	num, err := strconv.ParseInt(cleaned, 10, 64)
	if err != nil {
		return fmt.Errorf("cannot parse chain ID: %s", cleaned)
	}

	*c = Chain(num)
	return nil
}
