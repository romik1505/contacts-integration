package contact

import (
	"context"
	"gorm.io/gorm"
	"week3_docker/internal/model"
)

func applyListContactsFilter(q *gorm.DB, filter model.ListContactsFilter) *gorm.DB {
	if filter.Limit > 0 {
		q = q.Limit(filter.Limit)
	}
	if filter.Page > 0 {
		if filter.Limit < 1 {
			filter.Limit = 100
		}
		q = q.Offset((filter.Page - 1) * filter.Limit)
	}
	s := make(map[string]interface{})
	if filter.AccountID >= 0 {
		s["account_id"] = filter.AccountID
	}
	if filter.Type != "" {
		s["type"] = filter.Type
	}
	if filter.Sync != nil {
		s["sync"] = *filter.Sync
	}
	q = q.Where(s)

	return q
}

func (r Repository) ListContacts(ctx context.Context, filter model.ListContactsFilter) ([]model.Contact, error) {
	var contacts []model.Contact
	q := r.Store.DB.WithContext(ctx)
	q = applyListContactsFilter(q, filter).Order("id")
	if err := q.Find(&contacts).Error; err != nil {
		return nil, err
	}
	return contacts, nil
}

func (r Repository) CreateContact(ctx context.Context, contact *model.Contact) error {
	if err := r.Store.DB.WithContext(ctx).Create(contact).Error; err != nil {
		return err
	}
	return nil
}

func (r Repository) InsertContacts(ctx context.Context, contacts []*model.Contact) (int64, error) {
	res := r.Store.DB.WithContext(ctx).Create(contacts)
	if res.Error != nil {
		return 0, res.Error
	}

	return res.RowsAffected, nil
}

func (r Repository) UpdateContact(ctx context.Context, contact *model.Contact) error {
	if err := r.Store.DB.WithContext(ctx).Updates(contact).Error; err != nil {
		return err
	}
	return nil
}

func (r Repository) DeleteAccountContacts(ctx context.Context, accountID uint64) error {
	if err := r.Store.DB.WithContext(ctx).Where("account_id = ?", accountID).Delete(&model.Contact{}).Error; err != nil {
		return err
	}
	return nil
}
