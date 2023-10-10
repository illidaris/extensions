package dependency

import (
	"context"
)

type DbAction func(ctx context.Context) error

type IUnitOfWork interface {
	Execute(ctx context.Context, fs ...DbAction) (e error)
}

type IRepository[T IPo] interface {
	BaseCreate(ctx context.Context, P T) (int64, error)
	BaseSave(ctx context.Context, p interface{}) (int64, error)
	BaseBatchCreate(ctx context.Context, p interface{}, ps interface{}) (int64, error)
	BaseUpdate(ctx context.Context, p interface{}, conds ...interface{}) (int64, error)
	BaseUpdates(ctx context.Context, p interface{}, ups interface{}, conds ...interface{}) (int64, error)
	UpdateByCond(ctx context.Context, p interface{}, conds ...interface{}) (int64, error)
	BaseDelete(ctx context.Context, p interface{}, conds ...interface{}) (int64, error)
	BaseGet(ctx context.Context, p interface{}, conds ...interface{}) (int64, error)
	BaseCount(ctx context.Context, p interface{}, conds ...interface{}) (int64, error)
	BaseAny(ctx context.Context, result interface{}, ids ...interface{}) (int64, error)
}
