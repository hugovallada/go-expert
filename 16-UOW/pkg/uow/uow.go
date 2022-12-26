package uow

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type RepositoryFactory func(tx *sql.Tx) any

type IUow interface {
	Register(name string, fc RepositoryFactory)
	GetRepository(ctx context.Context, name string) (any, error)
	Do(ctx context.Context, fn func(uow Uow) error) error
	CommitOrRollback() error
	Rollback() error
	Unregister(name string)
}

type Uow struct {
	Db           *sql.DB
	Tx           *sql.Tx
	Repositories map[string]RepositoryFactory
}

func NewUOW(ctx context.Context, db *sql.DB) *Uow {
	return &Uow{
		Db:           db,
		Repositories: make(map[string]RepositoryFactory),
	}
}

func (u *Uow) Register(name string, fc RepositoryFactory) {
	u.Repositories[name] = fc
}

func (u *Uow) Unregister(name string) {
	delete(u.Repositories, name)
}

func (u *Uow) Do(ctx context.Context, fn func(uow *Uow) error) error {
	if u.Tx != nil {
		return errors.New("transaction already started")
	}
	tx, err := u.Db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	u.Tx = tx
	err = fn(u)
	if err != nil {
		if errRb := u.Rollback(); errRb != nil {
			return errors.New(fmt.Sprintf("original error %s, rollback error: %s\n", err.Error(), errRb.Error()))
		}
		return err
	}
	return u.CommitOrRollback()
}

func (u *Uow) Rollback() error {
	if u.Tx == nil {
		return errors.New("transaction is not active")
	}
	err := u.Tx.Rollback()
	if err != nil {
		return err
	}
	u.Tx = nil
	return nil
}

func (u *Uow) CommitOrRollback() error {
	if u.Tx == nil {
		return errors.New("transaction is not active")
	}
	err := u.Tx.Commit()
	if err != nil {
		if errRb := u.Rollback(); errRb != nil {
			return errors.New(fmt.Sprintf("original error %s, rollback error: %s\n", err.Error(), errRb.Error()))
		}
		return err
	}
	u.Tx = nil
	return nil
}

func (u *Uow) GetRepository(ctx context.Context, name string) (any, error) {
	if u.Tx == nil {
		tx, err := u.Db.BeginTx(ctx, nil)
		if err != nil {
			return nil, err
		}
		u.Tx = tx
	}
	repo := u.Repositories[name](u.Tx)
	return repo, nil
}
