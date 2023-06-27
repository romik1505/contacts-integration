package account

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
	"week3_docker/internal/model"
)

func (r Repository) CreateAccount(ctx context.Context, account *model.Account) error {
	res := r.Store.DB.WithContext(ctx).Create(account)
	if res.Error != nil {
		return fmt.Errorf("CreateAccount: %v", res.Error)
	}
	return nil
}

func (r Repository) UpdateAccount(ctx context.Context, account *model.Account) error {
	tx := r.Store.DB.WithContext(ctx).Updates(*account)

	if tx.Error != nil {
		return fmt.Errorf("UpdateAccount: %v", tx.Error)
	}
	return nil
}

func (r Repository) DeleteAccount(ctx context.Context, id uint64) error {
	res := r.Store.DB.WithContext(ctx).
		Where("id = ?", id).
		Delete(&model.Account{})

	if res.Error != nil {
		return fmt.Errorf("DeleteAccount: %v", res.Error)
	}
	return nil
}

func applyListAccountFilter(q *gorm.DB, filter model.ListAccountFilter) *gorm.DB {
	if filter.JoinIntegrations {
		q = q.Joins("LEFT JOIN integrations ON integrations.account_id=accounts.id ").Preload("Integrations")
	}

	if filter.Limit > 0 {
		q = q.Limit(filter.Limit)
	}
	if filter.Page > 0 {
		if filter.Limit < 1 {
			filter.Limit = 100
		}
		q = q.Offset((filter.Page - 1) * filter.Limit)
	}
	if filter.NeedRefresh {
		q = q.Where("access_token IS NOT NULL AND expires <= ?", time.Now().Add(time.Hour*5).Unix()).Order("expires")
	}

	if filter.AmoAuthorized != nil {
		if *filter.AmoAuthorized {
			q = q.Where("access_token IS NOT NULL AND expires > ?", time.Now().Unix()).Order("id")
		} else {
			q = q.Where("access_token IS NULL OR expires <= ?", time.Now().Unix())
		}
	}

	return q
}

func (r Repository) ListAccounts(ctx context.Context, filter model.ListAccountFilter) ([]model.Account, error) {
	q := r.Store.DB.WithContext(ctx)
	q = applyListAccountFilter(q, filter)

	var accounts []model.Account

	res := q.Find(&accounts)

	if res.Error != nil {
		return nil, fmt.Errorf("ListAccounts: %v", res.Error)
	}
	return accounts, nil
}

func (r Repository) GetAccount(ctx context.Context, id uint64) (model.Account, error) {
	var account model.Account
	res := r.Store.DB.WithContext(ctx).First(&account, "id = ?", id)
	if res.Error != nil {
		return model.Account{}, fmt.Errorf("GetAccount: %v", res.Error)
	}
	return account, nil
}
