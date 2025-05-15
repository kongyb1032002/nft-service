package dtos

type DonationNftDto struct {
	TokenId  string `json:"tokenId"`
	TokenURI string `json:"tokenURI"`
	TxHash   string `json:"txHash"`
}
