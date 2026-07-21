package sqlx

import (
	"database/sql"
	"database/sql/driver"
	"reflect"
	"strings"
	"sync"

	"github.com/jmoiron/sqlx/reflectx"
)

var NameMapper = strings.ToLower
var origMapper = reflect.ValueOf(NameMapper)

var mpr *reflectx.Mapper

var mprMu sync.Mutex

func mapper() *reflectx.Mapper { _ = "STUB: not implemented"; return nil }

func isScannable(t reflect.Type) bool { _ = "STUB: not implemented"; return false }

type ColScanner interface {
	Columns() ([]string, error)
	Scan(dest ...interface{}) error
	Err() error
}

type Queryer interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	Queryx(query string, args ...interface{}) (*Rows, error)
	QueryRowx(query string, args ...interface{}) *Row
}

type Execer interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
}

type binder interface {
	DriverName() string
	Rebind(string) string
	BindNamed(string, interface{}) (string, []interface{}, error)
}

type Ext interface {
	binder
	Queryer
	Execer
}

type Preparer interface {
	Prepare(query string) (*sql.Stmt, error)
}

func isUnsafe(i interface{}) bool { _ = "STUB: not implemented"; return false }

func mapperFor(i interface{}) *reflectx.Mapper { _ = "STUB: not implemented"; return nil }

var _scannerInterface = reflect.TypeOf((*sql.Scanner)(nil)).Elem()

//lint:ignore U1000 ignoring this for now
var _valuerInterface = reflect.TypeOf((*driver.Valuer)(nil)).Elem()

type Row struct {
	err    error
	unsafe bool
	rows   *sql.Rows
	Mapper *reflectx.Mapper
}

func (r *Row) Scan(dest ...interface{}) error { _ = "STUB: not implemented"; return nil }

func (r *Row) Columns() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *Row) ColumnTypes() ([]*sql.ColumnType, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *Row) Err() error { _ = "STUB: not implemented"; return nil }

type DB struct {
	*sql.DB
	driverName string
	unsafe     bool
	Mapper     *reflectx.Mapper
}

//lint:ignore ST1003 changing this would break the package interface.
func NewDb(db *sql.DB, driverName string) *DB { _ = "STUB: not implemented"; return nil }

func (db *DB) DriverName() string { _ = "STUB: not implemented"; return "" }

