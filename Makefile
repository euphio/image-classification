# Variables
PROTO_DIR=proto
PYTHON_SERVICE=python_service
GO_SERVICE=go_service
VENV_DIR=$(PYTHON_SERVICE)/.venv

# --- Proto Generation ---
proto-python:
		$(VENV_DIR)/bin/python -m grpc_tools.protoc \
		-I$(PROTO_DIR) \
		--python_out=$(PYTHON_SERVICE) \
		--grpc_python_out=$(PYTHON_SERVICE) \
		$(PROTO_DIR)/image_recognition.proto

proto-go:
	protoc -I$(PROTO_DIR) \
		--go_out=. \
		--go-grpc_out=. \
		$(PROTO_DIR)/image_recognition.proto


proto: proto-python proto-go

# --- Setup (first time dev environment) ---
setup-python:
	python3 -m venv $(VENV_DIR)
	$(VENV_DIR)/bin/pip install --upgrade pip
	$(VENV_DIR)/bin/pip install -r $(PYTHON_SERVICE)/requirements.txt

setup-go:
	cd $(GO_SERVICE) && go mod tidy
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

setup: setup-python setup-go proto

# --- Running services locally ---
run-python-service:
	$(VENV_DIR)/bin/python $(PYTHON_SERVICE)/main.py

run-go-service:
	cd $(GO_SERVICE) && go run main.go

# --- Docker Compose for full stack ---
docker-build:
	docker compose build

docker-up:
	docker compose up

docker-down:
	docker compose down

# --- Dev Container Utilities ---
devcontainer-build:
	devcontainer build .

devcontainer-up:
	devcontainer up .

# --- Lint and Format (optional, recommended) ---
lint-python:
	$(VENV_DIR)/bin/flake8 $(PYTHON_SERVICE)

format-python:
	$(VENV_DIR)/bin/black $(PYTHON_SERVICE)

lint-go:
	cd $(GO_SERVICE) && golangci-lint run

format-go:
	cd $(GO_SERVICE) && gofmt -w .

lint: lint-python lint-go
format: format-python format-go