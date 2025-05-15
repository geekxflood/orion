# Orion Development Guidelines

This document provides guidelines and instructions for developing and maintaining the Orion project.

## Build/Configuration Instructions

### Building the Project

Orion is a Go project that follows standard Go build practices.

#### Local Build

```bash
# Build the project
go build -o orion .

# Run the built binary
./orion
```

#### Docker Build

```bash
# Build the Docker image
docker build -t orion:latest -f build/Dockerfile .

# Run the Docker container
docker run -d -p 9981:9981 -v /path/to/config.yaml:/config.yaml orion:latest /usr/local/bin/orion run --config /config.yaml
```

### Configuration

Orion uses a configuration file to define its behavior. The default location is `$HOME/.orion/config.yaml`.

#### Initialize Configuration

```bash
# Create a default configuration file
orion init
```

This will create a configuration file at `$HOME/.orion/config.yaml` with default values.

#### Configuration Format

Orion supports YAML, JSON, and TOML configuration formats. The configuration structure is as follows:

```yaml
# YAML example
module: module_name  # The module to use (e.g., "file", "http")
port: 9981           # The port to listen on
insecure: false      # Whether to disable TLS verification
interval: 5          # The interval for refreshing the configuration
endpoints:           # The list of endpoints
  - targets:         # The list of target addresses
      - "target1"
      - "target2"
    labels:          # Key-value pairs for labels
      label1: "value1"
      label2: "value2"
```

## Testing Information

### Running Tests

Orion uses Go's standard testing package. To run tests:

```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run tests for a specific package
go test -v ./internal/helpers
```

### Adding Tests

Tests should be added in files with the `_test.go` suffix, in the same package as the code being tested.

#### Example Test

Here's an example of a test for the `EncodeBase64` function in the `helpers` package:

```go
package helpers

import (
	"testing"
)

func TestEncodeBase64(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Simple string",
			input:    "hello",
			expected: "aGVsbG8=",
		},
		{
			name:     "Complex string",
			input:    "Hello, World! 123",
			expected: "SGVsbG8sIFdvcmxkISAxMjM=",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := EncodeBase64(tt.input)
			if result != tt.expected {
				t.Errorf("EncodeBase64(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}
```

### Test Coverage

To check test coverage:

```bash
# Generate test coverage report
go test -cover ./...

# Generate detailed coverage report
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## Additional Development Information

### Project Structure

Orion follows the standard Go project layout:

- `cmd/`: Contains the command-line interface code
- `internal/`: Contains the internal packages
  - `helpers/`: Utility functions
  - `httpclient/`: HTTP client implementation
  - `localtypes/`: Data structures
  - `modules/`: Module implementations
  - `unmarshaller/`: Configuration unmarshalling
- `build/`: Contains build-related files (e.g., Dockerfile)
- `doc/`: Contains documentation
- `config/`: Contains example configuration files

### Module Development

When developing a new module:

1. Create a new file in the `internal/modules` directory with the name `module_name.go`
2. Create a corresponding test file `module_name_test.go`
3. Implement the module functionality
4. Add the module to the switch statement in `cmd/run.go`

### Code Style

Orion follows standard Go code style and best practices:

- Use `gofmt` or `goimports` to format code
- Follow the [Effective Go](https://golang.org/doc/effective_go) guidelines
- Use meaningful variable and function names
- Add comments for exported functions and types
- Keep functions small and focused on a single responsibility
- Use error handling instead of panics for recoverable errors

### Debugging

For debugging, you can use the following techniques:

- Use the `-v` flag for verbose output
- Set the `insecure` flag to `true` to disable TLS verification for debugging
- Check the logs for error messages

### Continuous Integration

The project uses GitHub Actions for continuous integration. The workflow includes:

- Building the project
- Running tests
- Building and pushing Docker images

### Release Process

To release a new version:

1. Update the version in the code
2. Create a new tag with the version number
3. Push the tag to GitHub
4. The CI/CD pipeline will build and publish the release
