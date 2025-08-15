package user

import (
	"context"
	"database/sql"
	"fmt"

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

func (m *Repository) fetch(ctx context.Context, query string, args ...interface{}) (result []model.User, err error) {
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

	result = make([]model.User, 0)
	for rows.Next() {
		t := model.User{}
		err = rows.Scan(
			&t.ID,
			&t.Name,
			&t.Phone,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}
		result = append(result, t)
	}

	return result, nil
}

func (m *Repository) Fetch(ctx context.Context) (res []model.User, err error) {
	query := squirrel.Select(`id, name, phone`).
		From(`users`)
	queryStr, args, err := query.ToSql()
	res, err = m.fetch(ctx, queryStr, args...)
	if err != nil {
		return nil, err
	}

	return
}
func (m *Repository) GetByID(ctx context.Context, id int64) (res model.User, err error) {
	query := squirrel.Select(`id, name, phone`).
		From(`users`).
		Where(squirrel.Eq{"id": id})
	queryStr, args, err := query.ToSql()

	list, err := m.fetch(ctx, queryStr, args...)
	if err != nil {
		return model.User{}, err
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, model.ErrNotFound
	}

	return
}

func (m *Repository) Store(ctx context.Context, a *model.User) (err error) {
	query := squirrel.Insert("users").
		Columns("name", "phone").
		Values(a.Name, a.Phone)
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
	query := squirrel.Delete("users").
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
func (m *Repository) Update(ctx context.Context, ar *model.User) (err error) {
	query := squirrel.Update("users").
		Set("name", ar.Name).
		Set("phone", ar.Phone).
		Where(squirrel.Eq{"user_id": ar.ID})
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
