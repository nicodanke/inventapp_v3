package gapi

import (
	"context"
	"fmt"

	db "github.com/nicodanke/inventapp_v2/db/sqlc"
	"github.com/nicodanke/inventapp_v2/pb/requests/v1/user"
	"github.com/nicodanke/inventapp_v2/sse"
	"google.golang.org/protobuf/types/known/emptypb"
)

const (
	sse_delete_user = "delete-user"
)

func (server *Server) DeleteUser(ctx context.Context, req *user.DeleteUserRequest) (*emptypb.Empty, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(fmt.Sprintln("", err))
	}

	if req.GetId() == authPayload.UserID {
		return nil, unprocessableError(ACTION_NOT_ALLOWED, "User cannot auto delete itself")
	}

	arg := db.DeleteUserParams{
		AccountID: authPayload.AccountID,
		ID:        req.GetId(),
	}

	err = server.store.DeleteUser(ctx, arg)
	if err != nil {
		return nil, internalError(fmt.Sprintln("Failed to delete user with id:", req.GetId(), err))
	}

	// Notify delete user
	var data = map[string]any{}
	data["id"] = req.GetId()
	server.notifier.BoadcastMessageToAccount(sse.NewEventMessage(sse_delete_user, data), authPayload.AccountID, &authPayload.UserID)

	return &emptypb.Empty{}, nil
}
