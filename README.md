# DNS Manager Library

[![Go Report Card](https://goreportcard.com/badge/github.com/byteonabeach/dns-manager)](https://goreportcard.com/report/github.com/byteonabeach/dns-manager)
[![Go Reference](https://pkg.go.dev/badge/github.com/byteonabeach/dns-manager.svg)](https://pkg.go.dev/github.com/byteonabeach/dns-manager)


The `dns_manager` library provides a standard interface for managing DNS records, enabling seamless integration with third-party libraries and drivers. This library simplifies the process of adding, deleting, updating, and retrieving DNS records.

## Features

- **Record Types**: Supports various DNS record types including A, AAAA, MX, CNAME, TXT, and a placeholder for non-essential records.
- **Record Management**: Provides an interface for performing actions on DNS records.
- **Flexible Integration**: Designed to work with any third-party DNS management system through a standard interface.

## Installation

To use the `dns_manager` library, simply include it in your Go project:

```go
import "github.com/byteonabeach/dns-manager"
```

## Usage

### Record Types

The library defines several DNS record types:

```go
const (
    None  RecordType = "None essential"
    A     RecordType = "A"
    AAAA  RecordType = "AAAA"
    MX    RecordType = "MX"
    CNAME RecordType = "CNAME"
    TXT   RecordType = "TXT"
)
```

### Record Structure

A DNS record is represented by the `Record` struct:

```go
type Record struct {
    Type    RecordType `json:"type,omitempty"`
    Name    string     `json:"name,omitempty"`
    Content string     `json:"content,omitempty"`
    TTL     uint       `json:"ttl,omitempty"`
}
```

### Actions Interface

The `Actions` interface defines methods for managing DNS records:

```go
type Actions interface {
    AddRecord(*Record) (*Response, error)
    DeleteRecord(*Record) (*Response, error)
    UpdateRecord(*Record) (*Response, error)
    GetRecords() ([]*Record, error)
}
```

### Creating a New Manager

To create a new DNS manager, implement the `manager` interface and use the `New` function:

```go
func New(manager_config manager) (Actions, error) {
    return manager_config.New()
}
```

### Response Structure

The `Response` struct is used to handle responses from DNS operations:

```go
type Response struct {
    Error   string `json:"error"`
    Message string `json:"message"`
}
```

## License

This library is open-source and available under the MIT License.

## Contributing

Contributions are welcome! Please submit a pull request or open an issue for any enhancements or bug fixes.

## Contact

For questions or feedback, please reach out to the maintainer (OxFF, @byteonabeach at GitHub).
