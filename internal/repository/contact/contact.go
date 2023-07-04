package contact

import (
	"context"
	"gorm.io/gorm"
	"week3_docker/internal/model"
)

type ListContactsFilter struct {
	AccountID  uint64
	Page       int
	Limit      int
	Type       string
	AmoIDs     []uint64
	Sync       *bool
	FillReason *bool
}

func (r Repository) UpdateContact(ctx context.Context, contact *model.Contact) error {
	if err := r.Store.DB.WithContext(ctx).Updates(contact).Error; err != nil {
		return err
	}
	return nil
}

func (r Repository) UpdateContactsByIDs(ctx context.Context, ids []uint64, c *model.Contact) (int64, error) {
	res := r.Store.DB.WithContext(ctx).Where("id IN ?", ids).Updates(c)
	if res.Error != nil {
		return 0, res.Error
	}
	return res.RowsAffected, res.Error
}

func (r Repository) DeleteAccountContacts(ctx context.Context, accountID uint64) error {
	if err := r.Store.DB.WithContext(ctx).Where("account_id = ?", accountID).Delete(&model.Contact{}).Error; err != nil {
		return err
	}
	return nil
}

func (r Repository) CreateContact(ctx context.Context, contact *model.Contact) error {
	if err := r.Store.DB.WithContext(ctx).Create(contact).Error; err != nil {
		return err
	}
	return nil
}

func (r Repository) InsertContacts(ctx context.Context, contacts []model.Contact) (int64, error) {
	res := r.Store.DB.WithContext(ctx).Create(&contacts)
	if res.Error != nil {
		return 0, res.Error
	}
	return res.RowsAffected, nil
}
func (r Repository) ListContacts(ctx context.Context, filter ListContactsFilter) ([]model.Contact, error) {
	var contacts []model.Contact
	q := r.Store.DB.WithContext(ctx)
	q = applyListContactsFilter(q, filter).Order("id")
	if err := q.Find(&contacts).Error; err != nil {
		return nil, err
	}

	return contacts, nil
}

func applyListContactsFilter(q *gorm.DB, filter ListContactsFilter) *gorm.DB {
	if filter.Limit < 1 {
		filter.Limit = 100
	}
	if filter.Page < 1 {
		filter.Page = 1
	}

	q = q.Limit(filter.Limit).Offset((filter.Page - 1) * filter.Limit)

	s := make(map[string]interface{})
	if filter.AccountID != 0 {
		s["account_id"] = filter.AccountID
	}
	if filter.Type != "" {
		s["type"] = filter.Type
	}
	if filter.Sync != nil {
		s["sync"] = *filter.Sync
	}
	if filter.AmoIDs != nil {
		s["amo_id"] = filter.AmoIDs
	}

	if filter.FillReason != nil {
		if *filter.FillReason {
			q = q.Where("reason_out_sync <> \"\"")
		} else {
			q = q.Where("reason_out_sync = \"\"")
		}

	}

	q = q.Where(s)

	return q
}

func (r Repository) DeleteContact(ctx context.Context, contact *model.Contact) error {
	if err := r.Store.DB.WithContext(ctx).Unscoped().Delete(contact).Error; err != nil {
		return err
	}
	return nil
}

func (r Repository) UpdateContactsByAmoIDs(ctx context.Context, amoIDs []uint64, contact *model.Contact) error {
	res := r.Store.DB.WithContext(ctx).Model(&model.Contact{}).Where("amo_id IN ?", amoIDs).Updates(map[string]interface{}{
		"type": contact.Type,
		"sync": contact.Sync,
	})
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (r Repository) DeleteContacts(ctx context.Context, contacts []model.Contact) (int64, error) {
	res := r.Store.DB.WithContext(ctx).Unscoped().Delete(&contacts)
	if res.Error != nil {
		return 0, res.Error
	}

	return res.RowsAffected, nil
}

func (r Repository) DeleteContactsByAmoIDs(ctx context.Context, amoIDs []uint64) (int64, error) {
	res := r.Store.DB.WithContext(ctx).Unscoped().Where("amo_id IN ?", amoIDs).Delete(&model.Contact{})
	if res.Error != nil {
		return 0, res.Error
	}
	return res.RowsAffected, nil
}
