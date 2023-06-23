package account

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"time"
	"week3_docker/internal/model"
	"week3_docker/internal/store"
)

func (r Repository) CreateAccount(ctx context.Context, account *model.Account) error {
	query := r.Store.Builder().Insert("accounts").SetMap(map[string]interface{}{
		"subdomain":      account.Subdomain,
		"auth_code":      account.AuthCode,
		"integration_id": account.IntegrationID,
	})
	sql, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("CreateAccount: %v", err)
	}

	res, err := r.Store.DB.ExecContext(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("CreateAccount: %v", err)
	}
	id, _ := res.LastInsertId()
	account.ID = store.NewNullInt64(id)
	return nil
}

func (r Repository) UpdateAccount(ctx context.Context, account *model.Account) error {
	query := r.Store.Builder().Update("accounts").
		SetMap(map[string]interface{}{
			"subdomain":      account.Subdomain,
			"auth_code":      account.AuthCode,
			"integration_id": account.IntegrationID,
			"access_token":   account.AccessToken,
			"refresh_token":  account.RefreshToken,
			"expires":        account.Expires,
			"updated_at":     time.Now(),
		}).
		Where(sq.Eq{"id": account.ID})
	sql, args, err := query.ToSql()
	if err != nil {
		return fmt.Errorf("UpdateAccount: %v", err)
	}
	_, err = r.Store.DB.ExecContext(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("UpdateAccount: %v", err)
	}
	return nil
}

func (r Repository) DeleteAccount(ctx context.Context, id uint64) error {
	query := r.Store.Builder().Delete("accounts").Where(sq.Eq{"id": id})
	_, err := query.ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("DeleteAccount: %v", err)
	}
	return nil
}

func applyListAccountFilter(q sq.SelectBuilder, filter model.ListAccountFilter) sq.SelectBuilder {
	if filter.Limit > 0 {
		q = q.Limit(uint64(filter.Limit))
	}
	if filter.Page > 0 {
		if filter.Limit < 1 {
			filter.Limit = 100
		}
		q = q.Offset(uint64((filter.Page - 1) * filter.Limit))
	}
	if filter.NeedRefresh {
		q = q.Where(sq.And{
			sq.NotEq{"access_token": nil},
			sq.LtOrEq{"expires": time.Now().Add(time.Hour * 5).Unix()},
		})
	}

	return q
}

func (r Repository) ListAccounts(ctx context.Context, filter model.ListAccountFilter) ([]model.Account, error) {
	query := r.Store.Builder().Select("*").
		From("accounts")

	query = applyListAccountFilter(query, filter)

	sqlQuery, args, err := query.ToSql()
	if err != nil {
		return nil, fmt.Errorf("ListAccounts: %v", err)
	}

	rows, err := r.Store.DB.QueryxContext(ctx, sqlQuery, args...)
	if err != nil {
		return nil, fmt.Errorf("ListAccounts: %v", err)
	}

	accounts := make([]model.Account, 0, 10)
	for rows.Next() {
		var account model.Account
		err := rows.StructScan(&account)
		if err != nil {
			return nil, fmt.Errorf("ListAccounts: %v", err)
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}

func (r Repository) GetAccount(ctx context.Context, id uint64) (model.Account, error) {
	query := r.Store.Builder().Select("*").From("accounts").Where(sq.Eq{"id": id}).Limit(1)
	sql, args, err := query.ToSql()
	if err != nil {
		return model.Account{}, fmt.Errorf("GetAccount: %v", err)
	}
	var account model.Account
	err = r.Store.DB.QueryRowxContext(ctx, sql, args...).StructScan(&account)
	if err != nil {
		return model.Account{}, fmt.Errorf("GetAccount: %v", err)
	}
	return account, nil
}
