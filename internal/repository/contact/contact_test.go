package contact

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
	"week3_docker/internal/model"
	"week3_docker/internal/store"
)

func TestRepository_InsertContacts(t *testing.T) {
	st := store.NewStore()

	rep := NewRepository(st)

	type wantStruct struct {
		AffectedRowNum int64
		err            error
	}
	tests := []struct {
		name       string
		inp        []model.Contact
		want       wantStruct
		hookBefore func()
		hookAfter  func()
		wantErr    bool
	}{
		{
			name: "ok insert",
			inp: []model.Contact{
				{
					ID:        0,
					AmoID:     123,
					AccountID: 23,
					Name:      "CONTACT_1",
					Email:     "contact1@mail.com",
					Type:      "init",
				},
				{
					ID:        0,
					AmoID:     234,
					AccountID: 23,
					Name:      "CONTACT_2",
					Email:     "contact1@mail.com",
					Type:      "init",
				},
				{
					ID:        0,
					AmoID:     345,
					AccountID: 23,
					Name:      "CONTACT_3",
					Email:     "contact3@yandex.ru",
					Type:      "init",
				},
			},
			want: wantStruct{
				AffectedRowNum: 3,
				err:            nil,
			},
			hookBefore: func() {
				mustExec(t, &st, "delete from contacts")
				mustExec(t, &st, "delete from integrations")
				mustExec(t, &st, "delete from accounts")

				mustExec(t, &st, "insert into accounts (id) values (23)")
			},
			hookAfter: func() {
				mustExec(t, &st, "delete from contacts")
				mustExec(t, &st, "delete from integrations")
				mustExec(t, &st, "delete from accounts")
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.hookBefore != nil {
				tt.hookBefore()
			}

			resAff, err := rep.InsertContacts(context.Background(), tt.inp)
			require.Equal(t, resAff, tt.want.AffectedRowNum)

			if tt.wantErr {
				require.Equal(t, err, tt.want.err)
			} else {
				require.Nil(t, err)
			}

			var contacts []model.Contact
			rep.Store.DB.Order("amo_id").Find(&contacts)

			require.Equal(t, tt.inp, contacts)

			if tt.hookAfter != nil {
				tt.hookAfter()
			}
		})
	}
}

func mustExec(t *testing.T, db *store.Store, sql string) {
	require.Nil(t, db.DB.Exec(sql).Error)
}
