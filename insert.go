package squirrel

import (
	"database/sql"
	"io"

	"github.com/lann/builder"
)

type insertData struct {
	PlaceholderFormat PlaceholderFormat
	RunWith           BaseRunner
	Prefixes          []Sqlizer
	StatementKeyword  string
	Options           []string
	Into              string
	Columns           []string
	Values            [][]interface{}
	Suffixes          []Sqlizer
	Select            *SelectBuilder
}

func (d *insertData) Exec() (sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

func (d *insertData) Query() (*sql.Rows, error) { _ = "STUB: not implemented"; return nil, nil }

func (d *insertData) QueryRow() RowScanner { _ = "STUB: not implemented"; return *new(RowScanner) }

func (d *insertData) ToSql() (sqlStr string, args []interface{}, err error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func (d *insertData) appendValuesToSQL(w io.Writer, args []interface{}) ([]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (d *insertData) appendSelectToSQL(w io.Writer, args []interface{}) ([]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Builder

// InsertBuilder builds SQL INSERT statements.
type InsertBuilder builder.Builder

func init() {
	builder.Register(InsertBuilder{}, insertData{})
}

// Format methods

// PlaceholderFormat sets PlaceholderFormat (e.g. Question or Dollar) for the
// query.
func (b InsertBuilder) PlaceholderFormat(f PlaceholderFormat) InsertBuilder {
	_ = "STUB: not implemented"
	return *new(InsertBuilder)
}

// Runner methods

// RunWith sets a Runner (like database/sql.DB) to be used with e.g. Exec.
func (b InsertBuilder) RunWith(runner BaseRunner) InsertBuilder {
	_ = "STUB: not implemented"
	return *new(InsertBuilder)
}

// Exec builds and Execs the query with the Runner set by RunWith.
func (b InsertBuilder) Exec() (sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

// Query builds and Querys the query with the Runner set by RunWith.
func (b InsertBuilder) Query() (*sql.Rows, error) { _ = "STUB: not implemented"; return nil, nil }

// QueryRow builds and QueryRows the query with the Runner set by RunWith.
func (b InsertBuilder) QueryRow() RowScanner { _ = "STUB: not implemented"; return *new(RowScanner) }

// Scan is a shortcut for QueryRow().Scan.
func (b InsertBuilder) Scan(dest ...interface{}) error { _ = "STUB: not implemented"; return nil }

// SQL methods

// ToSql builds the query into a SQL string and bound args.
func (b InsertBuilder) ToSql() (string, []interface{}, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// MustSql builds the query into a SQL string and bound args.
// It panics if there are any errors.
func (b InsertBuilder) MustSql() (string, []interface{}) { _ = "STUB: not implemented"; return "", nil }

// Prefix adds an expression to the beginning of the query
func (b InsertBuilder) Prefix(sql string, args ...interface{}) InsertBuilder {
	_ = "STUB: not implemented"
	return *new(InsertBuilder)
}

// PrefixExpr adds an expression to the very beginning of the query
func (b InsertBuilder) PrefixExpr(expr Sqlizer) InsertBuilder {
	_ = "STUB: not implemented"
	return *new(InsertBuilder)
}

// Options adds keyword options before the INTO clause of the query.
func (b InsertBuilder) Options(options ...string) InsertBuilder {
	_ = "STUB: not implemented"
	return *new(InsertBuilder)
}

// Into sets the INTO clause of the query.
func (b InsertBuilder) Into(into string) InsertBuilder {
	_ = "STUB: not implemented"
	return *new(InsertBuilder)
}

// Columns adds insert columns to the query.
func (b InsertBuilder) Columns(columns ...string) InsertBuilder {
	_ = "STUB: not implemented"
	return *new(InsertBuilder)
}

// Values adds a single row's values to the query.
func (b InsertBuilder) Values(values ...interface{}) InsertBuilder {
	_ = "STUB: not implemented"
	return *new(InsertBuilder)
}

// Suffix adds an expression to the end of the query
func (b InsertBuilder) Suffix(sql string, args ...interface{}) InsertBuilder {
	_ = "STUB: not implemented"
	return *new(InsertBuilder)
}

// SuffixExpr adds an expression to the end of the query
func (b InsertBuilder) SuffixExpr(expr Sqlizer) InsertBuilder {
	_ = "STUB: not implemented"
	return *new(InsertBuilder)
}

// SetMap set columns and values for insert builder from a map of column name and value
// note that it will reset all previous columns and values was set if any
func (b InsertBuilder) SetMap(clauses map[string]interface{}) InsertBuilder {
	_ = "STUB: not implemented"
	// Keep the columns in a consistent order by sorting the column key string.
	return *new(InsertBuilder)
}

// Select set Select clause for insert query
// If Values and Select are used, then Select has higher priority
func (b InsertBuilder) Select(sb SelectBuilder) InsertBuilder {
	_ = "STUB: not implemented"
	return *new(InsertBuilder)
}

func (b InsertBuilder) statementKeyword(keyword string) InsertBuilder {
	_ = "STUB: not implemented"
	return *new(InsertBuilder)
}
