package clients

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"nft-service/configs"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type nftClient struct {
	client       *ethclient.Client
	privateKey   *ecdsa.PrivateKey
	publicAddr   common.Address
	contract     *bind.BoundContract
	contractAddr common.Address
	chainID      *big.Int
}

func NewNFTClient(cfg *configs.Config) *nftClient {
	client, err := ethclient.Dial(cfg.JsonRpcURL)
	if err != nil {
		panic(fmt.Errorf("failed to connect to blockchain: %w", err))
	}

	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(cfg.PrivateKey, "0x"))
	if err != nil {
		panic(fmt.Errorf("invalid private key: %w", err))
	}
	publicAddr := crypto.PubkeyToAddress(privateKey.PublicKey)

	abiFile, err := os.ReadFile("contracts/NFT/abi.json")
	if err != nil {
		panic(fmt.Errorf("cannot read abi.json: %w", err))
	}

	var parsedABI abi.ABI
	if err := json.Unmarshal(abiFile, &parsedABI); err != nil {
		panic(fmt.Errorf("invalid ABI format: %w", err))
	}

	contractAddr := common.HexToAddress(cfg.ContractAddr)
	contract := bind.NewBoundContract(contractAddr, parsedABI, client, client, client)

	return &nftClient{
		client:       client,
		privateKey:   privateKey,
		publicAddr:   publicAddr,
		contract:     contract,
		contractAddr: contractAddr,
		chainID:      big.NewInt(int64(cfg.ChainID)),
	}
}

func (c *nftClient) MintNFT(to string, tokenURI string) (int64, string, error) {
	auth, err := bind.NewKeyedTransactorWithChainID(c.privateKey, c.chainID)
	if err != nil {
		log.Printf("failed to create transactor: %v", err)
		return 0, "", fmt.Errorf("transactor failed: %w", err)
	}
	auth.From = c.publicAddr
	auth.Context = context.Background()
	auth.GasLimit = uint64(3000000)

	toAddress := common.HexToAddress(to)
	tx, err := c.contract.Transact(auth, "createToken", toAddress, tokenURI)
	if err != nil {
		log.Printf("failed to call createToken: %v", err)
		return 0, "", fmt.Errorf("call createToken failed: %w", err)
	}

	receipt, err := bind.WaitMined(context.Background(), c.client, tx)
	if err != nil {
		return 0, tx.Hash().Hex(), fmt.Errorf("tx mining error: %w", err)
	}

	// Event signature: NFTMinted(address,uint256,string)
	eventSig := []byte("NFTMinted(address,uint256,string)")
	eventID := crypto.Keccak256Hash(eventSig) // keccak256

	for _, vLog := range receipt.Logs {
		if len(vLog.Topics) == 3 && vLog.Topics[0] == eventID {
			// Decode topics
			tokenIdBig := new(big.Int).SetBytes(vLog.Topics[2].Bytes())
			tokenId := tokenIdBig.Int64()

			return tokenId, tx.Hash().Hex(), nil
		}
	}

	return 0, tx.Hash().Hex(), nil // tokenId mock
}

func (c *nftClient) GetNftURI(tokenId int64) (string, error) {
	tokenIdBig := big.NewInt(tokenId)
	var tokenURI string

	var result []any
	err := c.contract.Call(nil, &result, "tokenURI", tokenIdBig)
	if err == nil && len(result) > 0 {
		tokenURI = result[0].(string)
	}
	if err != nil {
		log.Printf("failed to call tokenURI: %v", err)
		return "", fmt.Errorf("failed to call tokenURI: %w", err)
	}

	return tokenURI, nil
}
