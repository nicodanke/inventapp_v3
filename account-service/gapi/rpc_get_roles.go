package gapi

import (
	"context"
	"fmt"

	db "github.com/nicodanke/inventapp_v3/account-service/db/sqlc"
	"github.com/nicodanke/inventapp_v3/account-service/pb/requests/v1/role"
	"github.com/nicodanke/inventapp_v3/account-service/validators"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func (server *Server) GetRoles(ctx context.Context, req *role.GetRolesRequest) (*role.GetRolesResponse, error) {
	authPayload, err := server.authorizeUser(ctx)
	if err != nil {
		return nil, unauthenticatedError(fmt.Sprintln("", err))
	}

	violations := validateGetRolesRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	limit := int32(50)
	if req.Size != nil {
		limit = req.GetSize()
	}

	offset := int32(0)
	if req.Page != nil {
		offset = req.GetPage() * limit
	}

	arg := db.GetRolesParams{
		AccountID: authPayload.AccountID,
		Limit:     limit,
		Offset:    offset,
	}

	result, err := server.store.GetRoles(ctx, arg)
	if err != nil {
		return nil, internalError(fmt.Sprintln("Failed to get roles:", err))
	}

	rsp := &role.GetRolesResponse{
		Roles: convertRoles(result),
	}
	return rsp, nil
}

func validateGetRolesRequest(req *role.GetRolesRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if req.Page != nil {
		if err := validators.ValidatePage(req.GetPage()); err != nil {
			violations = append(violations, fieldViolation("page", err))
		}
	}

	if req.Size != nil {
		if err := validators.ValidateSize(req.GetSize()); err != nil {
			violations = append(violations, fieldViolation("size", err))
		}
	}

	return violations
}
