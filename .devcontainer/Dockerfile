# This is now based on the pre-built Python image
FROM mcr.microsoft.com/devcontainers/python:3.11-bookworm

# Install protobuf compiler (needed by both Python and Go)
RUN apt-get update && apt-get install -y protobuf-compiler

# Python venv setup
WORKDIR /workspace
RUN python3 -m venv /workspace/.venv
ENV PATH="/workspace/.venv/bin:$PATH"

# Install Python deps
COPY python_service/requirements.txt /workspace/python_service/requirements.txt
RUN pip install --upgrade pip && pip install -r /workspace/python_service/requirements.txt

# Default CMD for keeping container alive
CMD ["sleep", "infinity"]
