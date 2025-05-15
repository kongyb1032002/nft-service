package controllers

import (
	"net/http"
	"nft-service/dtos"
	"nft-service/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RegisterNFTRoutes(r *gin.Engine, service services.NFTService) {
	nft := r.Group("/api/nft")
	{
		nft.POST("", func(c *gin.Context) {
			handleCreateNFT(c, service)
		})
		nft.GET("/wallet/:wallet", func(c *gin.Context) {
			handleGetNftByWallet(c, service)
		})
		nft.GET("/uri/:tokenURI", func(c *gin.Context) {
			handleGetNFTMetadata(c, service)
		})
		nft.GET("/:tokenId", func(c *gin.Context) {
			handleGetNft(c, service)
		})
	}
}

func handleCreateNFT(c *gin.Context, service services.NFTService) {
	var req dtos.CreateNFTRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "BAD_REQUEST", "error": err.Error()})
		return
	}
	res, err := service.IssueDonationNFT(req)
	if err != nil {
		code, msg := parseServiceError(err)
		c.JSON(http.StatusInternalServerError, gin.H{"code": code, "error": msg})
		return
	}
	c.JSON(http.StatusCreated, res)
}

func handleGetNftByWallet(c *gin.Context, service services.NFTService) {
	wallet := c.Param("wallet")
	res, err := service.GetNFTsByWallet(wallet)
	if err != nil {
		code, msg := parseServiceError(err)
		c.JSON(http.StatusInternalServerError, gin.H{"code": code, "error": msg})
		return
	}
	c.JSON(http.StatusOK, res)
}

func handleGetNFTMetadata(ctx *gin.Context, service services.NFTService) {
	tokenURI := ctx.Param("tokenURI")
	if tokenURI == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": "BAD_REQUEST", "error": "tokenURI is required"})
		return
	}

	metadata, err := service.GetNFTMetadata(tokenURI)
	if err != nil {
		code, msg := parseServiceError(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": code, "error": msg})
		return
	}

	ctx.JSON(http.StatusOK, metadata)
}

func handleGetNft(c *gin.Context, service services.NFTService) {
	tokenId := c.Param("tokenId")
	if tokenId == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": "BAD_REQUEST", "error": "tokenId is required"})
		return
	}
	tokenIdInt, err := strconv.ParseInt(tokenId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": "BAD_REQUEST", "error": "invalid tokenId"})
		return
	}
	res, err := service.GetNft(tokenIdInt)
	if err != nil {
		code, msg := parseServiceError(err)
		c.JSON(http.StatusInternalServerError, gin.H{"code": code, "error": msg})
		return
	}
	c.JSON(http.StatusOK, res)
}

// Helper: parse error để lấy code và message trả về client
func parseServiceError(err error) (string, string) {
	type coder interface{ Code() string }
	if ue, ok := err.(coder); ok {
		return ue.Code(), err.Error()
	}
	return "INTERNAL_ERROR", err.Error()
}
