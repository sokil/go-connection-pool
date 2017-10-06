package connectionPool

import (
	"net"
	"sync"
)

// ConnectionPool is a thread safe list of net.Conn instances
type ConnectionPool struct {
	mutex sync.RWMutex
	list  map[int]net.Conn
}

// NewConnectionPool is the factory method to create new connection pool
func NewConnectionPool() *ConnectionPool {
	pool := &ConnectionPool{
		list: make(map[int]net.Conn),
	}

	return pool
}

// Add collection to pool
func (pool *ConnectionPool) Add(connection net.Conn) int {
	pool.mutex.Lock()
	nextConnectionId := len(pool.list)
	pool.list[nextConnectionId] = connection
	pool.mutex.Unlock()
	return nextConnectionId
}

// Get connection by id
func (pool *ConnectionPool) Get(connectionId int) net.Conn {
	pool.mutex.RLock()
	connection := pool.list[connectionId]
	pool.mutex.RUnlock()
	return connection
}

// Remove connection from pool
func (pool *ConnectionPool) Remove(connectionId int) {
	pool.mutex.Lock()
	delete(pool.list, connectionId)
	pool.mutex.Unlock()
}

// Size of connections pool
func (pool *ConnectionPool) Size() int {
	return len(pool.list)
}

// Range iterates over pool
func (pool *ConnectionPool) Range(callback func(net.Conn, int)) {
	pool.mutex.RLock()
	for connectionId, connection := range pool.list {
		callback(connection, connectionId)
	}
	pool.mutex.RUnlock()
}
