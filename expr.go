package squirrel

const (
	// Portable true/false literals.
	sqlTrue  = "(1=1)"
	sqlFalse = "(1=0)"
)

type expr struct {
	sql  string
	args []interface{}
}

// Expr builds an expression from a SQL fragment and arguments.
//
// Ex:
//
//	Expr("FROM_UNIXTIME(?)", t)
func Expr(sql string, args ...interface{}) Sqlizer { _ = "STUB: not implemented"; return *new(Sqlizer) }

func (e expr) ToSql() (sql string, args []interface{}, err error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// no more placeholders

// escaped "??"; append it and step past

// sqlizer argument; expand it and append the result

// normal argument; append it and the placeholder

// step past the argument and placeholder

// append the remaining sql and arguments

type concatExpr []interface{}

func (ce concatExpr) ToSql() (sql string, args []interface{}, err error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// ConcatExpr builds an expression by concatenating strings and other expressions.
//
// Ex:
//
//	name_expr := Expr("CONCAT(?, ' ', ?)", firstName, lastName)
//	ConcatExpr("COALESCE(full_name,", name_expr, ")")
func ConcatExpr(parts ...interface{}) concatExpr {
	_ = "STUB: not implemented"
	return *

	// aliasExpr helps to alias part of SQL query generated with underlying "expr"
	new(concatExpr)
}

type aliasExpr struct {
	expr  Sqlizer
	alias string
}

// Alias allows to define alias for column in SelectBuilder. Useful when column is
// defined as complex expression like IF or CASE
// Ex:
//
//	.Column(Alias(caseStmt, "case_column"))
func Alias(expr Sqlizer, alias string) aliasExpr { _ = "STUB: not implemented"; return *new(aliasExpr) }

func (e aliasExpr) ToSql() (sql string, args []interface{}, err error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// Eq is syntactic sugar for use with Where/Having/Set methods.
type Eq map[string]interface{}

func (eq Eq) toSQL(useNotOpr bool) (sql string, args []interface{}, err error) {
	_ = "STUB: not implemented"

	// Empty Sql{} evaluates to true.
	return "", nil, nil
}

func (eq Eq) ToSql() (sql string, args []interface{}, err error) {
	_ = "STUB: not implemented"
	return "",

		// NotEq is syntactic sugar for use with Where/Having/Set methods.
		// Ex:
		//
		//	.Where(NotEq{"id": 1}) == "id <> 1"
		nil, nil
}

type NotEq Eq

func (neq NotEq) ToSql() (sql string, args []interface{}, err error) {
	_ = "STUB: not implemented"
	return "",

		// Like is syntactic sugar for use with LIKE conditions.
		// Ex:
		//
		//	.Where(Like{"name": "%irrel"})
		nil, nil
}

type Like map[string]interface{}

func (lk Like) toSql(opr string) (sql string, args []interface{}, err error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func (lk Like) ToSql() (sql string, args []interface{}, err error) {
	_ = "STUB: not implemented"
	return "",

		// NotLike is syntactic sugar for use with LIKE conditions.
		// Ex:
		//
		//	.Where(NotLike{"name": "%irrel"})
		nil, nil
}

type NotLike Like

func (nlk NotLike) ToSql() (sql string, args []interface{}, err error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// ILike is syntactic sugar for use with ILIKE conditions.
// Ex:
//
//	.Where(ILike{"name": "sq%"})
type ILike Like

func (ilk ILike) ToSql() (sql string, args []interface{}, err error) {
	_ = "STUB: not implemented"
	return "", nil, nil

	// NotILike is syntactic sugar for use with ILIKE conditions.
	// Ex:
	//
	//	.Where(NotILike{"name": "sq%"})
}

type NotILike Like

func (nilk NotILike) ToSql() (sql string, args []interface{}, err error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// Lt is syntactic sugar for use with Where/Having/Set methods.
// Ex:
//
//	.Where(Lt{"id": 1})
type Lt map[string]interface{}

func (lt Lt) toSql(opposite, orEq bool) (sql string, args []interface{}, err error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func (lt Lt) ToSql() (sql string, args []interface{}, err error) {
	_ = "STUB: not implemented"
	return "", nil,

		// LtOrEq is syntactic sugar for use with Where/Having/Set methods.
		// Ex:
		//
		//	.Where(LtOrEq{"id": 1}) == "id <= 1"
		nil
}

type LtOrEq Lt

func (ltOrEq LtOrEq) ToSql() (sql string, args []interface{}, err error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// Gt is syntactic sugar for use with Where/Having/Set methods.
// Ex:
//
//	.Where(Gt{"id": 1}) == "id > 1"
type Gt Lt

func (gt Gt) ToSql() (sql string, args []interface{}, err error) {
	_ = "STUB: not implemented"
	return "", nil, nil

	// GtOrEq is syntactic sugar for use with Where/Having/Set methods.
	// Ex:
	//
	//	.Where(GtOrEq{"id": 1}) == "id >= 1"
}

type GtOrEq Lt

func (gtOrEq GtOrEq) ToSql() (sql string, args []interface{}, err error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

type conj []Sqlizer

func (c conj) join(sep, defaultExpr string) (sql string, args []interface{}, err error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// And conjunction Sqlizers
type And conj

func (a And) ToSql() (string, []interface{}, error) { _ = "STUB: not implemented"; return "", nil, nil }

// Or conjunction Sqlizers
type Or conj

func (o Or) ToSql() (string, []interface{}, error) { _ = "STUB: not implemented"; return "", nil, nil }

func getSortedKeys(exp map[string]interface{}) []string { _ = "STUB: not implemented"; return nil }

func isListType(val interface{}) bool { _ = "STUB: not implemented"; return false }
