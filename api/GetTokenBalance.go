package api

import (
	"fmt"
	"math/big"
	"web3/contracts"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetTokenBalance(wallet string, contract string, network string) (big.Int, error) {
	var client, err = ethclient.Dial(GetRPC(network))
	if err != nil {
		return *big.NewInt(-1), err
	}

	walletAddress := common.HexToAddress(wallet)
	contractAddress := common.HexToAddress(contract)

	nft, err := contracts.NewIERC721ACaller(contractAddress, client)

	tx, err := nft.BalanceOf(nil, walletAddress)
	if err != nil {
		fmt.Println("error with tx", err)
		return *big.NewInt(-2), err
	}

	return *tx, nil
}
