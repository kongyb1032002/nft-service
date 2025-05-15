package usecase_errors

type ErrorCode string

const (
	ErrCodeIPFSUpload      ErrorCode = "IPFS_UPLOAD_ERROR"
	ErrCodeMintNFT         ErrorCode = "MINT_NFT_ERROR"
	ErrCodeGetNFTURI       ErrorCode = "GET_NFT_URI_ERROR"
	ErrCodeMarshalMetadata ErrorCode = "MARSHAL_METADATA_ERROR"
	ErrCodeFileWrite       ErrorCode = "FILE_WRITE_ERROR"
	ErrCodeFileRemove      ErrorCode = "FILE_REMOVE_ERROR"
	ErrCodeGetFile         ErrorCode = "GET_FILE_ERROR"
	ErrCodeUnmarshal       ErrorCode = "UNMARSHAL_ERROR"
)
