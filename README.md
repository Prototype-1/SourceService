# SourceService

This is the simulated microservice that acts as a data source for the TargetService in the Teqbae machine task.
It generates random User Profile JSON data each time you call its endpoint.

## Project Overview

- Exposes `/users/changes` endpoint on `:8080`.
- Each call returns 5-10 randomly generated user profiles, with new or updated `last_updated_at`.
- No database, purely in-memory + Go slice.

## How to run

### With Docker
```bash
docker build -t source-service .
docker run -p 8080:8080 source-service
```

### Or with Docker Compose
If you are running it with TargetService, just do:
```bash
docker-compose up --build
```
from the root project.

### Manually with Go
```bash
go run main.go
```

## API Endpoint

### GET /users/changes
Returns a list of randomly generated user profiles.

**Example Response:**
```json
[
  {
    "id": "8f2717d7-bd45-4911-b457-3b23dcbbc556",
    "name": "User145",
    "email": "user145@example.com",
    "mobile": "+911234567890",
    "status": "active",
    "last_updated_at": "2025-07-01T15:43:04+05:30"
  },
  {
    "id": "57ff7c4e-3539-4e55-bb46-5e802e3c7c0e",
    "name": "User89",
    "email": "user89@example.com",
    "mobile": "+911234567891",
    "status": "pending",
    "last_updated_at": "2025-07-01T15:44:04+05:30"
  }
]
```

## Highlights

- **Generates data dynamically** on each request (no persistence).
- Used to test:
  - **Concurrency**
  - **Business validation**
  - **Metrics tracking** in TargetService.

## Related Services

This SourceService works in conjunction with the **TargetService** which:
- Periodically fetches data from this service
- Applies business validation and data persistence
- Provides Prometheus metrics for monitoring

For complete setup instructions and information about the entire data synchronization system, please refer to the [TargetService README](../README.md).

---

*Developed for Teqbae machine task*