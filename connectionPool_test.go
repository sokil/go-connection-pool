package connectionPool

import (
	"testing"
	"reflect"
	"net"
)

func TestNewConnectionPool(t *testing.T) {
	pool := NewConnectionPool()
	typeName := reflect.TypeOf(pool).Name()
	if typeName != "ConnectionPool" {
		t.Errorf("Wrong type of factory method return value: \"%s\"", typeName)
	}
}
