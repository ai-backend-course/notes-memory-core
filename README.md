# Notes Memory Core — Go Backend Template

A clean, production-ready Go backend template featuring:

- Go Fiber web framework  
- Postgres 16 (Dockerized)  
- CRUD Notes API  
- In-memory metrics endpoint  
- Structured logging (zerolog)  
- Dockerfile + docker-compose  
- Organized project structure  
- Easily extendable into AI/RAG systems

This template is the foundation for the RAG extension found in the upcoming repository:

notes-memory-core-rag (coming next)

---

## Project Structure

    notes-memory-core/
    │
    ├── main.go
    ├── go.mod
    ├── Dockerfile
    ├── docker-compose.yml
    ├── .env.example
    │
    └── internal/
        ├── database/
        │   └── database.go
        ├── handlers/
        │   └── notes.go
        └── middleware/
            ├── logger.go
            └── metrics.go

---

## Getting Started

### 1. Clone the repository

    git clone https://github.com/ai-backend-course/notes-memory-core.git
    cd notes-memory-core

### 2. Set your environment variables

Copy `.env.example` to `.env` (for local development):

    cp .env.example .env

### 3. Start the API and database using Docker Compose

    docker-compose up --build

The API will be available at:

    http://localhost:8080

---

## API Endpoints

### Health Check

    GET /health

### Get all notes

    GET /notes

### Create a note

    POST /notes
    Content-Type: application/json

    {
      "title": "My note",
      "content": "This is a sample note"
    }

### Get metrics

    GET /metrics

Example response:

    {
      "total_requests": 10,
      "total_errors": 0,
      "avg_latency_ms": 1.2
    }

---

## Tech Stack

| Component   | Technology   |
|------------|--------------|
| Language   | Go 1.22+     |
| Framework  | Fiber v2     |
| Database   | Postgres 16  |
| Driver     | pgx / pgxpool|
| Logging    | Zerolog      |
| Metrics    | Custom middleware |
| Containers | Docker + Docker Compose |

---

## Why This Template Exists

This repo is designed as a clean foundation for backend development in Go.

It intentionally includes:

- Clean HTTP routing  
- Database migrations for a real table (`notes`)  
- Structured logging with zerolog  
- Metrics middleware for basic observability  
- Dockerized Postgres and API container  
- Realistic project layout

It is meant to be extended into:

- Vector search APIs  
- Embeddings pipelines  
- RAG (Retrieval-Augmented Generation) systems  
- Multi-service AI backends  

The next step in this series is the RAG-focused repository:

notes-memory-core-rag

---

## License

MIT License.
