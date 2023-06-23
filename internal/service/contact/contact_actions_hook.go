package contact

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"week3_docker/internal/mapper"
	contact "week3_docker/pkg/api/contact_service"
)

func (s Service) ContactActionsHook(ctx context.Context, req *contact.ContactActionsHookRequest) (*emptypb.Empty, error) {
	log.Println(req.GetContacts().GetAdd())

	addContacts := mapper.ConvertAmoContacts(req.GetContacts().GetAdd(), req.GetId(), "add")
	if len(addContacts) > 0 {
		_, err1 := s.cr.InsertContacts(ctx, addContacts)
		if err1 != nil {
			return nil, err1
		}
	}

	updateContacts := mapper.ConvertAmoContacts(req.GetContacts().GetUpdate(), req.GetId(), "update")
	if len(updateContacts) > 0 {
		_, err2 := s.cr.InsertContacts(ctx, updateContacts)
		if err2 != nil {
			return nil, err2
		}
	}

	deleteContacts := mapper.ConvertAmoContacts(req.GetContacts().GetDelete(), req.GetId(), "delete")
	if len(deleteContacts) > 0 {
		_, err3 := s.cr.InsertContacts(ctx, deleteContacts)
		if err3 != nil {
			return nil, err3
		}
	}
	log.Printf("ContactHook: add=%d, update=%d, delete=%d", len(addContacts), len(updateContacts), len(deleteContacts))
	return new(emptypb.Empty), nil
}
