package squirrel

import (
	"database/sql"
	"sync"
)

// Prepareer is the interface that wraps the Prepare method.
//
// Prepare executes the given query as implemented by database/sql.Prepare.
type Preparer interface {
	Prepare(query string) (*sql.Stmt, error)
}

// DBProxy groups the Execer, Queryer, QueryRower, and Preparer interfaces.
type DBProxy interface {
	Execer
	Queryer
	QueryRower
	Preparer
}

// NOTE: NewStmtCache is defined in stmtcacher_ctx.go (Go >= 1.8) or stmtcacher_noctx.go (Go < 1.8).

// StmtCache wraps and delegates down to a Preparer type
//
// It also automatically prepares all statements sent to the underlying Preparer calls
// for Exec, Query and QueryRow and caches the returns *sql.Stmt using the provided
// query as the key. So that it can be automatically re-used.
type StmtCache struct {
	prep  Preparer
	cache map[string]*sql.Stmt
	mu    sync.Mutex
}

// Prepare delegates down to the underlying Preparer and caches the result
// using the provided query as a key
func (sc *StmtCache) Prepare(query string) (*sql.Stmt, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Exec delegates down to the underlying Preparer using a prepared statement
func (sc *StmtCache) Exec(query string, args ...interface{}) (res sql.Result, err error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

// Query delegates down to the underlying Preparer using a prepared statement
func (sc *StmtCache) Query(query string, args ...interface{}) (rows *sql.Rows, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// QueryRow delegates down to the underlying Preparer using a prepared statement
func (sc *StmtCache) QueryRow(query string, args ...interface{}) RowScanner {
	_ = "STUB: not implemented"
	return *new(RowScanner)
}

// Clear removes and closes all the currently cached prepared statements
func (sc *StmtCache) Clear() (err error) { _ = "STUB: not implemented"; return nil }

type DBProxyBeginner interface {
	DBProxy
	Begin() (*sql.Tx, error)
}

type stmtCacheProxy struct {
	DBProxy
	db *sql.DB
}

func NewStmtCacheProxy(db *sql.DB) DBProxyBeginner {
	_ = "STUB: not implemented"
	return *new(DBProxyBeginner)
}

func (sp *stmtCacheProxy) Begin() (*sql.Tx, error) { _ = "STUB: not implemented"; return nil, nil }
