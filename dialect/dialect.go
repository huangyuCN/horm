package dialect

import "reflect"

type Dialect interface {
	DataTypeOf(typ reflect.Value) string
	TableExistSQL(table string) (string, []interface{})
}

var DialectMap = map[string]Dialect{}

func RegisterDialect(name string, dialect Dialect) {
	DialectMap[name] = dialect
}

func GetDialect(name string) (dialect Dialect, ok bool) {
	dialect, ok = DialectMap[name]
	return
}
