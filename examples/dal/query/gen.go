// Code generated by github.com/oo-pp307/gen. DO NOT EDIT.
// Code generated by github.com/oo-pp307/gen. DO NOT EDIT.
// Code generated by github.com/oo-pp307/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"
)

var (
	Q       = new(Query)
	Mytable *mytable
)

func SetDefault(db *gorm.DB) {
	*Q = *Use(db)
	Mytable = &Q.Mytable
}

func Use(db *gorm.DB) *Query {
	return &Query{
		db:      db,
		Mytable: newMytable(db),
	}
}

type Query struct {
	db *gorm.DB

	Mytable mytable
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:      db,
		Mytable: q.Mytable.clone(db),
	}
}

type queryCtx struct {
	Mytable mytableDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Mytable: *q.Mytable.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	return &QueryTx{q.clone(q.db.Begin(opts...))}
}

type QueryTx struct{ *Query }

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
