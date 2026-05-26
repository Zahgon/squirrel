//go:build go1.8
// +build go1.8

package squirrel

import (
	"context"
	"database/sql"
	"errors"
)

// NoContextSupport is returned if a db doesn't support Context.
var NoContextSupport = errors.New("DB does not support Context")

// ExecerContext is the interface that wraps the ExecContext method.
//
// Exec executes the given query as implemented by database/sql.ExecContext.
type ExecerContext interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

// QueryerContext is the interface that wraps the QueryContext method.
//
// QueryContext executes the given query as implemented by database/sql.QueryContext.
type QueryerContext interface {
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
}

// QueryRowerContext is the interface that wraps the QueryRowContext method.
//
// QueryRowContext executes the given query as implemented by database/sql.QueryRowContext.
type QueryRowerContext interface {
	QueryRowContext(ctx context.Context, query string, args ...interface{}) RowScanner
}

// RunnerContext groups the Runner interface, along with the Context versions of each of
// its methods
type RunnerContext interface {
	Runner
	QueryerContext
	QueryRowerContext
	ExecerContext
}

// WrapStdSqlCtx wraps a type implementing the standard SQL interface plus the context
// versions of the methods with methods that squirrel expects.
func WrapStdSqlCtx(stdSqlCtx StdSqlCtx) RunnerContext {
	_ = "STUB: not implemented"
	return *new(RunnerContext)
}

// StdSqlCtx encompasses the standard methods of the *sql.DB type, along with the Context
// versions of those methods, and other types that wrap these methods.
type StdSqlCtx interface {
	StdSql
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
}

type stdsqlCtxRunner struct {
	StdSqlCtx
}

func (r *stdsqlCtxRunner) QueryRow(query string, args ...interface{}) RowScanner {
	_ = "STUB: not implemented"
	return *new(RowScanner)
}

func (r *stdsqlCtxRunner) QueryRowContext(ctx context.Context, query string, args ...interface{}) RowScanner {
	_ = "STUB: not implemented"
	return *new(RowScanner)
}

// ExecContextWith ExecContexts the SQL returned by s with db.
func ExecContextWith(ctx context.Context, db ExecerContext, s Sqlizer) (res sql.Result, err error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

// QueryContextWith QueryContexts the SQL returned by s with db.
func QueryContextWith(ctx context.Context, db QueryerContext, s Sqlizer) (rows *sql.Rows, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// QueryRowContextWith QueryRowContexts the SQL returned by s with db.
func QueryRowContextWith(ctx context.Context, db QueryRowerContext, s Sqlizer) RowScanner {
	_ = "STUB: not implemented"
	return *new(RowScanner)
}
