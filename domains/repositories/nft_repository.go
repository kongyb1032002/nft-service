package repositories

import (
	"nft-service/models"

	"gorm.io/gorm"
)

type NFTRepository interface {
	SaveToken(token *models.Token) error
	GetNFTsByWallet(wallet string) ([]models.Token, error)
}

type nftRepository struct {
	db *gorm.DB
}

func NewNFTRepository(db *gorm.DB) NFTRepository {
	return &nftRepository{db: db}
}

func (r *nftRepository) SaveToken(token *models.Token) error {
	if err := r.db.Create(token).Error; err != nil {
		return err
	}
	return nil
}

func (r *nftRepository) GetNFTsByWallet(wallet string) ([]models.Token, error) {
	var tokens []models.Token
	if err := r.db.Where("owner = ?", wallet).Find(&tokens).Error; err != nil {
		return nil, err
	}
	return tokens, nil
}
