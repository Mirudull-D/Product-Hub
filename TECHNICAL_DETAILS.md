# Technical Specifications and Performance Analysis

This document provides in-depth technical details regarding the architecture, performance benchmarks, and operational configuration of the Product-Hub service.

---

## Performance Benchmarks (k6)

The system is validated using k6 load testing scripts to simulate realistic production traffic patterns.

### Load Profile (load_test.js)
The performance evaluation follows a structured load curve:
- Ramp-up Phase: Linear increase from 0 to 70 Virtual Users (VUs) over 30 seconds.
- Peak Load Phase: Sustained pressure of 100 VUs for 60 seconds.
- Ramp-down Phase: Gradual reduction to 50 VUs over 30 seconds.

### Service Level Objectives (SLOs)
The following thresholds are enforced during performance validation:
- Latency (http_req_duration): p95 must remain below 1000ms.
- Transactional Integrity: Zero failures allowed during stock deduction and order creation cycles.
- Cache Efficiency: Product listing responses should maintain sub-100ms latency when served from Redis.

---

## API Reference (v1)

### Identity Service

#### POST /register
Creates a new user profile.
- Payload Requirements:
  - firstName (required)
  - lastName (required)
  - email (required, unique, valid format)
  - password (required, min 3 chars)
- Success Response: 201 Created

#### POST /login
Authenticates credentials and issues a JWT.
- Payload Requirements:
  - email (required)
  - password (required)
- Success Response: 201 Created
- Output: JSON object containing "token" string.

### Commerce Service

#### GET /product
Retrieves all available products.
- Caching Strategy: First attempt reads from Redis ("products" key). On cache miss, queries PostgreSQL and hydrates the cache with a 10-minute TTL.
- Success Response: 200 OK
- Output: Array of Product objects.

#### POST /cart/checkout
Processes an order. Requires valid JWT in Authorization header.
- Transaction Logic:
  1. Validates user existence from context.
  2. Verifies stock availability for all requested items.
  3. Begins PostgreSQL transaction.
  4. Deducts quantity from products table.
  5. Inserts record into orders table.
  6. Inserts multiple records into order_items table.
  7. Commits transaction.
- Success Response: 201 Created
- Output: JSON object containing "orderId" and "totalPrice".

---

## Database Schema and Architecture

The persistence layer is designed for high relational integrity.

### Tables
- users: Stores identity data and hashed passwords (bcrypt).
- products: Inventory management with fields for description, image URL, quantity, and price.
- orders: Tracks order status (default: 'pending') and total cost, linked to user ID.
- order_items: Normalized table linking products to orders, capturing price at time of purchase to handle future price fluctuations.

### Constraints
- Foreign Keys: orders.user_id (ON DELETE CASCADE), order_items.order_id (ON DELETE CASCADE), order_items.product_id (ON DELETE RESTRICT).
- Types: Prices are stored as NUMERIC(10,2) to prevent rounding errors associated with floating-point math.

---

## Containerization and Deployment

### Dockerfile Specification
The project uses a multi-stage Docker build for optimized image size:
1. Builder Stage: Uses golang:1.26 to compile the binary.
2. Runtime Stage: Uses debian:bookworm-slim for the final execution environment.

### Deployment Commands
Build:
```bash
docker build -t product-hub:latest .
```

Run:
```bash
docker run -p 8080:8080 \
  -e connString="postgres://user:password@host.docker.internal:5432/dbname" \
  -e redisAddr="host.docker.internal:6379" \
  product-hub:latest
```

---

## Project Organization

The codebase follows a modular structure to separate concerns:
- cmd/: Application entry points and initialization.
- config/: Centralized environment and configuration management.
- db/: SQLC definitions, generated code, and schema migrations.
- redis/: Redis client factory and configuration.
- service/: Domain logic implementations (User, Product, Cart, Order).
- types/: Shared data structures and interfaces.
- utils/: Common utility functions for JSON handling and validation.
