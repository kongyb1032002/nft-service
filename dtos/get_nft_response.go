package dtos

import "nft-service/models"

type GetNftReponse struct {
	TokenId   string              `json:"tokenId"`
	Owner     string              `json:"owner"`
	ProjectId string              `json:"projectId"`
	TokenURI  string              `json:"tokenURI"`
	Metadata  models.MetadataData `json:"metadata"`
}
