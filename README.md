<p align="center">
    <img width=30% src="assets/logo.png">
</p>

# Orion

## Description

Orion is a web server designed to simplify target configuration for Prometheus.

Prometheus offers powerful service discovery capabilities, but it can be challenging when your resources are not defined in a service discovery method supported natively by Prometheus. This often requires maintaining a static configuration file.

Orion solves this problem by providing a web server that serves a list of targets in a format that Prometheus can understand, using the `http_sd_configs` configuration. With Orion, you can keep your Prometheus configuration static and let it handle the task of serving the targets to Prometheus.

In addition, Orion is designed to be extensible, allowing you to define your modules for retrieving targets.

## TODO

In no particular order:

- [ ] Add a customizable module for retrieving targets from a not-supported remote source
- [x] Add a module for retrieving targets from a local file
- [ ] Enable remote cache usage (Redis, Memcached, etc.)
- [ ] Use __meta_ labels to add more information to the targets and allow relabelling from serving the /targets
- [ ] Improve the documentation
- [ ] Add tests
- [ ] Generate and publish OCI images
- [ ] Generate and publish a helm chart

## Usage

You can define a local configuration file or use the default one.

```bash
orion run -config /path/to/config/file
```

Configuration format support:

- **YAML**
- **JSON**
- **TOML**

**Docker**:

You can run it with Docker:

```bash
docker run -d -p 9981:9981 -v /path/to/config/file:/config.yaml ghcr.io/leboncoin/orion:latest /usr/local/bin/orion run --config /config/config.yaml
```

![](assets/buildoci.gif)

### Configuration

The configuration file is a list of targets with the following format:

```yaml
module: "module_name" # Define which module Orion will have to use, refer to the modules section for more information.
port: "9981" # Define the port on which Orion will listen. Default: 9981
insecure: false # Define if Orion will use TLS or not. Default: false
interval: "5" # Define the interval in seconds between each refresh of the targets. Default: 5
```

#### Modules

Orion supports multiple modules to retrieve targets.

- **file**:
  - This module will read a file and return its content.
  - Use the `--local-file` flag to override the file path.
  - If this module is used, it is expected that the configuration file is a list of `endpoints`.
  - Example:
  
    ```yaml
    ---
    module: "file"
    port: "9981"
    insecure: false
    interval: "60"
    endpoints:
      - targets:
          - 10.0.10.2:9100
          - 10.0.10.3:9100
          - 10.0.10.4:9100
          - 10.0.10.5:9100
        labels:
          __meta_datacenter: london
          __meta_prometheus_job: node
      - targets:
          - 10.0.40.2:9100
          - 10.0.40.3:9100
        labels:
          __meta_datacenter: london
          __meta_prometheus_job: alertmanager
    ```

    ```json
    {
        "module": "file",
        "port": "9981",
        "insecure": false,
        "interval": "60",
        "endpoints": [
            {
                "targets": [
                    "10.0.10.2:9100",
                    "10.0.10.3:9100",
                    "10.0.10.4:9100",
                    "10.0.10.5:9100"
                ],
                "labels": {
                    "__meta_datacenter": "london",
                    "__meta_prometheus_job": "node"
                }
            },
            {
                "targets": [
                    "10.0.40.2:9100",
                    "10.0.40.3:9100"
                ],
                "labels": {
                    "__meta_datacenter": "london",
                    "__meta_prometheus_job": "alertmanager"
                }
            }
        ]
    }
    ```

- **http**:
  - This module will perform an HTTP request to a remote endpoint and return its content.
  - Use the key `auth` to define the authentication method.
  - Example:

    ```yaml
    ---
    module: "http"
    port: "9981"
    insecure: false
    interval: "60s"
    timeout: "30s" # Timeout for HTTP requests
    rate_limit: 10 # Requests per minute
    retry:
      attempts: 3
      backoff: "5s"

    logging:
      level: "info" # Options: debug, info, warn, error
      format: "json" # Options: json, plain
      output: "file" # Options: file, stdout
      file_path: "/var/log/myapp.log"

    cache:
      enabled: true
      duration: "10m" # Cache duration
      max_size: 100 # Maximum number of items in cache

    health_check:
      endpoint: "/health"
      interval: "30s"

    metrics:
      prometheus_enabled: true
      custom_metrics_enabled: true
      custom_metrics:
        - name: "my_custom_metric"
          type: "gauge" # Options: gauge, counter, histogram, summary

    module_http:
      auth:
        type: "basic" # Options: basic, token, oauth
        username: "${HTTP_AUTH_USERNAME}"
        password: "${HTTP_AUTH_PASSWORD}"
        token_url: "http://localhost:8080/auth"
        token_ttl: "1h"
      endpoint: "http://localhost:8080/data"
      data_type: "json"
      mapping:
        - targets:
            - "key_in_target"
          labels: 
            - label_name_1: "key_in_filter_1"

    high_availability:
      load_balancing_method: "round_robin" # Options: round_robin, least_connections
      failover_strategy: "next_available" # Options: next_available, random

    security:
      tls_config:
        cert_file: "/path/to/cert.pem"
        key_file: "/path/to/key.pem"
        ca_cert_file: "/path/to/ca.pem"

    dynamic_config:
      reload_enabled: true
      reload_interval: "5m"

    internationalization:
      default_locale: "en_US"
      supported_locales: ["en_US", "fr_FR", "es_ES"]
    ```

