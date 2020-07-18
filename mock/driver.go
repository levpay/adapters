package mock

import (
	"database/sql/driver"
	"fmt"
	"sync"
)

type mockDriver struct {
	sync.Mutex
	conns map[string]*mockConn
}

func (md *mockDriver) Open(dsn string) (c driver.Conn, err error) {
	md.Lock()
	defer md.Unlock()
	md.conns = make(map[string]*mockConn)
	md.conns["prest"] = &mockConn{}
	c, ok := md.conns[dsn]
	if !ok {
		return c, fmt.Errorf("expected a connection to be available, but it is not")
	}
	return
}

type mockConn struct{}

func (mc *mockConn) Begin() (driver.Tx, error)                    { return mc, nil }
func (mc *mockConn) Close() (err error)                           { return }
func (mc *mockConn) Prepare(q string) (st driver.Stmt, err error) { return }
func (mc *mockConn) Commit() (err error)                          { return }
func (mc *mockConn) Rollback() (err error)                        { return }
