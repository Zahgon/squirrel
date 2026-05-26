package squirrel

type wherePart part

func newWherePart(pred interface{}, args ...interface{}) Sqlizer {
	_ = "STUB: not implemented"
	return *new(Sqlizer)
}

func (p wherePart) ToSql() (sql string, args []interface{}, err error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// no-op