## Endpoints

- **/targets**:
  - Returns a list of targets in JSON format compliant with Prometheus `http_sd_configs` configuration.
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

- **/config**:
  - Returns the configuration file in JSON format.
  - Example:

  ```json
  {
    "module": "file",
    "port": "9981",
    "insecure": false,
    "targets": [
      {
        "targets": ["localhost:8080"],
        "labels": {
          "job": "prometheus",
          "instance": "localhost:8080"
        }
      }
    ]
    ```

## Service discovery

Prometheus can discover new targets using service discovery (`sd`). Service discovery allows Prometheus to automatically find and monitor new targets without manual configuration.

There are various service discovery methods supported by Prometheus, including DNS queries, values files, proprietary software (Consul, Puppet, Eureka, etc.), and dedicated cloud/infrastructure providers (AWS, GCE, etc.). One interesting provider is `http_sd`, which allows Prometheus to fetch targets from a specified URL.

For example, the `http_sd` configuration can include the following options:
This is the extract from the doc:

```yaml
# URL from which the targets are fetched.
url: <string>

# Refresh interval to re-query the endpoint.
[ refresh_interval: <duration> | default = 60s ]

# Authentication information used to authenticate to the API server.
# Note that `basic_auth`, `authorization` and `oauth2` options are
# mutually exclusive.
# `password` and `password_file` are mutually exclusive.

# Optional HTTP basic authentication information.
basic_auth:
  [ username: <string> ]
  [ password: <secret> ]
  [ password_file: <string> ]

# Optional `Authorization` header configuration.
authorization:
  # Sets the authentication type.
  [ type: <string> | default: Bearer ]
  # Sets the credentials. It is mutually exclusive with
  # `credentials_file`.
  [ credentials: <secret> ]
  # Sets the credentials to the credentials read from the configured file.
  # It is mutually exclusive with `credentials`.
  [ credentials_file: <filename> ]

# Optional OAuth 2.0 configuration.
oauth2:
  [ <oauth2> ]

# Optional proxy URL.
[ proxy_url: <string> ]
# Comma-separated string that can contain IPs, CIDR notation, domain names
# that should be excluded from proxying. IP and domain names can
# contain port numbers.
[ no_proxy: <string> ]
# Use proxy URL indicated by environment variables (HTTP_PROXY, https_proxy, HTTPs_PROXY, https_proxy, and no_proxy)
[ proxy_from_environment: <boolean> | default: false ]
# Specifies headers to send to proxies during CONNECT requests.
[ proxy_connect_header:
  [ <string>: [<secret>, ...] ] ]

# Configure whether HTTP requests follow HTTP 3xx redirects.
[ follow_redirects: <boolean> | default = true ]

# Whether to enable HTTP2.
[ enable_http2: <boolean> | default: true ]

# TLS configuration.
tls_config:
  [ <tls_config> ]
```

## Project objectives

This project aims to build a fast and reliable `http_sd_configs` provider endpoint for Prometheus.

This provider will cache and refresh its data at regular intervals. It must be highly available and allow for node deficiency in a multiple-instance deployment.

Example of a sequence diagram big picture:

**File Module**

```mermaid
sequenceDiagram
    participant prometheus
    participant orion
    participant exporterA
    participant targetA
    loop Every X seconds
        prometheus->>orion: GET /targets
        orion->>prometheus: response: tragets:[targetA[{}]]
        critical Failed to GET /targets
            prometheus->>prometheus: reuse same target values
        end
    end
    Note over prometheus,exporterA: prometheus relabelling target
    prometheus->>exporterA: GET /metrics?targetA
    exporterA->>targetA: exporter query targetA
    targetA->>exporterA: response
    exporterA->>prometheus: /metrics for targetA
    
```

