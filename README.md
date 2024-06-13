<div style="display: flex; justify-content: center;">
    <img style="width: 50%; height: auto;" src="assets/logo_nobg.png" alt="logo">
</div>

# Orion

## Table of Contents

<!-- TOC -->
* [Orion](#orion)
  * [Table of Contents](#table-of-contents)
  * [Project Description](#project-description)
  * [Key Features](#key-features)
  * [Endpoints](#endpoints)
  * [Usages](#usages)
    * [Initialize Configuration](#initialize-configuration)
    * [Start Orion](#start-orion)
  * [Configuration](#configuration)
    * [Configuration Structure](#configuration-structure)
    * [Module Configuration](#module-configuration)
      * [File](#file)
      * [REST](#rest)
<!-- TOC -->

## Project Description

Prometheus is renowned for its robust service discovery capabilities, but complexities arise when dealing with
resources that fall outside its native service discovery methods. This often necessitates the maintenance of static
configuration files, a process that can be cumbersome and error-prone. Orion  addresses this challenge by providing
a versatile web server capable of serving a dynamically generated list of targets. These targets are formatted for
seamless compatibility with Prometheus, specifically utilizing the `http_sd_configs` configuration.

`Orion` is an advanced web server engineered to streamline the configuration and management of Prometheus targets. It
adeptly handles queries across a diverse range of data sources, utilizing various protocols such as REST API, SOAP,
GraphQL, SQL, gRPC, and file-based scraping. The system stands out for its high modularity; each protocol is managed
by a dedicated module, ensuring seamless integration and consistent handling.

Orion's design philosophy emphasizes extensibility and adaptability. It allows for the creation and integration of 
custom modules, enabling users to tailor the system for retrieving targets from a wide array of data sources. 
This flexibility makes Orion an invaluable tool for organizations looking to harness the full potential of Prometheus 
in a variety of complex and evolving technological landscapes.

## Key Features

- **Modular Design**: Separate modules for each data source type (REST, SOAP, GraphQL, ...) and Prometheus target retrieval.
- **Common Interface**: A consistent interface across all modules for initialization, request handling, and response parsing.
- **Configurable**: Uses YAML, JSON, or TOML configuration files for defining query parameters and target settings.
- **Concurrent Execution**: Leverages Go's goroutines for concurrent query execution.
- **Extensible**: Easily adaptable to accommodate new data sources.

## Endpoints

- **/targets**: Returns a list of targets in JSON format compliant with Prometheus `http_sd_configs`.
- **/config**: Returns the configuration file in JSON format.

## Usages

You can define a local configuration file or use the default one.

### Initialize Configuration

Initialize a new configuration file. The command will create a new configuration file in the `$HOME/.orion` directory.

```bash
orion init
```

### Start Orion

```bash
orion run --config /path/to/config/file
```

**Docker**:

Run Orion with Docker:

```bash
docker run -d -p 9981:9981 -v /path/to/config/file:/config.yaml ghcr.io/geekxflood/orion:latest /usr/local/bin/orion run --config /config.yaml
```

## Configuration

The configuration file is a list of targets with settings for different data sources and protocols.

### Configuration Structure

The configuration file supports an array of targets, each specifying the data source type and relevant settings.

```yaml
targets:
  - type: "REST"
    url: "https://example.com/api"
    method: "GET"
    headers: { }
    response_type: "json"
    parser_rules: { }
    timeout: "30s"
    retry_count: 3
    retry_interval: "1s"

  - type: "File"
    file_path: "/path/to/local/file.json"
    file_format: "json"

module: "module_name"
port: "9981"
insecure: false
interval: "5s"
```

### Module Configuration

When defining a target, you can specify the module to use for querying the data source. 
When default value are define for the module, the field become optional and a fallback to the default value.

- [File](doc/file.md)
- [REST](doc/rest.md)
