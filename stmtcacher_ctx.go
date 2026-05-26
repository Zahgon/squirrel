//go:build go1.8
// +build go1.8

package squirrel

import (
	"context"
	"database/sql"
)

// PrepareerContext is the interface that wraps the Prepare and PrepareContext methods.
//
// Prepare executes the given query as implemented by database/sql.Prepare.
// PrepareContext executes the given query as implemented by database/sql.PrepareContext.
type PreparerContext interface {
	Preparer
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
}

// DBProxyContext groups the Execer, Queryer, QueryRower and PreparerContext interfaces.
type DBProxyContext interface {
	Execer
	Queryer
	QueryRower
	PreparerContext
}

// NewStmtCache returns a *StmtCache wrapping a PreparerContext that caches Prepared Stmts.
//
// Stmts are cached based on the string value of their queries.
func NewStmtCache(prep PreparerContext) *StmtCache { _ = "STUB: not implemented"; return nil }

// NewStmtCacher is deprecated
//
// Use NewStmtCache instead
func NewStmtCacher(prep PreparerContext) DBProxyContext {
	_ = "STUB: not implemented"
	return *new(DBProxyContext)
}

// PrepareContext delegates down to the underlying PreparerContext and caches the result
// using the provided query as a key
func (sc *StmtCache) PrepareContext(ctx context.Context, query string) (*sql.Stmt, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ExecContext delegates down to the underlying PreparerContext using a prepared statement
func (sc *StmtCache) ExecContext(ctx context.Context, query string, args ...interface{}) (res sql.Result, err error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

// QueryContext delegates down to the underlying PreparerContext using a prepared statement
func (sc *StmtCache) QueryContext(ctx context.Context, query string, args ...interface{}) (rows *sql.Rows, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// QueryRowContext delegates down to the underlying PreparerContext using a prepared statement
func (sc *StmtCache) QueryRowContext(ctx context.Context, query string, args ...interface{}) RowScanner {
	_ = "STUB: not implemented"
	return *new(RowScanner)
}