**HTTP Module**

```mermaid
sequenceDiagram
    participant sourceB
    participant orion
    participant prometheus
    participant exporterB
    participant targetB
    loop Every X seconds
      orion->>sourceB: GET /data
      sourceB->>orion: response: data
      critical Failed to GET /data
          Note over orion: retry and backoff and meanwhile serve cached data
          orion->>sourceB: GET /data
      end
      orion->>orion: cache data
    end
    loop Every X seconds
        prometheus->>orion: GET /targets
        orion->>prometheus: response: tragets:[targetB[{}]]
        critical Failed to GET /targets
            prometheus->>prometheus: reuse same target values
        end
    end
    Note over prometheus,exporterB: prometheus relabelling target
    prometheus->>exporterB: GET /metrics?targetB
    exporterB->>targetB: exporter query targetB
    targetB->>exporterB: response
    exporterB->>prometheus: /metrics for targetB
```

## Improvement

### http_module

- Extend the configuration to support this wide array of features:

```yaml
---
module: "http"
port: "9981"
insecure: false # Define if Orion will use TLS or not. Default: false
interval: "60s" # Define the interval in seconds between each refresh of the targets. Default: 5
timeout: "30s" # Timeout for HTTP requests
rate_limit: 10 # Requests per minute

# Retry configuration (optional)
retry:
  attempts: 3
  backoff: "5s"

# Logging configuration (optional)
logging:
  level: "info" # Options: debug, info, warn, error
  format: "json" # Options: json, plain
  output: "file" # Options: file, stdout
  file_path: "/var/log/myapp.log"

# Cache configuration (optional)
cache:
  enabled: true
  duration: "10m" # Cache duration
  max_size: 100 # Maximum number of items in cache

# Health check configuration (optional)
health_check:
  endpoint: "/health"
  interval: "30s"

# Prometheus metrics configuration (optional)
metrics:
  prometheus_enabled: true
  custom_metrics_enabled: true
  custom_metrics:
    - name: "my_custom_metric"
      type: "gauge" # Options: gauge, counter, histogram, summary

# HTTP module configuration (required)
module_http:

  # Authentication configuration (optional)
  auth:
    type: "${HTTP_AUTH_TYPE}" # Placeholder for auth type: basic, token, oauth, none
    credentials:
      username: "${HTTP_AUTH_USERNAME}"
      password: "${HTTP_AUTH_PASSWORD}"
    token:
      url: "${HTTP_AUTH_TOKEN_URL}" # URL to retrieve the token
      request_method: "POST" # HTTP method for token retrieval: GET, POST
      request_body: "${HTTP_AUTH_TOKEN_REQUEST_BODY}" # JSON, form-encoded data, etc.
      headers: # Additional headers for token request
        Content-Type: "application/json"
      ttl: "1h" # Time-to-live for token
    oauth:
      client_id: "${OAUTH_CLIENT_ID}"
      client_secret: "${OAUTH_CLIENT_SECRET}"
      auth_url: "${OAUTH_AUTH_URL}"
      scopes: ["scope1", "scope2"] # OAuth scopes
  
  # Endpoint configuration (required)
  request:
    method: "GET" # HTTP method: GET, POST, PUT, DELETE
    url: "http://localhost:8080/data"
    headers: # Custom headers for the data request
      Accept: "application/json"
    params: # Query parameters for the data request
      param1: "value1"
      param2: "value2"
    timeout: "30s" # Timeout for the data request
  data_type: "json" # Expected data type: json, xml, text
  data_mapping:
    json_path: "data.targets" # JSONPath or XMLPath to locate the desired data
    mappings: # Mapping of data to Prometheus format
      - targets:
          - json_path: "ip_address"
        labels: 
          label_name_1: "json_path: label_value_1"
          label_name_2: "json_path: label_value_2"

# High availability configuration (optional)
high_availability:
  load_balancing_method: "round_robin" # Options: round_robin, least_connections
  failover_strategy: "next_available" # Options: next_available, random

# TLS configuration (optional)
security:
  tls_config:
    cert_file: "/path/to/cert.pem"
    key_file: "/path/to/key.pem"
    ca_cert_file: "/path/to/ca.pem"

# Dynamic configuration for the HTTP module (optional)
dynamic_config:
  reload_enabled: true
  reload_interval: "5m"

# Internationalization configuration for the HTTP module (optional)
internationalization: 
  default_locale: "en_US"
  supported_locales: ["en_US", "fr_FR", "es_ES"]
```
