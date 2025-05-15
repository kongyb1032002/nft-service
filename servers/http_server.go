package servers

import (
	"nft-service/configs"
	"nft-service/controllers"
	"nft-service/domains/repositories"
	"nft-service/domains/usecases"
	"nft-service/infrastructure/clients"
	"nft-service/infrastructure/middlewares"
	"nft-service/services"

	"github.com/gin-gonic/gin"
)

type HTTPServer struct {
	r   *gin.Engine
	cfg *configs.Config
}

func NewHTTPServer(cfg *configs.Config) *HTTPServer {
	db := configs.InitDB(cfg)
	r := gin.Default()
	r.Use(middlewares.CORSMiddleware())

	ipfsClient := clients.NewIpfsClient(cfg)
	nftClient := clients.NewNFTClient(cfg)

	nftRepo := repositories.NewNFTRepository(db)

	nftUsecase := usecases.NewNFTUsecase(nftClient)
	ipfsUsecase := usecases.NewIPFSUsecase(ipfsClient)

	nftService := services.NewNFTService(nftUsecase, ipfsUsecase, nftRepo)
	controllers.RegisterNFTRoutes(r, nftService)

	return &HTTPServer{r: r, cfg: cfg}
}

func (s *HTTPServer) Run() {
	s.r.Run(s.cfg.ServerAddress)
}
