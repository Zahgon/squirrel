package squirrel

import (
	"database/sql"

	"github.com/lann/builder"
)

type deleteData struct {
	PlaceholderFormat PlaceholderFormat
	RunWith           BaseRunner
	Prefixes          []Sqlizer
	From              string
	WhereParts        []Sqlizer
	OrderBys          []string
	Limit             string
	Offset            string
	Suffixes          []Sqlizer
}

func (d *deleteData) Exec() (sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

func (d *deleteData) ToSql() (sqlStr string, args []interface{}, err error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// Builder

// DeleteBuilder builds SQL DELETE statements.
type DeleteBuilder builder.Builder

func init() {
	builder.Register(DeleteBuilder{}, deleteData{})
}

// Format methods

// PlaceholderFormat sets PlaceholderFormat (e.g. Question or Dollar) for the
// query.
func (b DeleteBuilder) PlaceholderFormat(f PlaceholderFormat) DeleteBuilder {
	_ = "STUB: not implemented"
	return *new(DeleteBuilder)
}

// Runner methods

// RunWith sets a Runner (like database/sql.DB) to be used with e.g. Exec.
func (b DeleteBuilder) RunWith(runner BaseRunner) DeleteBuilder {
	_ = "STUB: not implemented"
	return *new(DeleteBuilder)
}

// Exec builds and Execs the query with the Runner set by RunWith.
func (b DeleteBuilder) Exec() (sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

// SQL methods

// ToSql builds the query into a SQL string and bound args.
func (b DeleteBuilder) ToSql() (string, []interface{}, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// MustSql builds the query into a SQL string and bound args.
// It panics if there are any errors.
func (b DeleteBuilder) MustSql() (string, []interface{}) { _ = "STUB: not implemented"; return "", nil }

// Prefix adds an expression to the beginning of the query
func (b DeleteBuilder) Prefix(sql string, args ...interface{}) DeleteBuilder {
	_ = "STUB: not implemented"
	return *new(DeleteBuilder)
}

// PrefixExpr adds an expression to the very beginning of the query
func (b DeleteBuilder) PrefixExpr(expr Sqlizer) DeleteBuilder {
	_ = "STUB: not implemented"
	return *new(DeleteBuilder)
}

// From sets the table to be deleted from.
func (b DeleteBuilder) From(from string) DeleteBuilder {
	_ = "STUB: not implemented"
	return *new(DeleteBuilder)
}

// Where adds WHERE expressions to the query.
//
// See SelectBuilder.Where for more information.
func (b DeleteBuilder) Where(pred interface{}, args ...interface{}) DeleteBuilder {
	_ = "STUB: not implemented"
	return *new(DeleteBuilder)
}

// OrderBy adds ORDER BY expressions to the query.
func (b DeleteBuilder) OrderBy(orderBys ...string) DeleteBuilder {
	_ = "STUB: not implemented"
	return *new(DeleteBuilder)
}

// Limit sets a LIMIT clause on the query.
func (b DeleteBuilder) Limit(limit uint64) DeleteBuilder {
	_ = "STUB: not implemented"
	return *new(DeleteBuilder)
}

// Offset sets a OFFSET clause on the query.
func (b DeleteBuilder) Offset(offset uint64) DeleteBuilder {
	_ = "STUB: not implemented"
	return *new(DeleteBuilder)
}

// Suffix adds an expression to the end of the query
func (b DeleteBuilder) Suffix(sql string, args ...interface{}) DeleteBuilder {
	_ = "STUB: not implemented"
	return *new(DeleteBuilder)
}

// SuffixExpr adds an expression to the end of the query
func (b DeleteBuilder) SuffixExpr(expr Sqlizer) DeleteBuilder {
	_ = "STUB: not implemented"
	return *new(DeleteBuilder)
}

func (b DeleteBuilder) Query() (*sql.Rows, error) { _ = "STUB: not implemented"; return nil, nil }

func (d *deleteData) Query() (*sql.Rows, error) { _ = "STUB: not implemented"; return nil, nil }
