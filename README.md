# go-connection-pool

Connection pool is a thread safe list of net.Conn

[![Go Report Card](https://goreportcard.com/badge/github.com/sokil/go-connection-pool)](https://goreportcard.com/report/github.com/sokil/go-connection-pool)
[![GoDoc](https://godoc.org/github.com/sokil/go-connection-pool?status.svg)](https://godoc.org/github.com/sokil/go-connection-pool)
[![Code Climate](https://codeclimate.com/github/sokil/go-connection-pool/badges/gpa.svg)](https://codeclimate.com/github/sokil/go-connection-pool)
## Basic usage

```go
socket, err := net.Listen("tcp", "127.0.0.1:8080")
  
// prepare connection pool
connectionPool := connectionPool.NewConnectionPool()
  
// accept connection
connection, err := socket.Accept()
    
// add connection to pool
connectionId := connectionPool.Add(connection)

// get connection and read
reader := bufio.NewReader(connectionPool.Get(connectionId))

// count of connections in pool
size := connectionPool.Size()

// send message to all connections in pool
connectionPool.Range(func(targetConnection net.Conn, targetConnectionId int) {
    writer := bufio.NewWriter(targetConnection)
    writer.WriteString("Some message\n")
    writer.Flush()
})

// remove connection from bool
connectionPool.Remove(connectionId)
```