func Open(driverName, dataSourceName string) (*DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func MustOpen(driverName, dataSourceName string) *DB { _ = "STUB: not implemented"; return nil }

func (db *DB) MapperFunc(mf func(string) string) { _ = "STUB: not implemented"; return }

func (db *DB) Rebind(query string) string { _ = "STUB: not implemented"; return "" }

func (db *DB) Unsafe() *DB { _ = "STUB: not implemented"; return nil }

func (db *DB) BindNamed(query string, arg interface{}) (string, []interface{}, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func (db *DB) NamedQuery(query string, arg interface{}) (*Rows, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (db *DB) NamedExec(query string, arg interface{}) (sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

func (db *DB) Select(dest interface{}, query string, args ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (db *DB) Get(dest interface{}, query string, args ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (db *DB) MustBegin() *Tx { _ = "STUB: not implemented"; return nil }

func (db *DB) Beginx() (*Tx, error) { _ = "STUB: not implemented"; return nil, nil }

func (db *DB) Queryx(query string, args ...interface{}) (*Rows, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (db *DB) QueryRowx(query string, args ...interface{}) *Row {
	_ = "STUB: not implemented"
	return nil
}

func (db *DB) MustExec(query string, args ...interface{}) sql.Result {
	_ = "STUB: not implemented"
	return *new(sql.Result)
}

func (db *DB) Preparex(query string) (*Stmt, error) { _ = "STUB: not implemented"; return nil, nil }

func (db *DB) PrepareNamed(query string) (*NamedStmt, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Conn struct {
	*sql.Conn
	driverName string
	unsafe     bool
	Mapper     *reflectx.Mapper
}

type Tx struct {
	*sql.Tx
	driverName string
	unsafe     bool
	Mapper     *reflectx.Mapper
}

func (tx *Tx) DriverName() string { _ = "STUB: not implemented"; return "" }

func (tx *Tx) Rebind(query string) string { _ = "STUB: not implemented"; return "" }

func (tx *Tx) Unsafe() *Tx { _ = "STUB: not implemented"; return nil }

func (tx *Tx) BindNamed(query string, arg interface{}) (string, []interface{}, error) {
	_ = "STUB: not implemented"
	return "", nil, nil
}

func (tx *Tx) NamedQuery(query string, arg interface{}) (*Rows, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (tx *Tx) NamedExec(query string, arg interface{}) (sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

func (tx *Tx) Select(dest interface{}, query string, args ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (tx *Tx) Queryx(query string, args ...interface{}) (*Rows, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (tx *Tx) QueryRowx(query string, args ...interface{}) *Row {
	_ = "STUB: not implemented"
	return nil
}

func (tx *Tx) Get(dest interface{}, query string, args ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (tx *Tx) MustExec(query string, args ...interface{}) sql.Result {
	_ = "STUB: not implemented"
	return *new(sql.Result)
}

func (tx *Tx) Preparex(query string) (*Stmt, error) { _ = "STUB: not implemented"; return nil, nil }

func (tx *Tx) Stmtx(stmt interface{}) *Stmt { _ = "STUB: not implemented"; return nil }

func (tx *Tx) NamedStmt(stmt *NamedStmt) *NamedStmt { _ = "STUB: not implemented"; return nil }

func (tx *Tx) PrepareNamed(query string) (*NamedStmt, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Stmt struct {
	*sql.Stmt
	unsafe bool
	Mapper *reflectx.Mapper
}

func (s *Stmt) Unsafe() *Stmt { _ = "STUB: not implemented"; return nil }

func (s *Stmt) Select(dest interface{}, args ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Stmt) Get(dest interface{}, args ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Stmt) MustExec(args ...interface{}) sql.Result {
	_ = "STUB: not implemented"
	return *new(sql.Result)
}

func (s *Stmt) QueryRowx(args ...interface{}) *Row { _ = "STUB: not implemented"; return nil }

func (s *Stmt) Queryx(args ...interface{}) (*Rows, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type qStmt struct{ *Stmt }

func (q *qStmt) Query(query string, args ...interface{}) (*sql.Rows, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (q *qStmt) Queryx(query string, args ...interface{}) (*Rows, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (q *qStmt) QueryRowx(query string, args ...interface{}) *Row {
	_ = "STUB: not implemented"
	return nil
}

func (q *qStmt) Exec(query string, args ...interface{}) (sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

type Rows struct {
	*sql.Rows
	unsafe bool
	Mapper *reflectx.Mapper

	started bool
	fields  [][]int
	values  []interface{}
}

func (r *Rows) SliceScan() ([]interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *Rows) MapScan(dest map[string]interface{}) error { _ = "STUB: not implemented"; return nil }

func (r *Rows) StructScan(dest interface{}) error { _ = "STUB: not implemented"; return nil }

func Connect(driverName, dataSourceName string) (*DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func MustConnect(driverName, dataSourceName string) *DB { _ = "STUB: not implemented"; return nil }

func Preparex(p Preparer, query string) (*Stmt, error) { _ = "STUB: not implemented"; return nil, nil }

func Select(q Queryer, dest interface{}, query string, args ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func Get(q Queryer, dest interface{}, query string, args ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func LoadFile(e Execer, path string) (*sql.Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func MustExec(e Execer, query string, args ...interface{}) sql.Result {
	_ = "STUB: not implemented"
	return *new(sql.Result)
}

func (r *Row) SliceScan() ([]interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *Row) MapScan(dest map[string]interface{}) error { _ = "STUB: not implemented"; return nil }

func (r *Row) scanAny(dest interface{}, structOnly bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Row) StructScan(dest interface{}) error { _ = "STUB: not implemented"; return nil }

func SliceScan(r ColScanner) ([]interface{}, error) { _ = "STUB: not implemented"; return nil, nil }

func MapScan(r ColScanner, dest map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

type rowsi interface {
	Close() error
	Columns() ([]string, error)
	Err() error
	Next() bool
	Scan(...interface{}) error
}

func structOnlyError(t reflect.Type) error { _ = "STUB: not implemented"; return nil }

func scanAll(rows rowsi, dest interface{}, structOnly bool) error {
	_ = "STUB: not implemented"
	return nil
}

func StructScan(rows rowsi, dest interface{}) error { _ = "STUB: not implemented"; return nil }

func baseType(t reflect.Type, expected reflect.Kind) (reflect.Type, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Type), nil
}

func fieldsByTraversal(v reflect.Value, traversals [][]int, values []interface{}, ptrs bool) error {
	_ = "STUB: not implemented"
	return nil
}

func missingFields(transversals [][]int) (field int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}
