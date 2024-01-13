# Orion

## Description

Orion is a web server that serve targets configuration for `prometheus`.

## Implementation Details

- **Language**: Go (Golang)
- **Libraries**:
  - Cobra for CLI and configuration management.
  - HTTP package for the server implementation.

## Endpoints

- **/targets**:
  - Returns a list of targets in JSON format.
  - Example:

    ```json
    [
      {
        "targets": ["localhost:8080"],
        "labels": {
          "job": "prometheus",
          "instance": "localhost:8080"
        }
      }
    ]
    ```

- **/health**:
  - Returns a 200 OK response if the server is healthy.
  - Returns a 500 Internal Server Error response if the server is unhealthy.

- **/ready**:
  - Returns a 200 OK response if the server is ready.
  - Returns a 500 Internal Server Error response if the server is not ready.

## Workflow

### 1. Initialization

- **Configuration Handling**:
  - Utilize the Cobra library for command-line argument parsing.
  - Load configuration through command flags or a config file.

### 2. Cache Management (Go Routine 1)

- **Cache Generation and Updating**:
  - Implement a caching mechanism (in-memory or external like Redis).
  - Initially populate the cache.
  - Continuously update the cache based on predefined triggers or intervals.

### 3. HTTP Server (Go Routine 2)

- **Serving Cache Values**:
  - Start an HTTP server after the initial cache is populated.
  - Serve cached values over HTTP endpoints.
  - Include readiness endpoints for Kubernetes integration.

### 4. Continuous Operation

- **Ongoing Cache Updates**:
  - Ensure the first Go routine runs indefinitely.
  - Periodically update the cache as required.

