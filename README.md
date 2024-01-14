# Orion

## Description

Orion is a web server that serves target configuration for Prometheus.

## Usage

You can define a local configuration file or use the default one.

```bash
orion run -config /path/to/config/file
```

Configuration format support:

- **YAML**
- **JSON**
- **TOML**

### Configuration

The configuration file is a list of targets with the following format:

```yaml
module: "module_name" # Define which module Orion will have to use, refer to the modules section for more information.
port: "9981" # Define the port on which Orion will listen. Default: 9981
insecure: false # Define if Orion will use TLS or not. Default: false
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
        "targets": [
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
