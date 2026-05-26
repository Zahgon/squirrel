package squirrel

import (
	"database/sql"

	"github.com/lann/builder"
)

type updateData struct {
	PlaceholderFormat PlaceholderFormat
	RunWith           BaseRunner
	Prefixes          []Sqlizer
	Table             string
	SetClauses        []setClause
	From              Sqlizer
	WhereParts        []Sqlizer
	OrderBys          []string
	Limit             string
	Offset            string
	Suffixes          []Sqlizer
}

type setClause struct {
	column string
	value  interface{}
}

func (d *updateData) Exec() (sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

func (d *updateData) Query() (*sql.Rows, error) { _ = "STUB: not implemented"; return nil, nil }

func (d *updateData) QueryRow() RowScanner { _ = "STUB: not implemented"; return *new(RowScanner) }

func (d *updateData) ToSql() (sqlStr string, args []interface{}, err error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// Builder

// UpdateBuilder builds SQL UPDATE statements.
type UpdateBuilder builder.Builder

func init() {
	builder.Register(UpdateBuilder{}, updateData{})
}

// Format methods

// PlaceholderFormat sets PlaceholderFormat (e.g. Question or Dollar) for the
// query.
func (b UpdateBuilder) PlaceholderFormat(f PlaceholderFormat) UpdateBuilder {
	_ = "STUB: not implemented"
	return *new(UpdateBuilder)
}

// Runner methods

// RunWith sets a Runner (like database/sql.DB) to be used with e.g. Exec.
func (b UpdateBuilder) RunWith(runner BaseRunner) UpdateBuilder {
	_ = "STUB: not implemented"
	return *new(UpdateBuilder)
}

// Exec builds and Execs the query with the Runner set by RunWith.
func (b UpdateBuilder) Exec() (sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

func (b UpdateBuilder) Query() (*sql.Rows, error) { _ = "STUB: not implemented"; return nil, nil }

func (b UpdateBuilder) QueryRow() RowScanner { _ = "STUB: not implemented"; return *new(RowScanner) }

func (b UpdateBuilder) Scan(dest ...interface{}) error { _ = "STUB: not implemented"; return nil }

// SQL methods

// ToSql builds the query into a SQL string and bound args.
func (b UpdateBuilder) ToSql() (string, []interface{}, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// MustSql builds the query into a SQL string and bound args.
// It panics if there are any errors.
func (b UpdateBuilder) MustSql() (string, []interface{}) { _ = "STUB: not implemented"; return "", nil }

// Prefix adds an expression to the beginning of the query
func (b UpdateBuilder) Prefix(sql string, args ...interface{}) UpdateBuilder {
	_ = "STUB: not implemented"
	return *new(UpdateBuilder)
}

// PrefixExpr adds an expression to the very beginning of the query
func (b UpdateBuilder) PrefixExpr(expr Sqlizer) UpdateBuilder {
	_ = "STUB: not implemented"
	return *new(UpdateBuilder)
}

// Table sets the table to be updated.
func (b UpdateBuilder) Table(table string) UpdateBuilder {
	_ = "STUB: not implemented"
	return *new(UpdateBuilder)
}

// Set adds SET clauses to the query.
func (b UpdateBuilder) Set(column string, value interface{}) UpdateBuilder {
	_ = "STUB: not implemented"
	return *new(UpdateBuilder)
}

// SetMap is a convenience method which calls .Set for each key/value pair in clauses.
func (b UpdateBuilder) SetMap(clauses map[string]interface{}) UpdateBuilder {
	_ = "STUB: not implemented"
	return *new(UpdateBuilder)
}

// From adds FROM clause to the query
// FROM is valid construct in postgresql only.
func (b UpdateBuilder) From(from string) UpdateBuilder {
	_ = "STUB: not implemented"
	return *new(UpdateBuilder)
}

// FromSelect sets a subquery into the FROM clause of the query.
func (b UpdateBuilder) FromSelect(from SelectBuilder, alias string) UpdateBuilder {
	_ = "STUB: not implemented"
	// Prevent misnumbered parameters in nested selects (#183).
	return *new(UpdateBuilder)
}

// Where adds WHERE expressions to the query.
//
// See SelectBuilder.Where for more information.
func (b UpdateBuilder) Where(pred interface{}, args ...interface{}) UpdateBuilder {
	_ = "STUB: not implemented"
	return *new(UpdateBuilder)
}

// OrderBy adds ORDER BY expressions to the query.
func (b UpdateBuilder) OrderBy(orderBys ...string) UpdateBuilder {
	_ = "STUB: not implemented"
	return *new(UpdateBuilder)
}

// Limit sets a LIMIT clause on the query.
func (b UpdateBuilder) Limit(limit uint64) UpdateBuilder {
	_ = "STUB: not implemented"
	return *new(UpdateBuilder)
}

// Offset sets a OFFSET clause on the query.
func (b UpdateBuilder) Offset(offset uint64) UpdateBuilder {
	_ = "STUB: not implemented"
	return *new(UpdateBuilder)
}

// Suffix adds an expression to the end of the query
func (b UpdateBuilder) Suffix(sql string, args ...interface{}) UpdateBuilder {
	_ = "STUB: not implemented"
	return *new(UpdateBuilder)
}

// SuffixExpr adds an expression to the end of the query
func (b UpdateBuilder) SuffixExpr(expr Sqlizer) UpdateBuilder {
	_ = "STUB: not implemented"
	return *new(UpdateBuilder)
}
