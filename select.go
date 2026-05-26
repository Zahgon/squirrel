package squirrel

import (
	"database/sql"

	"github.com/lann/builder"
)

type selectData struct {
	PlaceholderFormat PlaceholderFormat
	RunWith           BaseRunner
	Prefixes          []Sqlizer
	Options           []string
	Columns           []Sqlizer
	From              Sqlizer
	Joins             []Sqlizer
	WhereParts        []Sqlizer
	GroupBys          []string
	HavingParts       []Sqlizer
	OrderByParts      []Sqlizer
	Limit             string
	Offset            string
	Suffixes          []Sqlizer
}

func (d *selectData) Exec() (sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

func (d *selectData) Query() (*sql.Rows, error) { _ = "STUB: not implemented"; return nil, nil }

func (d *selectData) QueryRow() RowScanner { _ = "STUB: not implemented"; return *new(RowScanner) }

func (d *selectData) ToSql() (sqlStr string, args []interface{}, err error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func (d *selectData) toSqlRaw() (sqlStr string, args []interface{}, err error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// Builder

// SelectBuilder builds SQL SELECT statements.
type SelectBuilder builder.Builder

func init() {
	builder.Register(SelectBuilder{}, selectData{})
}

// Format methods

// PlaceholderFormat sets PlaceholderFormat (e.g. Question or Dollar) for the
// query.
func (b SelectBuilder) PlaceholderFormat(f PlaceholderFormat) SelectBuilder {
	_ = "STUB: not implemented"
	return *new(SelectBuilder)
}

// Runner methods

// RunWith sets a Runner (like database/sql.DB) to be used with e.g. Exec.
// For most cases runner will be a database connection.
//
// Internally we use this to mock out the database connection for testing.
func (b SelectBuilder) RunWith(runner BaseRunner) SelectBuilder {
	_ = "STUB: not implemented"
	return *new(SelectBuilder)
}

// Exec builds and Execs the query with the Runner set by RunWith.
func (b SelectBuilder) Exec() (sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

// Query builds and Querys the query with the Runner set by RunWith.
func (b SelectBuilder) Query() (*sql.Rows, error) { _ = "STUB: not implemented"; return nil, nil }

// QueryRow builds and QueryRows the query with the Runner set by RunWith.
func (b SelectBuilder) QueryRow() RowScanner { _ = "STUB: not implemented"; return *new(RowScanner) }

// Scan is a shortcut for QueryRow().Scan.
func (b SelectBuilder) Scan(dest ...interface{}) error { _ = "STUB: not implemented"; return nil }

// SQL methods

// ToSql builds the query into a SQL string and bound args.
func (b SelectBuilder) ToSql() (string, []interface{}, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func (b SelectBuilder) toSqlRaw() (string, []interface{}, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// MustSql builds the query into a SQL string and bound args.
// It panics if there are any errors.
func (b SelectBuilder) MustSql() (string, []interface{}) { _ = "STUB: not implemented"; return "", nil }

// Prefix adds an expression to the beginning of the query
func (b SelectBuilder) Prefix(sql string, args ...interface{}) SelectBuilder {
	_ = "STUB: not implemented"
	return *new(SelectBuilder)
}

// PrefixExpr adds an expression to the very beginning of the query
func (b SelectBuilder) PrefixExpr(expr Sqlizer) SelectBuilder {
	_ = "STUB: not implemented"
	return *new(SelectBuilder)
}

// Distinct adds a DISTINCT clause to the query.
func (b SelectBuilder) Distinct() SelectBuilder {
	_ = "STUB: not implemented"
	return *new(SelectBuilder)
}

// Options adds select option to the query
func (b SelectBuilder) Options(options ...string) SelectBuilder {
	_ = "STUB: not implemented"
	return *new(SelectBuilder)
}

// Columns adds result columns to the query.
func (b SelectBuilder) Columns(columns ...string) SelectBuilder {
	_ = "STUB: not implemented"
	return *new(SelectBuilder)
}

// RemoveColumns remove all columns from query.
// Must add a new column with Column or Columns methods, otherwise
// return a error.
func (b SelectBuilder) RemoveColumns() SelectBuilder {
	_ = "STUB: not implemented"
	return *new(SelectBuilder)
}

// Column adds a result column to the query.
// Unlike Columns, Column accepts args which will be bound to placeholders in
// the columns string, for example:
//
//	Column("IF(col IN ("+squirrel.Placeholders(3)+"), 1, 0) as col", 1, 2, 3)
func (b SelectBuilder) Column(column interface{}, args ...interface{}) SelectBuilder {
	_ = "STUB: not implemented"
	return *new(SelectBuilder)
}

// From sets the FROM clause of the query.
func (b SelectBuilder) From(from string) SelectBuilder {
	_ = "STUB: not implemented"
	return *new(SelectBuilder)
}

// FromSelect sets a subquery into the FROM clause of the query.
func (b SelectBuilder) FromSelect(from SelectBuilder, alias string) SelectBuilder {
	_ = "STUB: not implemented"
	// Prevent misnumbered parameters in nested selects (#183).
	return *new(SelectBuilder)
}

// JoinClause adds a join clause to the query.
func (b SelectBuilder) JoinClause(pred interface{}, args ...interface{}) SelectBuilder {
	_ = "STUB: not implemented"
	return *new(SelectBuilder)
}

// Join adds a JOIN clause to the query.
func (b SelectBuilder) Join(join string, rest ...interface{}) SelectBuilder {
	_ = "STUB: not implemented"
	return *new(SelectBuilder)
}

// LeftJoin adds a LEFT JOIN clause to the query.
func (b SelectBuilder) LeftJoin(join string, rest ...interface{}) SelectBuilder {
	_ = "STUB: not implemented"
	return *new(SelectBuilder)
}

// RightJoin adds a RIGHT JOIN clause to the query.
func (b SelectBuilder) RightJoin(join string, rest ...interface{}) SelectBuilder {
	_ = "STUB: not implemented"
	return *new(SelectBuilder)
}

// InnerJoin adds a INNER JOIN clause to the query.
func (b SelectBuilder) InnerJoin(join string, rest ...interface{}) SelectBuilder {
	_ = "STUB: not implemented"
	return *new(SelectBuilder)
}

// CrossJoin adds a CROSS JOIN clause to the query.
func (b SelectBuilder) CrossJoin(join string, rest ...interface{}) SelectBuilder {
	_ = "STUB: not implemented"
	return *new(SelectBuilder)
}

// Where adds an expression to the WHERE clause of the query.
//
// Expressions are ANDed together in the generated SQL.
//
// Where accepts several types for its pred argument:
//
// nil OR "" - ignored.
//
// string - SQL expression.
// If the expression has SQL placeholders then a set of arguments must be passed
// as well, one for each placeholder.
//
// map[string]interface{} OR Eq - map of SQL expressions to values. Each key is
// transformed into an expression like "<key> = ?", with the corresponding value
// bound to the placeholder. If the value is nil, the expression will be "<key>
// IS NULL". If the value is an array or slice, the expression will be "<key> IN
// (?,?,...)", with one placeholder for each item in the value. These expressions
// are ANDed together.
//
// Where will panic if pred isn't any of the above types.
func (b SelectBuilder) Where(pred interface{}, args ...interface{}) SelectBuilder {
	_ = "STUB: not implemented"
	return *new(SelectBuilder)
}

// GroupBy adds GROUP BY expressions to the query.
func (b SelectBuilder) GroupBy(groupBys ...string) SelectBuilder {
	_ = "STUB: not implemented"
	return *new(SelectBuilder)
}

// Having adds an expression to the HAVING clause of the query.
//
// See Where.
func (b SelectBuilder) Having(pred interface{}, rest ...interface{}) SelectBuilder {
	_ = "STUB: not implemented"
	return *new(SelectBuilder)
}

// OrderByClause adds ORDER BY clause to the query.
func (b SelectBuilder) OrderByClause(pred interface{}, args ...interface{}) SelectBuilder {
	_ = "STUB: not implemented"
	return *new(SelectBuilder)
}

// OrderBy adds ORDER BY expressions to the query.
func (b SelectBuilder) OrderBy(orderBys ...string) SelectBuilder {
	_ = "STUB: not implemented"
	return *new(SelectBuilder)
}

// Limit sets a LIMIT clause on the query.
func (b SelectBuilder) Limit(limit uint64) SelectBuilder {
	_ = "STUB: not implemented"
	return *new(SelectBuilder)
}

// Limit ALL allows to access all records with limit
func (b SelectBuilder) RemoveLimit() SelectBuilder {
	_ = "STUB: not implemented"
	return *new(SelectBuilder)
}

// Offset sets a OFFSET clause on the query.
func (b SelectBuilder) Offset(offset uint64) SelectBuilder {
	_ = "STUB: not implemented"
	return *new(SelectBuilder)
}

// RemoveOffset removes OFFSET clause.
func (b SelectBuilder) RemoveOffset() SelectBuilder {
	_ = "STUB: not implemented"
	return *new(SelectBuilder)
}

// Suffix adds an expression to the end of the query
func (b SelectBuilder) Suffix(sql string, args ...interface{}) SelectBuilder {
	_ = "STUB: not implemented"
	return *new(SelectBuilder)
}

// SuffixExpr adds an expression to the end of the query
func (b SelectBuilder) SuffixExpr(expr Sqlizer) SelectBuilder {
	_ = "STUB: not implemented"
	return *new(SelectBuilder)
}
