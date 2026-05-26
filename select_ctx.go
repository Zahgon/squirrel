//go:build go1.8
// +build go1.8

package squirrel

import (
	"context"
	"database/sql"
)

func (d *selectData) ExecContext(ctx context.Context) (sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

func (d *selectData) QueryContext(ctx context.Context) (*sql.Rows, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *selectData) QueryRowContext(ctx context.Context) RowScanner {
	_ = "STUB: not implemented"
	return *new(RowScanner)
}

// ExecContext builds and ExecContexts the query with the Runner set by RunWith.
func (b SelectBuilder) ExecContext(ctx context.Context) (sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

// QueryContext builds and QueryContexts the query with the Runner set by RunWith.
func (b SelectBuilder) QueryContext(ctx context.Context) (*sql.Rows, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// QueryRowContext builds and QueryRowContexts the query with the Runner set by RunWith.
func (b SelectBuilder) QueryRowContext(ctx context.Context) RowScanner {
	_ = "STUB: not implemented"
	return *new(RowScanner)
}

// ScanContext is a shortcut for QueryRowContext().Scan.
func (b SelectBuilder) ScanContext(ctx context.Context, dest ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}
