package dexigo

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/JustBugLord/dexigo/chains"
)

func (okx *Okx) SearchToken(content string) (*SearchTokenResponse, error) {
	return okx.SearchTokenAdvance(content, 1, 1)
}

func (okx *Okx) SearchTokenAdvance(token string, chainType, searchType int) (*SearchTokenResponse, error) {
	response, err := okx.rb.Get(fmt.Sprintf("https://web3.okx.com/priapi/v1/dx/trade/multi/tokens/single/search?chainType=%d&searchType=%d&inputContent=%s", chainType, searchType, token))
	if err != nil {
		return nil, errors.New("fail to get token info: " + err.Error())
	}
	if response.StatusCode != 200 {
		return nil, errors.New("fail to get token info: " + response.Body)
	}
	var result *SearchTokenResponse
	if err := json.Unmarshal([]byte(response.Body), &result); err != nil {
		return nil, errors.New("fail to unmarshal token info: " + err.Error())
	}
	return result, nil
}

func (okx *Okx) TokenInfo(address string, chainId chains.Chain) (*TokenInfoResponse, error) {
	response, err := okx.rb.Get(fmt.Sprintf("https://web3.okx.com/priapi/v1/dx/market/v2/latest/info?tokenContractAddress=%s&chainId=%d", address, chainId))
	if err != nil {
		return nil, errors.New("fail to get token info: " + err.Error())
	}
	if response.StatusCode != 200 {
		return nil, errors.New("fail to get token info: " + response.Body)
	}
	var result *TokenInfoResponse
	if err := json.Unmarshal([]byte(response.Body), &result); err != nil {
		return nil, errors.New("fail to unmarshal token info: " + err.Error())
	}
	return result, nil
}

func (okx *Okx) AllNetworkTokens() (*AllNetworkTokensResponse, error) {
	response, err := okx.rb.Get("https://web3.okx.com/priapi/v1/dx/trade/multi/swap/allNetWorkTokens")
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New("fail to get all network tokens: " + response.Body)
	}
	var result *AllNetworkTokensResponse
	if err := json.Unmarshal([]byte(response.Body), &result); err != nil {
		return nil, errors.New("fail to unmarshal body: " + err.Error())
	}
	return result, nil
}
