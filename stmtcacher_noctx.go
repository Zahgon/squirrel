//go:build !go1.8
// +build !go1.8

package squirrel

// NewStmtCacher returns a DBProxy wrapping prep that caches Prepared Stmts.
//
// Stmts are cached based on the string value of their queries.
func NewStmtCache(prep Preparer) *StmtCache { _ = "STUB: not implemented"; return nil }

// NewStmtCacher is deprecated
//
// Use NewStmtCache instead
func NewStmtCacher(prep Preparer) DBProxy { _ = "STUB: not implemented"; return *new(DBProxy) }
