# yagokv

Yet Another Key-Value store written in Go, without any external dependencies.

**This project is not meant for production usage, just as a playground**

## Overview

yagokv is designed to be a lightweight key-value store. The project uses a bucket-based storage mechanism with concurrent safe operations. It also provides an API layer for managing GET, SET, and DELETE operations and an HTTP server for handling client requests.

## Architecture

- **Storage Layer:**  
  Implements the core in-memory key-value store under the `internal/kvs` package. It uses multiple buckets with their own read-write mutexes to provide thread-safe and efficient concurrent access.

- **API Layer:**  
  Exposes HTTP endpoints for basic CRUD operations. This layer abstracts the underlying storage and allows for future scaling enhancements.

- **Server:**  
  The HTTP server (see `internal/server/server.go`) initializes and serves your requests, integrating the API and storage components.

## Getting Started

### Prerequisites

- Go 1.23.x (ensure you have the proper version installed as per `go.mod`)

### Building and Running

To build the project:

```sh
go build -o yagokv
```

Then, run the executable:

```sh
./yagokv
```

### Testing

Unit tests for the key-value store are located in the kvs directory. Run them with:

### Future Enhancements

### License

This project is licensed under the MIT License.
