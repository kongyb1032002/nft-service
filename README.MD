Hiểu rồi, mình sẽ viết lại README chi tiết hơn, đầy đủ từng phần cụ thể, vừa dễ hiểu, vừa đủ thông tin vận hành & phát triển dự án.

---

# nft-service

`nft-service` là backend service viết bằng Go, cung cấp API để tạo và quản lý NFT dùng làm chứng nhận donation trên blockchain. Metadata của NFT được lưu trên IPFS, smart contract chuẩn ERC-721 được triển khai trên mạng Ethereum-compatible.

---

## Mục lục

1. [Tổng quan dự án](#tổng-quan-dự-án)
2. [Cấu trúc dự án](#cấu-trúc-dự-án)
3. [Yêu cầu hệ thống](#yêu-cầu-hệ-thống)
4. [Cài đặt & khởi chạy](#cài-đặt--khởi-chạy)
5. [Hướng dẫn phát triển](#hướng-dẫn-phát-triển)
6. [Quy trình build và deploy](#quy-trình-build-và-deploy)
7. [API chính](#api-chính)
8. [Hỗ trợ & liên hệ](#hỗ-trợ--liên-hệ)

---

## Tổng quan dự án

`nft-service` phục vụ việc tạo NFT đại diện cho chứng nhận donation của người dùng trong các dự án blockchain. Service tích hợp:

* **IPFS**: Lưu trữ metadata của NFT (tên, mô tả, dữ liệu donation)
* **Smart contract ERC-721**: Mint NFT trên blockchain
* **PostgreSQL**: Lưu trữ dữ liệu ứng dụng, mapping tokenId với người sở hữu
* **REST API**: Cung cấp các endpoint CRUD NFT cho client

---

## Cấu trúc dự án

* **cmd/**: Điểm vào chính ứng dụng (`cmd.go`)
* **configs/**: Cấu hình app & database (`config.go`, `db_config.go`)
* **contracts/NFT/**: Source và ABI smart contract ERC-721
* **controllers/**: HTTP handler logic, tương tác request/response
* **domains/**: Business logic - usecases, repository, lỗi
* **dtos/**: Các struct DTO truyền nhận dữ liệu qua API
* **infrastructure/**: Clients (IPFS, blockchain), middlewares (CORS...)
* **models/**: Định nghĩa entity & data model (metadata, token)
* **pkg/**: Helper và util functions chung
* **servers/**: HTTP server & routing setup
* **services/**: Service xử lý tương tác blockchain, IPFS
* **dist/**: Thư mục chứa binary build & image Docker
* **docs/**: Tài liệu API và spec
* **Makefile**: Các command build, deploy, chạy dev nhanh
* **docker-compose.yaml**: Các dịch vụ phụ trợ (Postgres, IPFS) chạy local dev

---

## Yêu cầu hệ thống

* **Go 1.20+**
* **Docker & Docker Compose** (phiên bản mới để chạy dịch vụ phụ trợ)
* **Server Linux có SSH** để deploy binary & restart service
* Cấu hình hệ thống phù hợp với Go build (Linux amd64)

---

## Cài đặt & khởi chạy

### 1. Clone source về

```bash
git clone https://github.com/yourorg/nft-service.git
cd nft-service
```

### 2. Chuẩn bị môi trường dev

```bash
make dev-setup
```

Lệnh này sẽ:

* Chạy `go mod tidy` và `go mod vendor` để cài dependency
* Copy `.env.example` thành `.env` nếu chưa có
* Khởi động Docker Compose để chạy Postgres, IPFS local

### 3. Khởi chạy app local

```bash
go run cmd/cmd.go
```

Hoặc build và chạy binary:

```bash
make build
./dist/nft-service-v0.0.1
```

---

## Hướng dẫn phát triển

* Các logic business tập trung trong `domains/usecases`
* API handler trong `controllers/`
* Khai báo struct request/response trong `dtos/`
* Tương tác IPFS & blockchain qua `infrastructure/clients`
* Mỗi module đều có test riêng, viết test bổ sung khi cần

---

## Quy trình build và deploy

### Build binary cho môi trường production

```bash
make build
```

Thao tác:

* Build Go binary với GOOS=linux, GOARCH=amd64
* Copy binary qua SCP lên server tại `/opt/nft-service/`
* Restart service Linux bằng systemctl

### Build và chuyển Docker image (chứa các service phụ trợ)

```bash
make build-image
make load-image
```

### Mở SSH tunnel

Giúp forward port database và IPFS khi cần thao tác từ máy local:

```bash
make tunel
```

---

## API chính

* `POST /api/nft` - Tạo NFT donation mới
* `GET /api/nft/wallet/:walletAddress` - Lấy NFT theo ví
* `GET /api/nft/:tokenId` - Chi tiết NFT

Xem tài liệu chi tiết API tại `docs/api.md`.
