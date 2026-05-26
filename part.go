package squirrel

import (
	"io"
)

type part struct {
	pred interface{}
	args []interface{}
}

func newPart(pred interface{}, args ...interface{}) Sqlizer {
	_ = "STUB: not implemented"
	return *new(Sqlizer)
}

func (p part) ToSql() (sql string, args []interface{}, err error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

// no-op

func nestedToSql(s Sqlizer) (string, []interface{}, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func appendToSql(parts []Sqlizer, w io.Writer, sep string, args []interface{}) ([]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
