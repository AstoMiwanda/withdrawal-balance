package withdrawalhistory

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

func (m *Repository) fetch(ctx context.Context, query string, args ...interface{}) (result []model.WithdrawalHistory, err error) {
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

	result = make([]model.WithdrawalHistory, 0)
	for rows.Next() {
		t := model.WithdrawalHistory{}
		err = rows.Scan(
			&t.ID,
			&t.UserID,
			&t.WalletID,
			&t.Amount,
			&t.BankAccountNumber,
			&t.BankName,
			&t.Status,
			&t.TransactionReference,
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

func (m *Repository) Fetch(ctx context.Context) (res []model.WithdrawalHistory, err error) {
	query := squirrel.Select(`*`).
		From(`withdrawal_histories`)
	queryStr, args, err := query.ToSql()
	res, err = m.fetch(ctx, queryStr, args...)
	if err != nil {
		return nil, err
	}

	return
}
func (m *Repository) GetByID(ctx context.Context, id int64) (res model.WithdrawalHistory, err error) {
	query := squirrel.Select(`*`).
		From(`withdrawal_histories`).
		Where(squirrel.Eq{"id": id})
	queryStr, args, err := query.ToSql()

	list, err := m.fetch(ctx, queryStr, args)
	if err != nil {
		return model.WithdrawalHistory{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, model.ErrNotFound
	}

	return
}

func (m *Repository) Store(ctx context.Context, data *model.WithdrawalHistory) (err error) {
	query := squirrel.Insert("withdrawal_histories").
		Columns("user_id", "wallet_id", "amount", "bank_account_number", "bank_name", "status", "transaction_reference", "created_at", "updated_at").
		Values(data.UserID, data.WalletID, data.Amount, data.BankAccountNumber, data.BankName, data.Status, data.TransactionReference, time.Now(), time.Now())
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
	data.ID = lastID
	return
}

func (m *Repository) Delete(ctx context.Context, id int64) (err error) {
	query := squirrel.Delete("withdrawal_histories").
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
func (m *Repository) Update(ctx context.Context, data *model.WithdrawalHistory) (err error) {
	query := squirrel.Update("withdrawal_histories").
		Set("user_id", data.UserID).
		Set("wallet_id", data.WalletID).
		Set("amount", data.Amount).
		Set("bank_account_number", data.BankAccountNumber).
		Set("bank_name", data.BankName).
		Set("status", data.Status).
		Set("transaction_reference", data.TransactionReference).
		Set("updated_at", time.Now()).
		Where(squirrel.Eq{"id": data.ID})
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
