# GRP Server

## Introduction

A personal **enterprise-grade** server, written in Go with modular architecture, ready for high scalability, robust security and optimized performance. Exposure to AWS via API Gateway → EC2, it uses managed services (DynamoDB, S3, Redis) and end-to-end encryption (SSH, AES-GCM). This server was designed to be used by me (and by other people as there is support for this) on a daily basis, from anywhere with internet and a Browser.

---

## Key Features

### 1. File Management

- Upload/Download files via **AWS S3**
- Write-behind Redis cache for instant responses
- Duplicate checking, type and size control
- Managed download, resumption support

### 2. Chat System

- Full message history
- Metadata (date, time, user)
- Message size control

### 3. Metrics System

- Files by extension
- Records by domain
- Storage by type and domain

### 4. Tag Manager

- Full tag CRUD
- Advanced search and filters
- Redis cache with configurable TTL (120s)

### 5. Password Manager

- AES-GCM encryption
- Secure storage in DynamoDB
- Change history

### 6. e-books

- Digital library with metadata
- Upload/Download and categorization

### 7. Task System

- Task CRUD
- Prioritization, status tracking and deadlines

### 8. Authentication and Security

- JWT signed with SSH private key
- Refresh tokens, rate limiting
- HTTPS/TLS via AWS API Gateway
- Strict input validation

---

## Infrastructure and Cloud Computing (AWS)

- **AWS EC2**: Go server hosting
- **AWS API Gateway**: Reverse proxy, CORS, authentication, throttling
- **AWS S3**: File storage
- **DynamoDB**: Core NoSQL database
- **Redis**: High-speed cache, TTL invalidation

---

## Project Structure

![image](https://github.com/user-attachments/assets/d27b8c2e-0238-448c-9352-dd32a203652f)



    grp@server/
    ├── cmd/
    │   └── main.go                # Entry point: initializes middlewares, routes, and services
    ├── core/
    │   ├── api/                   # Handlers and routes per domain
    │   │   ├── admin/             # /api/admin/logs
    │   │   ├── auth/              # /api/auth/*
    │   │   ├── chat/              # /api/chat/*
    │   │   ├── download/          # /api/download
    │   │   ├── e-books/           # /api/e-books/*
    │   │   ├── files/             # /api/files/*
    │   │   ├── health/            # /health
    │   │   ├── metrics/           # /api/metrics
    │   │   ├── passwords/         # /api/passwords/*
    │   │   ├── tags/              # /api/tags/*
    │   │   ├── tasks/             # /api/tasks/*
    │   │   └── users/             # /api/users/*
    │   ├── db/                    # Persistence layer
    │   │   ├── dynamo/            # DynamoDB operations
    │   │   ├── redis/             # Redis client and cache logic
    │   │   └── s3/                # AWS S3 operations
    │   ├── domains/               # Business logic
    │   ├── middleware/            # Authentication, cache, logs
    │   ├── crypto/                # Hash, JWT, AES, SSH utilities
    │   ├── utils/                 # General helpers
    │   └── config/                # Parsing .env, configurations
    ├── docs/                      # Swagger (swagger.yaml/json)
    ├── ssh/                       # SSH keys for JWT
    ├── storage/                   # Temporary upload storage
    ├── Makefile                   # Build, run, documentation scripts
    └── .env                       # Environment variables

This structure reflects the organization of my project, highlighting the modularity of the domains and components responsible for business logic, persistence, authentication, encryption, cache, among others.

---

That's it, this was my first project in golang, using aws services, and made specifically for me, you can see the front-end of this project in the link below as soon as it is ready:

https://github.com/Grs2080w/grp-front.git
