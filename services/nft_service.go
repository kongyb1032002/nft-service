package services

import (
	"fmt"
	"nft-service/domains/repositories"
	"nft-service/domains/usecases"
	"nft-service/dtos"
	"nft-service/models"
)

type NFTService interface {
	IssueDonationNFT(req dtos.CreateNFTRequest) (*dtos.DonationNftDto, error)
	GetNFTsByWallet(wallet string) ([]models.Token, error)
	GetNFTMetadata(tokenURI string) (*models.Metadata, error)
	GetNft(tokenId int64) (*dtos.GetNftReponse, error)
}

type nftService struct {
	ipfsUsecase usecases.IPFSUsecase
	nftUsecase  usecases.NFTUsecase
	nftRepo     repositories.NFTRepository // Thêm repository

}

func NewNFTService(
	nftUsecase usecases.NFTUsecase,
	ipfsUsecase usecases.IPFSUsecase,
	nftRepo repositories.NFTRepository, // Nhận repository từ bên ngoài
) NFTService {
	return &nftService{
		ipfsUsecase: ipfsUsecase,
		nftUsecase:  nftUsecase,
		nftRepo:     nftRepo,
	}
}

func (s *nftService) IssueDonationNFT(req dtos.CreateNFTRequest) (*dtos.DonationNftDto, error) {
	metadata := models.Metadata{
		Name:        req.Name,
		Description: req.Description,
		Data: models.MetadataData{
			WalletAddress:  req.WalletAddress,
			ProjectID:      req.ProjectID,
			ProjectName:    req.ProjectName,
			DonationAmount: req.DonationAmount,
			ProjectOwner:   req.ProjectOwner,
			Note:           req.Note,
		},
	}

	// Step 1: Upload metadata to IPFS
	tokenURI, err := s.ipfsUsecase.UploadMetadata(metadata)
	if err != nil {
		return nil, fmt.Errorf("upload to IPFS failed: %w", err)
	}

	// Step 2: Mint NFT on-chain
	tokenId, txHash, err := s.nftUsecase.MintNFT(req.WalletAddress, tokenURI)
	if err != nil {
		return nil, fmt.Errorf("mint NFT failed: %w", err)
	}

	// Step 3: Lưu thông tin vào cơ sở dữ liệu qua repository
	token := &models.Token{
		Owner:    req.WalletAddress,
		TokenID:  tokenId,
		TokenURI: tokenURI,
	}
	if err := s.nftRepo.SaveToken(token); err != nil {
		return nil, fmt.Errorf("failed to save token to database: %w", err)
	}

	return &dtos.DonationNftDto{
		TokenId:  fmt.Sprintf("%d", tokenId), // Convert tokenId to string
		TokenURI: tokenURI,
		TxHash:   txHash,
	}, nil // Add the missing error return value
}

func (s *nftService) GetNFTsByWallet(wallet string) ([]models.Token, error) {
	nfts, err := s.nftRepo.GetNFTsByWallet(wallet)
	if err != nil {
		return nil, err
	}
	if len(nfts) == 0 {
		return nil, fmt.Errorf("no NFTs found for wallet: %s", wallet)
	}
	return nfts, nil
}

func (s *nftService) GetNFTMetadata(tokenURI string) (*models.Metadata, error) {
	// Lấy metadata từ IPFS thông qua IPFS client
	metadata, err := s.ipfsUsecase.GetMetadata(tokenURI)
	if err != nil {
		return nil, err
	}

	// Parse JSON metadata
	return &metadata, nil
}

func (s *nftService) GetNft(tokenId int64) (*dtos.GetNftReponse, error) {
	uri, err := s.nftUsecase.GetNftURI(tokenId)
	if err != nil {
		return nil, err
	}

	metadata, err := s.ipfsUsecase.GetMetadata(uri)

	var response = dtos.GetNftReponse{
		TokenId:   fmt.Sprintf("%d", tokenId),
		Owner:     metadata.Data.ProjectOwner,
		ProjectId: metadata.Data.ProjectID,
		TokenURI:  uri,
		Metadata:  metadata.Data,
	}
	return &response, nil
}
