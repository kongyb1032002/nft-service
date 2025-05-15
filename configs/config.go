package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	JsonRpcURL    string `mapstructure:"JSON_RPC_URL"`
	IPFSEndpoint  string `mapstructure:"IPFS_ENDPOINT"`
	ChainID       int    `mapstructure:"CHAIN_ID"`
	PrivateKey    string `mapstructure:"PRIVATE_KEY"`
	ContractAddr  string `mapstructure:"CONTRACT_ADDRESS"`
	DbHost        string `mapstructure:"DB_HOST"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile(".env") // CHỈ ĐỊNH TRỰC TIẾP TÊN FILE
	viper.AutomaticEnv()        // Ưu tiên biến môi trường nếu có

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
