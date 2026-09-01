//go:build go1.8
// +build go1.8

package sqlx

import (
	"context"
	"database/sql"
)

type namedPreparerContext interface {
	PreparerContext
	binder
}

func prepareNamedContext(ctx context.Context, p namedPreparerContext, query string) (*NamedStmt, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (n *NamedStmt) ExecContext(ctx context.Context, arg interface{}) (sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}

func (n *NamedStmt) QueryContext(ctx context.Context, arg interface{}) (*sql.Rows, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (n *NamedStmt) QueryRowContext(ctx context.Context, arg interface{}) *Row {
	_ = "STUB: not implemented"
	return nil
}

func (n *NamedStmt) MustExecContext(ctx context.Context, arg interface{}) sql.Result {
	_ = "STUB: not implemented"
	return *new(sql.Result)
}

func (n *NamedStmt) QueryxContext(ctx context.Context, arg interface{}) (*Rows, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (n *NamedStmt) QueryRowxContext(ctx context.Context, arg interface{}) *Row {
	_ = "STUB: not implemented"
	return nil
}

func (n *NamedStmt) SelectContext(ctx context.Context, dest interface{}, arg interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (n *NamedStmt) GetContext(ctx context.Context, dest interface{}, arg interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func NamedQueryContext(ctx context.Context, e ExtContext, query string, arg interface{}) (*Rows, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NamedExecContext(ctx context.Context, e ExtContext, query string, arg interface{}) (sql.Result, error) {
	_ = "STUB: not implemented"
	return *new(sql.Result), nil
}
