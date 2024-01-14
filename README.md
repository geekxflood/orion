# Orion

## Description

Orion is a web server designed to simplify target configuration for Prometheus.

Prometheus offers powerful service discovery capabilities, but it can be challenging when your resources are not defined in a service discovery method supported natively by Prometheus. This often requires maintaining a static configuration file.

Orion solves this problem by providing a web server that serves a list of targets in a format that Prometheus can understand, using the `http_sd_configs` configuration. With Orion, you can keep your Prometheus configuration static and let it handle the task of serving the targets to Prometheus.

In addition, Orion is designed to be extensible, allowing you to define your own modules for retrieving targets.

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
