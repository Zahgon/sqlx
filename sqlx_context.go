//go:build go1.8
// +build go1.8

package sqlx

import (
	"context"
	"database/sql"
)

func ConnectContext(ctx context.Context, driverName, dataSourceName string) (*DB, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type QueryerContext interface {
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryxContext(ctx context.Context, query string, args ...interface{}) (*Rows, error)
	QueryRowxContext(ctx context.Context, query string, args ...interface{}) *Row
}

type PreparerContext interface {
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
}

type ExecerContext interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
}

type ExtContext interface {
	binder
	QueryerContext
	ExecerContext
}

func SelectContext(ctx context.Context, q QueryerContext, dest interface{}, query string, args ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func PreparexContext(ctx context.Context, p PreparerContext, query string) (*Stmt, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func GetContext(ctx context.Context, q QueryerContext, dest interface{}, query string, args ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func LoadFileContext(ctx context.Context, e ExecerContext, path string) (*sql.Result, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func MustExecContext(ctx context.Context, e ExecerContext, query string, args ...interface{}) sql.Result {
	_ = "STUB: not implemented"
	return *new(sql.Result)
}

func (db *DB) PrepareNamedContext(ctx context.Context, query string) (*NamedStmt, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (db *DB) NamedQueryContext(ctx context.Context, query string, arg interface{}) (*Rows, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (db *DB) NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

func (db *DB) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (db *DB) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (db *DB) PreparexContext(ctx context.Context, query string) (*Stmt, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (db *DB) QueryxContext(ctx context.Context, query string, args ...interface{}) (*Rows, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (db *DB) QueryRowxContext(ctx context.Context, query string, args ...interface{}) *Row {
	_ = "STUB: not implemented"
	return nil
}

func (db *DB) MustBeginTx(ctx context.Context, opts *sql.TxOptions) *Tx {
	_ = "STUB: not implemented"
	return nil
}

func (db *DB) MustExecContext(ctx context.Context, query string, args ...interface{}) sql.Result {
	_ = "STUB: not implemented"
	return *new(sql.Result)
}

func (db *DB) BeginTxx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (db *DB) Connx(ctx context.Context) (*Conn, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *Conn) BeginTxx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Conn) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Conn) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Conn) PreparexContext(ctx context.Context, query string) (*Stmt, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Conn) QueryxContext(ctx context.Context, query string, args ...interface{}) (*Rows, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Conn) QueryRowxContext(ctx context.Context, query string, args ...interface{}) *Row {
	_ = "STUB: not implemented"
	return nil
}

func (c *Conn) Rebind(query string) string { _ = "STUB: not implemented"; return "" }

func (tx *Tx) StmtxContext(ctx context.Context, stmt interface{}) *Stmt {
	_ = "STUB: not implemented"
	return nil
}

func (tx *Tx) NamedStmtContext(ctx context.Context, stmt *NamedStmt) *NamedStmt {
	_ = "STUB: not implemented"
	return nil
}

func (tx *Tx) PreparexContext(ctx context.Context, query string) (*Stmt, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (tx *Tx) PrepareNamedContext(ctx context.Context, query string) (*NamedStmt, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (tx *Tx) MustExecContext(ctx context.Context, query string, args ...interface{}) sql.Result {
	_ = "STUB: not implemented"
	return *new(sql.Result)
}

func (tx *Tx) QueryxContext(ctx context.Context, query string, args ...interface{}) (*Rows, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (tx *Tx) SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (tx *Tx) GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (tx *Tx) QueryRowxContext(ctx context.Context, query string, args ...interface{}) *Row {
	_ = "STUB: not implemented"
	return nil
}

func (tx *Tx) NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

func (s *Stmt) SelectContext(ctx context.Context, dest interface{}, args ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Stmt) GetContext(ctx context.Context, dest interface{}, args ...interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *Stmt) MustExecContext(ctx context.Context, args ...interface{}) sql.Result {
	_ = "STUB: not implemented"
	return *new(sql.Result)
}

func (s *Stmt) QueryRowxContext(ctx context.Context, args ...interface{}) *Row {
	_ = "STUB: not implemented"
	return nil
}

func (s *Stmt) QueryxContext(ctx context.Context, args ...interface{}) (*Rows, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (q *qStmt) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (q *qStmt) QueryxContext(ctx context.Context, query string, args ...interface{}) (*Rows, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (q *qStmt) QueryRowxContext(ctx context.Context, query string, args ...interface{}) *Row {
	_ = "STUB: not implemented"
	return nil
}

func (q *qStmt) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}
