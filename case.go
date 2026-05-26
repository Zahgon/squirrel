package squirrel

import (
	"bytes"

	"github.com/lann/builder"
)

func init() {
	builder.Register(CaseBuilder{}, caseData{})
}

// sqlizerBuffer is a helper that allows to write many Sqlizers one by one
// without constant checks for errors that may come from Sqlizer
type sqlizerBuffer struct {
	bytes.Buffer
	args []interface{}
	err  error
}

// WriteSql converts Sqlizer to SQL strings and writes it to buffer
func (b *sqlizerBuffer) WriteSql(item Sqlizer) { _ = "STUB: not implemented"; return }

func (b *sqlizerBuffer) ToSql() (string, []interface{}, error) {
	_ = "STUB: not implemented"
	return "", nil, nil

	// whenPart is a helper structure to describe SQLs "WHEN ... THEN ..." expression
}

type whenPart struct {
	when Sqlizer
	then Sqlizer
}

func newWhenPart(when interface{}, then interface{}) whenPart {
	_ = "STUB: not implemented"
	return *new(whenPart)
}

// caseData holds all the data required to build a CASE SQL construct
type caseData struct {
	What      Sqlizer
	WhenParts []whenPart
	Else      Sqlizer
}

// ToSql implements Sqlizer
func (d *caseData) ToSql() (sqlStr string, args []interface{}, err error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// CaseBuilder builds SQL CASE construct which could be used as parts of queries.
type CaseBuilder builder.Builder

// ToSql builds the query into a SQL string and bound args.
func (b CaseBuilder) ToSql() (string, []interface{}, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// MustSql builds the query into a SQL string and bound args.
// It panics if there are any errors.
func (b CaseBuilder) MustSql() (string, []interface{}) { _ = "STUB: not implemented"; return "", nil }

// what sets optional value for CASE construct "CASE [value] ..."
func (b CaseBuilder) what(expr interface{}) CaseBuilder {
	_ = "STUB: not implemented"
	return *new(CaseBuilder)
}

// When adds "WHEN ... THEN ..." part to CASE construct
func (b CaseBuilder) When(when interface{}, then interface{}) CaseBuilder {
	_ = "STUB: not implemented"
	// TODO: performance hint: replace slice of WhenPart with just slice of parts
	// where even indices of the slice belong to "when"s and odd indices belong to "then"s
	return *new(CaseBuilder)
}

// What sets optional "ELSE ..." part for CASE construct
func (b CaseBuilder) Else(expr interface{}) CaseBuilder {
	_ = "STUB: not implemented"
	return *new(CaseBuilder)
}
