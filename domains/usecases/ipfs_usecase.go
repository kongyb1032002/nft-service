package usecases

import (
	"encoding/json"
	"log"
	"nft-service/models"
	"os"
)

type IPFSUsecase interface {
	UploadMetadata(metadata models.Metadata) (string, error)
	GetMetadata(tokenURI string) (models.Metadata, error)
}

type ipfsUsecase struct {
	client IPFSClient
}

func NewIPFSUsecase(client IPFSClient) IPFSUsecase {
	return &ipfsUsecase{client: client}
}

func (u *ipfsUsecase) UploadMetadata(metadata models.Metadata) (string, error) {
	// Convert metadata to JSON bytes
	metadataBytes, err := json.Marshal(metadata)
	if err != nil {
		log.Printf("Error marshalling metadata: %v", err)
		return "", err
	}

	// Save metadata to a local JSON file
	fileName := "metadata.json"
	err = os.WriteFile(fileName, metadataBytes, 0644)
	if err != nil {
		log.Printf("Error writing metadata to file: %v", err)
		return "", err
	}

	// Upload the metadata to IPFS
	ipfsHash, err := u.client.Upload(metadataBytes)
	if err != nil {
		log.Printf("Error uploading metadata to IPFS: %v", err)
		return "", err
	}

	// Optionally, remove the local file after upload
	err = os.Remove(fileName)
	if err != nil {
		log.Printf("Error removing local file: %v", err)
		return "", err
	}

	return ipfsHash, nil
}

func (u *ipfsUsecase) GetMetadata(tokenURI string) (models.Metadata, error) {
	metadataBytes, err := u.client.GetFile(tokenURI)
	if err != nil {
		log.Printf("Error getting metadata from IPFS: %v", err)
		return models.Metadata{}, err
	}
	var metadata models.Metadata
	err = json.Unmarshal(metadataBytes, &metadata)
	if err != nil {
		log.Printf("Error unmarshalling metadata: %v", err)
		return models.Metadata{}, err
	}
	return metadata, nil
}
