package account

import (
	"context"
	"fmt"
	gorm "gorm.io/gorm"
	"gorm.io/gorm/clause"
	"time"
	"week3_docker/internal/model"
)

func (r Repository) CreateAccount(ctx context.Context, account *model.Account) error {
	res := r.Store.DB.WithContext(ctx).
		Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			UpdateAll: true,
		}).
		Create(account)
	if res.Error != nil {
		return fmt.Errorf("CreateAccount: %v", res.Error)
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

func (r Repository) GetAccount(ctx context.Context, id uint64) (model.Account, error) {
	var account model.Account
	res := r.Store.DB.WithContext(ctx)
	if id != 0 {
		res = res.Where("id = ?", id)
	}

	res = res.First(&account)
	if res.Error != nil {
		return model.Account{}, fmt.Errorf("GetAccount: %v", res.Error)
	}
	return account, nil
}

func (r Repository) ListAccounts(ctx context.Context, filter ListAccountFilter) ([]model.Account, error) {
	q := r.Store.DB.WithContext(ctx)
	q = applyListAccountFilter(q, filter)

	var accounts []model.Account

	res := q.Find(&accounts)

	if res.Error != nil {
		return nil, res.Error
	}

	return accounts, nil
}

func applyListAccountFilter(q *gorm.DB, filter ListAccountFilter) *gorm.DB {
	if filter.JoinIntegrations {
		q = q.Preload("Integrations")
	}

	if filter.Limit < 1 {
		filter.Limit = 100
	}
	if filter.Page < 1 {
		filter.Page = 1
	}

	q = q.Limit(filter.Limit).Offset((filter.Page - 1) * filter.Limit)

	if filter.NeedRefresh {
		q = q.Where("access_token <> \"\" AND expires <= ?", time.Now().Add(time.Hour*5).Unix()).Order("expires")
	}

	if filter.AmoAuthorized != nil {
		if *filter.AmoAuthorized {
			q = q.Where("access_token <> \"\" AND expires > ?", time.Now().Unix()).Order("id")
		} else {
			q = q.Where("access_token = \"\" OR expires <= ?", time.Now().Unix())
		}
	}

	return q
}

func (r Repository) UpdateAccount(ctx context.Context, account *model.Account) error {
	tx := r.Store.DB.WithContext(ctx).Updates(*account)

	if tx.Error != nil {
		return fmt.Errorf("UpdateAccount: %v", tx.Error)
	}
	return nil
}

type ListAccountFilter struct {
	Page             int
	Limit            int
	NeedRefresh      bool
	AmoAuthorized    *bool
	JoinIntegrations bool
}
