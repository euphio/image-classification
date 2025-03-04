
# Image Classification Microservice

This project is a **microservice-based image classification system** using **ImageNet**. It consists of two services:

- A **Python service** (using gRPC) that handles image classification.
- A **Go service** (using HTTP and gRPC) that serves as a gateway to the Python service.

This setup includes full support for **VS Code Dev Containers**, **Docker Compose**, and a comprehensive `Makefile` for managing development tasks.

---

## Quickstart

### Clone the Repository

```sh
git clone <repository-url>
cd <project-folder>
```

### Running (On Host Machine)

For running both services together in Docker:

```sh
make docker-build
make docker-up
# When done
make docker-down
```

---

### Start using Dev Container (recommended)

This project comes with a **VS Code Dev Container** setup, allowing you to develop in a fully containerised environment with Python and Go pre-installed.

1. Open the project in VS Code.
2. When prompted, **Reopen in Container**.
3. Run `make setup` to initialise dependencies.

### Manual Setup (if not using Dev Container)

#### Requirements

- Python 3.11+
- Go 1.22+
- Docker & Docker Compose

#### Setup

```sh
make setup
```

---

## Architecture

```
+--------------------+     gRPC      +--------------------------+
| Go Service (API)   | <-----------> | Python Service           |
| (HTTP Gateway)     |               | (Image Classification)   |
+--------------------+               +--------------------------+
```

---

## Development Workflow

### Proto Generation

Whenever you update the protobuf file (`proto/image_recognition.proto`), regenerate the stubs:

```sh
make proto
```

### Running Services on Dev Machine

Run the services locally:

```sh
make run-python-service  # Python classification service
make run-go-service      # Go gateway service
```

---

## VS Code Dev Container Commands

If using **Dev Containers**, you can also manage the container directly:

```sh
make devcontainer-build  # Build the dev container image
make devcontainer-up     # Start the dev container
```

---

## Lint and Format

```sh
make lint    # Check formatting/linting
make format  # Auto-format code
```

---

## Summary of Makefile Targets

| Target             | Description                                     |
|-------------------|-------------------------------------------------|
| `setup`           | Install dependencies (Python + Go)              |
| `proto`           | Regenerate gRPC stubs                           |
| `run-python-service` | Run Python service                        |
| `run-go-service`     | Run Go service                            |
| `docker-build`    | Build Docker images                             |
| `docker-up`       | Start services with Docker Compose              |
| `docker-down`     | Stop services with Docker Compose               |
| `devcontainer-build` | Build VS Code Dev Container                 |
| `devcontainer-up`    | Start VS Code Dev Container                 |
| `lint`            | Run linters                                     |
| `format`          | Run formatters                                  |

---

## Notes

- The project uses **ImageNet** models via PyTorch inside the Python service.
- Protobuf stubs for Python and Go are automatically generated via `make proto`.

---

## License

MIT