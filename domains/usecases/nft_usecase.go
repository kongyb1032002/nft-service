package usecases

import (
	"fmt"
	"log"
)

type IPFSClient interface {
	Upload(metadata []byte) (string, error)
	GetFile(tokenURI string) ([]byte, error)
}

type NFTClient interface {
	MintNFT(to string, tokenURI string) (int64, string, error)
	GetNftURI(tokenId int64) (string, error)
}

type NFTUsecase interface {
	MintNFT(to string, tokenURI string) (int64, string, error)
	GetNftURI(tokenId int64) (string, error)
}

type nftUsecase struct {
	client NFTClient
}

func NewNFTUsecase(client NFTClient) NFTUsecase {
	return &nftUsecase{client: client}
}

func (u *nftUsecase) MintNFT(to string, tokenURI string) (int64, string, error) {
	tokenId, txHash, err := u.client.MintNFT(to, tokenURI)
	if err != nil {
		log.Printf("Error minting NFT: %v", err)
		return 0, "", err
	}
	if tokenId == 0 {
		log.Printf("Mint failed: tokenId is 0, txHash: %s", txHash)
		return 0, txHash, fmt.Errorf("mint failed: tokenId is 0")
	}
	return tokenId, txHash, nil
}

func (u *nftUsecase) GetNftURI(tokenId int64) (string, error) {
	tokenURI, err := u.client.GetNftURI(tokenId)
	if err != nil {
		log.Printf("Error getting NFT URI: %v", err)
		return "", err
	}
	return tokenURI, nil
}
