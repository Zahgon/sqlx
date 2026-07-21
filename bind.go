package sqlx

import (
	"reflect"
	"sync"
)

const (
	UNKNOWN = iota
	QUESTION
	DOLLAR
	NAMED
	AT
)

var defaultBinds = map[int][]string{
	DOLLAR:   {"postgres", "pgx", "pq-timeouts", "cloudsqlpostgres", "ql", "nrpostgres", "cockroach"},
	QUESTION: {"mysql", "sqlite3", "nrmysql", "nrsqlite3"},
	NAMED:    {"oci8", "ora", "goracle", "godror"},
	AT:       {"sqlserver", "azuresql"},
}

var binds sync.Map

func init() {
	for bind, drivers := range defaultBinds {
		for _, driver := range drivers {
			BindDriver(driver, bind)
		}
	}

}

func BindType(driverName string) int { _ = "STUB: not implemented"; return 0 }

func BindDriver(driverName string, bindType int) { _ = "STUB: not implemented"; return }

func Rebind(bindType int, query string) string { _ = "STUB: not implemented"; return "" }

func rebindBuff(bindType int, query string) string { _ = "STUB: not implemented"; return "" }

func asSliceForIn(i interface{}) (v reflect.Value, ok bool) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), false
}

func In(query string, args ...interface{}) (string, []interface{}, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func appendReflectSlice(args []interface{}, v reflect.Value, vlen int) []interface{} {
	_ = "STUB: not implemented"
	return nil
}
