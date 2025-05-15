package clients

import (
	"bytes"
	"fmt"
	"io"
	"nft-service/configs"

	shell "github.com/ipfs/go-ipfs-api"
)

type IpfsClient struct {
	apiUrl string
	sh     *shell.Shell
}

func NewIpfsClient(cfg *configs.Config) *IpfsClient {
	return &IpfsClient{
		apiUrl: cfg.IPFSEndpoint,
		sh:     shell.NewShell(cfg.IPFSEndpoint),
	}
}

func (c *IpfsClient) Upload(data []byte) (string, error) {
	reader := bytes.NewReader(data)
	cid, err := c.sh.Add(reader)
	if err != nil {
		return "", err
	}
	return cid, nil
}

func (c *IpfsClient) GetFile(cid string) ([]byte, error) {
	reader, err := c.sh.Cat(cid)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	return io.ReadAll(reader)
}

func (c *IpfsClient) DeleteFile(cid string) error {
	return c.sh.Unpin(cid) // Chỉ unpin, không xóa khỏi network nếu đã public
}

func (c *IpfsClient) GetFileInfo(cid string) (string, error) {
	links, err := c.sh.List(cid)
	if err != nil {
		return "", err
	}

	result := fmt.Sprintf("File: %s\n", cid)
	for _, link := range links {
		result += fmt.Sprintf(" - %s (%d bytes)\n", link.Name, link.Size)
	}
	return result, nil
}
