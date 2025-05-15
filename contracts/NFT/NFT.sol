// contracts/Market.sol
// SPDX-License-Identifier: MIT OR Apache-2.0
pragma solidity ^0.8.4;

import "@openzeppelin/contracts/utils/Counters.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/token/ERC721/ERC721.sol";

contract NFT is ERC721URIStorage {
    using Counters for Counters.Counter;
    Counters.Counter private _tokenIds;
    address contractAddress;

    event NFTMinted(
        address indexed owner,
        uint256 indexed tokenId,
        string tokenURI
    );

    constructor(address marketplaceAddress) ERC721("LaunchPad", "LCP") {
        contractAddress = marketplaceAddress;
    }

    function createToken(
        address receipt,
        string memory tokenURI
    ) public returns (uint) {
        _tokenIds.increment();
        uint256 newItemId = _tokenIds.current();

        _mint(receipt, newItemId);
        _setTokenURI(newItemId, tokenURI);
        setApprovalForAll(contractAddress, true);

        emit NFTMinted(receipt, newItemId, tokenURI);
        return newItemId;
    }
}
