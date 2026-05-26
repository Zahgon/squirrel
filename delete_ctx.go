//go:build go1.8
// +build go1.8

package squirrel

import (
	"context"
	"database/sql"
)

func (d *deleteData) ExecContext(ctx context.Context) (sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

func (d *deleteData) QueryContext(ctx context.Context) (*sql.Rows, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *deleteData) QueryRowContext(ctx context.Context) RowScanner {
	_ = "STUB: not implemented"
	return *new(RowScanner)
}

// ExecContext builds and ExecContexts the query with the Runner set by RunWith.
func (b DeleteBuilder) ExecContext(ctx context.Context) (sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

// QueryContext builds and QueryContexts the query with the Runner set by RunWith.
func (b DeleteBuilder) QueryContext(ctx context.Context) (*sql.Rows, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// QueryRowContext builds and QueryRowContexts the query with the Runner set by RunWith.
func (b DeleteBuilder) QueryRowContext(ctx context.Context) RowScanner {
	_ = "STUB: not implemented"
	return *new(RowScanner)
}

// ScanContext is a shortcut for QueryRowContext().Scan.
func (b DeleteBuilder) ScanContext(ctx context.Context, dest ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}
