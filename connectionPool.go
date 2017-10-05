package connectionPool

import (
	"net"
	"sync"
)

// connection pool
type ConnectionPool struct {
	mutex sync.RWMutex
	list  map[int]net.Conn
}

// factory method to get new connection pool
func NewConnectionPool() ConnectionPool {
	pool := ConnectionPool{
		list: make(map[int]net.Conn),
	}

	return pool
}

// add collection to pool
func (pool *ConnectionPool) Add(connection net.Conn) int {
	pool.mutex.Lock()
	nextConnectionId := len(pool.list)
	pool.list[nextConnectionId] = connection
	pool.mutex.Unlock()
	return nextConnectionId
}

// get connection by id
func (pool *ConnectionPool) Get(connectionId int) net.Conn {
	pool.mutex.RLock()
	connection := pool.list[connectionId]
	pool.mutex.RUnlock()
	return connection
}

// remove connection from pool
func (pool *ConnectionPool) Remove(connectionId int) {
	pool.mutex.Lock()
	delete(pool.list, connectionId)
	pool.mutex.Unlock()
}

// get size of connections pool
func (pool *ConnectionPool) Size() int {
	return len(pool.list)
}

// iterator
func (pool *ConnectionPool) Range(callback func(net.Conn, int)) {
	pool.mutex.RLock()
	for connectionId, connection := range pool.list {
		callback(connection, connectionId)
	}
	pool.mutex.RUnlock()
}
