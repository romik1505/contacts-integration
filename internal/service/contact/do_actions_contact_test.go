package contact

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
	"week3_docker/internal/client/unisender"
	"week3_docker/internal/model"
	mock_unisender "week3_docker/pkg/mocks/client/unisender"
	mock_contact "week3_docker/pkg/mocks/repository/contact"
)

func TestService_DoAddContacts(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	uni := mock_unisender.NewMockIUnisenderClient(ctl)
	cr := mock_contact.NewMockIRepository(ctl)
	contactService := NewService(nil, uni, nil, cr, nil, nil)

	tests := []struct {
		name       string
		input      model.ContactActionsTask
		wantErr    bool
		res        error
		hookBefore func()
		hookAfter  func()
	}{
		{
			name: ">2 tries",
			input: model.ContactActionsTask{
				TryNumber: 3,
			},
			wantErr: false,
		},
		{
			name: "ok",
			input: model.ContactActionsTask{
				Type:            "add",
				AccountID:       12345678,
				UnisenderKey:    "unisender_api_key_123",
				UnisenderListID: 23,
				Contacts: []model.Contact{
					{
						AmoID:     444,
						AccountID: 12345,
						Name:      "valid valid",
						Email:     "valid@email.ru",
						Type:      "add",
						Sync:      false,
					}, {
						AmoID:     555,
						AccountID: 12345,
						Name:      "simple simple",
						Email:     "simple@email.com",
						Type:      "add",
						Sync:      false,
					},
				},
				IDs: []uint64{444, 555},
			},
			wantErr: false,
			res:     nil,
			hookBefore: func() {
				cr.EXPECT().InsertContacts(gomock.Any(), []model.Contact{
					{
						AmoID:     444,
						AccountID: 12345,
						Name:      "valid valid",
						Email:     "valid@email.ru",
						Type:      "add",
						Sync:      false,
					}, {
						AmoID:     555,
						AccountID: 12345,
						Name:      "simple simple",
						Email:     "simple@email.com",
						Type:      "add",
						Sync:      false,
					}},
				).DoAndReturn(func(ctx context.Context, inp []model.Contact) (int64, error) {
					for i := range inp {
						inp[i].ID = uint64((i + 100) * 2)
					}
					return 2, nil
				})

				uni.EXPECT().ImportContacts(gomock.Any(), unisender.ImportContactsRequest{
					Format:     "json",
					ApiKey:     "unisender_api_key_123",
					FieldNames: []string{"email", "Name", "email_list_ids", "delete"},
					Data: [][]string{
						{"valid@email.ru", "valid valid", "23", "0"},
						{"simple@email.com", "simple simple", "23", "0"},
					},
				}).Return(unisender.ImportContactsResponse{
					Result: &unisender.Result{
						Total:     2,
						Inserted:  2,
						NewEmails: 2,
						Invalid:   0,
					},
				}, nil)

				cr.EXPECT().UpdateContact(gomock.Any(), gomock.Any()).Times(2)
			},
			hookAfter: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.hookBefore != nil {
				tt.hookBefore()
			}

			resp := contactService.DoAddContacts(context.Background(), tt.input)
			if tt.wantErr {
				require.NotNil(t, resp)
			} else {
				require.Equal(t, resp, tt.res)
			}

			if tt.hookAfter != nil {
				tt.hookAfter()
			}
		})
	}
}
