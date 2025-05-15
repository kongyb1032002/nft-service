package dtos

type CreateNFTRequest struct {
	WalletAddress  string  `json:"walletAddress" binding:"required"`
	ProjectID      string  `json:"projectId" binding:"required"`
	DonationAmount float64 `json:"donationAmount" binding:"required"`
	Name           string  `json:"name" binding:"required"`
	Description    string  `json:"description" binding:"required"`
	ProjectName    string  `json:"projectName" binding:"required"`
	ProjectOwner   string  `json:"projectOwner" binding:"required"`
	Note           string  `json:"note"`
}
