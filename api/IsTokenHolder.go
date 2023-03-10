package api

import (
	"math/big"
	"web3/contracts"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func IsTokenHolder(wallet string, contract string, network string) (bool, error) {
	var client, err = ethclient.Dial(GetRPC(network))
	if err != nil {
		return false, err
	}

	walletAddress := common.HexToAddress(wallet)
	contractAddress := common.HexToAddress(contract)

	nft, err := contracts.NewIERC721ACaller(contractAddress, client)

	tx, err := nft.BalanceOf(nil, walletAddress)
	if err != nil {
		return false, err
	}

	return tx.Cmp(big.NewInt(0)) > 0, nil
}
