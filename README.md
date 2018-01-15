# go-connection-pool

Connection pool for go's net.Conn

[![Go Report Card](https://goreportcard.com/badge/github.com/sokil/go-connection-pool)](https://goreportcard.com/report/github.com/sokil/go-connection-pool)
[![GoDoc](https://godoc.org/github.com/sokil/go-connection-pool?status.svg)](https://godoc.org/github.com/sokil/go-connection-pool)

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
  
  // send message to all connections in pool
  connectionPool.Range(func(targetConnection net.Conn, targetConnectionId int) {
    writer := bufio.NewWriter(targetConnection)
    writer.WriteString("Some message\n")
    writer.Flush()
  })
      
  // remove connection from bool
  connectionPool.Remove(connectionId)
```
