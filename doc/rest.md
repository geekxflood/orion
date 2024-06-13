# REST

Configuration field for the `REST` module.

| Field          | Description                                 | Value Type        | Default Value | Available Values |
|----------------|---------------------------------------------|-------------------|---------------|------------------|
| url            | The URL of the REST API endpoint.           | string            |               |                  |
| method         | The HTTP method to use for the request.     | string            | `GET`         | `GET`, `POST`    |
| headers        | The headers to include in the request.      | map[string]string | {}            |                  |
| response_type  | The type of response expected from the API. | string            | `json`        | `json`           |
| parser_rules   | Rules for parsing the response.             | map[string]string | {}            |                  |
| timeout        | The duration before the request times out.  | string            | 5s            |                  |
| retry_count    | The number of times to retry the request.   | int               | 3             |                  |
| retry_interval | The interval between retries.               | string            | 1s            |                  |
