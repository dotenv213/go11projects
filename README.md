# 🚀 11 Essential Go Projects: From Zero to Hero

Welcome to my Golang learning journey! This repository contains **11 practical projects** demonstrating the path from basic HTTP servers to advanced Cloud-Native microservices, Real-time systems, and Clean Architecture.

Each project focuses on specific concepts required for a modern Backend Engineer.

## 📂 Projects Overview

| # | Project Name | Key Concepts Learned | Tech Stack |
|---|---|---|---|
| 01 | [**Simple Web Server**](./01-simple-server) | `net/http`, Static Files, Form Parsing | Go Std Lib |
| 02 | [**Movies CRUD API**](./02-movies-crud) | REST API, Slices, JSON Handling | Fiber v2 |
| 03 | [**Bookstore (Dockerized)**](./03-bookstore) | **MySQL**, **GORM**, Docker Compose | Fiber, GORM, Docker |
| 04 | [**JWT Authentication**](./04-jwt-auth) | **Security**, Hashing, Tokens, Cookies | JWT v5, Bcrypt |
| 05 | [**Concurrent Checker**](./05-concurrent-checker) | **Concurrency**, Goroutines, Channels, WaitGroup | Go Concurrency |
| 06 | [**URL Shortener**](./06-url-shortener) | **Redis**, Caching, Rate Limiting | Fiber, Redis |
| 07 | [**Real-time Chat**](./07-realtime-chat) | **WebSockets**, Mutex, Broadcast Pattern | Fiber, JS WebSocket |
| 08 | [**Cloud File Uploader**](./08-s3-uploader) | **Object Storage**, S3/MinIO, Multipart Forms | MinIO SDK |
| 09 | [**gRPC Microservice**](./09-grpc-service) | **gRPC**, **Protobuf**, Code Gen, RPC | gRPC-Go |
| 10 | [**Weather CLI Tool**](./10-weather-cli) | **CLI Tools**, API Consumption, Cobra | Cobra, OpenMeteo |
| 11 | [**Clean Banking System**](./11-clean-bank) | **Clean Architecture**, **TDD**, Mocking, DI | Testify, Mockery |

## 🛠️ How to Run
Each project folder contains its own `README.md` with specific instructions. Generally, you can navigate to a folder and run:

```bash
cd 01-simple-server
go run main.go
```

For dockerized projects (like 03, 06, 08): 
```bash
docker-compose up -d
go run main.go
```

👨‍💻 About Me
I am a passionate Backend Developer transitioning to Go. This repository showcases my ability to build scalable, concurrent, and maintainable software using Go's ecosystem.