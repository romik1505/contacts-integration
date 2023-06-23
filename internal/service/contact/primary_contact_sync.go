package contact

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	contact "week3_docker/pkg/api/contact_service"
)

func (s Service) PrimaryContactsSync(ctx context.Context, req *contact.PrimaryContactSyncRequest) (*emptypb.Empty, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}
