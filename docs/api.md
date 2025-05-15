---

**Base URL:** `/api/nft`

---

### 1. Tạo NFT chứng nhận donation

**Method:** `POST`
**Endpoint:** `/api/nft`

**Request body (JSON):**

```json
{
   "walletAddress": "0x1234abcd...",
   "projectId": "abc123",
   "donationAmount": 100000,
   "name": "Launchpad Donor Badge",
   "description": "You have donated to a project on our Launchpad",
   "projectName": "Green Energy DAO",
   "projectOwner": "0x5678def...",
   "note": "Thank you for supporting sustainable energy!"
}
```

**Metadata trên IPFS:**

```json
{
   "name": "Launchpad Donor Badge",
   "description": "You have donated to a project on our Launchpad",
   "data": {
      "walletAddress": "0x1234abcd...",
      "projectId": "abc123",
      "projectName": "Green Energy DAO",
      "donationAmount": 100000,
      "projectOwner": "0x5678def...",
      "note": "Thank you for supporting sustainable energy!"
   }
}
```

**Success Response:**

-  **Code:** `201 Created`
-  **Body:**

```json
{
   "tokenId": 12,
   "tokenURI": "ipfs://Qm...",
   "transactionHash": "0x456..."
}
```

**Error Responses:**

-  `400 Bad Request` – Thiếu trường bắt buộc hoặc định dạng không hợp lệ

```json
{
   "error": "Missing required field: walletAddress"
}
```

-  `500 Internal Server Error` – Lỗi nội bộ khi upload IPFS hoặc tương tác blockchain

```json
{
   "error": "Failed to mint NFT: contract execution failed"
}
```

---

### 2. Lấy danh sách toàn bộ NFT

**Method:** `GET`
**Endpoint:** `/api/nft`

**Success Response:**

-  **Code:** `200 OK`
-  **Body:**

```json
[
   {
      "tokenId": 1,
      "owner": "0x1234abcd...",
      "projectId": "abc123",
      "tokenURI": "ipfs://Qm...",
      "name": "Launchpad Donor Badge"
   }
]
```

**Error:**

-  `500 Internal Server Error`

```json
{
   "error": "Failed to fetch NFT list"
}
```

---

### 3. Lấy danh sách NFT theo địa chỉ ví

**Method:** `GET`
**Endpoint:** `/api/nft/by-wallet/:walletAddress`

**Success Response:**

-  **Code:** `200 OK`
-  **Body:**

```json
[
   {
      "tokenId": 2,
      "projectId": "abc123",
      "tokenURI": "ipfs://Qm...",
      "name": "Launchpad Donor Badge"
   }
]
```

**Errors:**

-  `400 Bad Request` – Thiếu hoặc sai định dạng `walletAddress`

```json
{
   "error": "Invalid wallet address"
}
```

-  `500 Internal Server Error`

```json
{
   "error": "Failed to fetch NFT by wallet"
}
```

---

### 4. Lấy chi tiết một NFT

**Method:** `GET`
**Endpoint:** `/api/nft/:tokenId`

**Success Response:**

-  **Code:** `200 OK`
-  **Body:**

```json
{
   "tokenId": 2,
   "owner": "0x1234abcd...",
   "projectId": "abc123",
   "tokenURI": "ipfs://Qm...",
   "metadata": {
      "name": "Launchpad Donor Badge",
      "description": "You have donated to a project on our Launchpad",
      "data": {
         "walletAddress": "0x1234abcd...",
         "projectId": "abc123",
         "projectName": "Green Energy DAO",
         "donationAmount": 100000,
         "projectOwner": "0x5678def...",
         "note": "Thank you for supporting sustainable energy!"
      }
   }
}
```

**Errors:**

-  `404 Not Found` – Không tìm thấy NFT

```json
{
   "error": "NFT not found"
}
```

-  `500 Internal Server Error`

```json
{
   "error": "Failed to retrieve NFT detail"
}
```

---

Nếu bạn cần mình sinh thêm ví dụ cURL test nhanh hoặc bộ Swagger/OpenAPI spec `.yaml` để import vào Postman thì mình có thể hỗ trợ luôn. Bạn muốn không?
