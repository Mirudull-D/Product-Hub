# Product-Hub

Product-Hub is a high-performance e-commerce backend service implemented in Go. It utilizes PostgreSQL for transactional data persistence and Redis for distributed caching to ensure scalability under high load.

[View Comprehensive Technical Specifications](./TECHNICAL_DETAILS.md)

---

## Performance Overview

The system is optimized for low-latency response times and high throughput:
- p95 Latency: < 1000ms under 100 concurrent virtual users.
- Error Rate: < 1% during peak load testing.
- Database Transactions: 100% ACID compliance for order processing.
- Caching: Redis-backed product retrieval with 10-minute TTL.

---

## API Endpoints

### Identity and Health
- POST /api/v1/register: User account creation.
- POST /api/v1/login: Authentication and JWT issuance.
- GET /api/v1/health: System and database health status.

### Products and Commerce
- GET /api/v1/product: Retrieve product catalog (cached).
- POST /api/v1/cart/checkout: Process order and manage inventory (authenticated).

---

## Technology Stack

- Runtime: Go 1.21+
- Database: PostgreSQL (Persistence), Redis (Distributed Cache)
- API: Gorilla Mux (Routing), JWT (Authentication)
- Persistence Layer: SQLC (Type-safe query generation), pgx (Driver)
- Validation: go-playground/validator

---

## Quick Start

### 1. Configuration
Initialize the environment configuration:
```bash
cp .env.example .env
```

### 2. Execution
Download dependencies and start the service:
```bash
go mod tidy
go run cmd/main.go
```

### 3. Verification
Execute the health check:
```bash
curl http://localhost:8080/api/v1/health
```

---

## Testing

Unit Testing:
```bash
go test ./...
```

Load Testing (k6):
```bash
k6 run load_test.js
```
