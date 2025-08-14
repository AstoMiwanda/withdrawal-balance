package wallet

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/sirupsen/logrus"

	"withdrawal-balance/model"
)

type Repository struct {
	Conn *sql.DB
}

func NewRepository(conn *sql.DB) *Repository {
	return &Repository{conn}
}

func (m *Repository) fetch(ctx context.Context, query string, args ...interface{}) (result []model.Wallet, err error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			logrus.Error(errRow)
		}
	}()

	result = make([]model.Wallet, 0)
	for rows.Next() {
		t := model.Wallet{}
		err = rows.Scan(
			&t.ID,
			&t.UserID,
			&t.Balance,
			&t.CreatedAt,
			&t.UpdatedAt,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *Repository) Fetch(ctx context.Context) (res []model.Wallet, err error) {
	query := squirrel.Select(`*`).
		From(`wallets`)
	queryStr, args, err := query.ToSql()
	res, err = m.fetch(ctx, queryStr, args...)
	if err != nil {
		return nil, err
	}

	return
}
func (m *Repository) GetByID(ctx context.Context, id int64) (res model.Wallet, err error) {
	query := squirrel.Select(`*`).
		From(`wallets`).
		Where(squirrel.Eq{"id": id})
	queryStr, args, err := query.ToSql()

	list, err := m.fetch(ctx, queryStr, args)
	if err != nil {
		return model.Wallet{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, model.ErrNotFound
	}

	return
}

func (m *Repository) Store(ctx context.Context, a *model.Wallet) (err error) {
	query := squirrel.Insert("wallets").
		Columns("user_id", "balance", "created_at").
		Values(a.UserID, a.Balance, time.Now())
	queryStr, args, err := query.ToSql()
	if err != nil {
		return err
	}

	stmt, err := m.Conn.PrepareContext(ctx, queryStr)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		return
	}
	a.ID = lastID
	return
}

func (m *Repository) Delete(ctx context.Context, id int64) (err error) {
	query := squirrel.Delete("wallets").
		Where(squirrel.Eq{"id": id})
	queryStr, args, err := query.ToSql()
	if err != nil {
		return
	}

	stmt, err := m.Conn.PrepareContext(ctx, queryStr)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return
	}

	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return
	}

	if rowsAfected != 1 {
		err = fmt.Errorf("weird  Behavior. Total Affected: %d", rowsAfected)
		return
	}

	return
}
func (m *Repository) Update(ctx context.Context, ar *model.Wallet) (err error) {
	query := squirrel.Update("wallets").
		Set("user_id", ar.UserID).
		Set("balance", ar.Balance).
		Where(squirrel.Eq{"id": ar.ID})
	queryStr, args, err := query.ToSql()
	if err != nil {
		return
	}

	stmt, err := m.Conn.PrepareContext(ctx, queryStr)
	if err != nil {
		return
	}

	res, err := stmt.ExecContext(ctx, args...)
	if err != nil {
		return
	}
	affect, err := res.RowsAffected()
	if err != nil {
		return
	}
	if affect != 1 {
		err = fmt.Errorf("weird  Behavior. Total Affected: %d", affect)
		return
	}

	return
}
