package squirrel

// RowScanner is the interface that wraps the Scan method.
//
// Scan behaves like database/sql.Row.Scan.
type RowScanner interface {
	Scan(...interface{}) error
}

// Row wraps database/sql.Row to let squirrel return new errors on Scan.
type Row struct {
	RowScanner
	err error
}

// Scan returns Row.err or calls RowScanner.Scan.
func (r *Row) Scan(dest ...interface{}) error { _ = "STUB: not implemented"; return nil }
