package squirrel

import "github.com/lann/builder"

// StatementBuilderType is the type of StatementBuilder.
type StatementBuilderType builder.Builder

// Select returns a SelectBuilder for this StatementBuilderType.
func (b StatementBuilderType) Select(columns ...string) SelectBuilder {
	_ = "STUB: not implemented"
	return *new(SelectBuilder)
}

// Insert returns a InsertBuilder for this StatementBuilderType.
func (b StatementBuilderType) Insert(into string) InsertBuilder {
	_ = "STUB: not implemented"
	return *new(InsertBuilder)
}

// Replace returns a InsertBuilder for this StatementBuilderType with the
// statement keyword set to "REPLACE".
func (b StatementBuilderType) Replace(into string) InsertBuilder {
	_ = "STUB: not implemented"
	return *new(InsertBuilder)
}

// Update returns a UpdateBuilder for this StatementBuilderType.
func (b StatementBuilderType) Update(table string) UpdateBuilder {
	_ = "STUB: not implemented"
	return *new(UpdateBuilder)
}

// Delete returns a DeleteBuilder for this StatementBuilderType.
func (b StatementBuilderType) Delete(from string) DeleteBuilder {
	_ = "STUB: not implemented"
	return *new(DeleteBuilder)
}

// PlaceholderFormat sets the PlaceholderFormat field for any child builders.
func (b StatementBuilderType) PlaceholderFormat(f PlaceholderFormat) StatementBuilderType {
	_ = "STUB: not implemented"
	return *new(StatementBuilderType)
}

// RunWith sets the RunWith field for any child builders.
func (b StatementBuilderType) RunWith(runner BaseRunner) StatementBuilderType {
	_ = "STUB: not implemented"
	return *new(StatementBuilderType)
}

// Where adds WHERE expressions to the query.
//
// See SelectBuilder.Where for more information.
func (b StatementBuilderType) Where(pred interface{}, args ...interface{}) StatementBuilderType {
	_ = "STUB: not implemented"
	return *new(StatementBuilderType)
}

// StatementBuilder is a parent builder for other builders, e.g. SelectBuilder.
var StatementBuilder = StatementBuilderType(builder.EmptyBuilder).PlaceholderFormat(Question)

// Select returns a new SelectBuilder, optionally setting some result columns.
//
// See SelectBuilder.Columns.
func Select(columns ...string) SelectBuilder { _ = "STUB: not implemented"; return *new(SelectBuilder) }

// Insert returns a new InsertBuilder with the given table name.
//
// See InsertBuilder.Into.
func Insert(into string) InsertBuilder { _ = "STUB: not implemented"; return *new(InsertBuilder) }

// Replace returns a new InsertBuilder with the statement keyword set to
// "REPLACE" and with the given table name.
//
// See InsertBuilder.Into.
func Replace(into string) InsertBuilder { _ = "STUB: not implemented"; return *new(InsertBuilder) }

// Update returns a new UpdateBuilder with the given table name.
//
// See UpdateBuilder.Table.
func Update(table string) UpdateBuilder { _ = "STUB: not implemented"; return *new(UpdateBuilder) }

// Delete returns a new DeleteBuilder with the given table name.
//
// See DeleteBuilder.Table.
func Delete(from string) DeleteBuilder { _ = "STUB: not implemented"; return *new(DeleteBuilder) }

// Case returns a new CaseBuilder
// "what" represents case value
func Case(what ...interface{}) CaseBuilder { _ = "STUB: not implemented"; return *new(CaseBuilder) }
