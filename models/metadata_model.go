package models

type Metadata struct {
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Data        MetadataData `json:"data"`
}

type MetadataData struct {
	WalletAddress  string  `json:"walletAddress"`
	ProjectID      string  `json:"projectId"`
	ProjectName    string  `json:"projectName"`
	DonationAmount float64 `json:"donationAmount"`
	ProjectOwner   string  `json:"projectOwner"`
	Note           string  `json:"note"`
	Metadata       []byte  `json:"metadata,omitempty"`
}
