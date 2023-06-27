package integration

import (
	"context"
	"gorm.io/gorm"
	"week3_docker/internal/model"
)

func (r Repository) CreateIntegration(ctx context.Context, integration *model.Integration) error {
	if err := r.Store.DB.WithContext(ctx).Create(integration).Error; err != nil {
		return err
	}
	return nil
}

func (r Repository) UpdateIntegration(ctx context.Context, integration *model.Integration) error {
	if err := r.Store.DB.WithContext(ctx).Create(integration).Error; err != nil {
		return err
	}
	return nil
}

func applyListIntegratonFilter(q *gorm.DB, f model.ListIntegrationFilter) *gorm.DB {
	if f.Limit > 0 {
		q = q.Limit(f.Limit)
	}
	if f.Page > 0 {
		if f.Limit < 1 {
			f.Limit = 100
		}
		q = q.Offset((f.Page - 1) * f.Limit)
	}
	if f.AccountID != 0 {
		q = q.Where("account_id = ?", f.AccountID)
	}
	return q
}

func (r Repository) ListIntegration(ctx context.Context, filter model.ListIntegrationFilter) ([]model.Integration, error) {
	var integrations []model.Integration
	q := r.Store.DB.WithContext(ctx)
	q = applyListIntegratonFilter(q, filter).Order("id")

	if err := q.Find(&integrations).Error; err != nil {
		return nil, err
	}
	return integrations, nil
}

func (r Repository) GetIntegration(ctx context.Context, inp *model.Integration) (model.Integration, error) {
	var integration model.Integration
	if err := r.Store.DB.WithContext(ctx).Where(inp).First(&integration).Error; err != nil {
		return model.Integration{}, err
	}
	return integration, nil
}

func (r Repository) DeleteIntegrations(ctx context.Context, accountID uint64) error {
	if err := r.Store.DB.Where("account_id = ?", accountID).Delete(&model.Integration{}).Error; err != nil {
		return err
	}
	return nil
}
