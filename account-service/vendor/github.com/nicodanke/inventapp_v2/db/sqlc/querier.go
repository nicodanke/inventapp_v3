// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Account, error)
	CreateAccountPlan(ctx context.Context, arg CreateAccountPlanParams) (AccountPlan, error)
	CreateRole(ctx context.Context, arg CreateRoleParams) (Role, error)
	CreateRolePermission(ctx context.Context, arg CreateRolePermissionParams) (RolePermission, error)
	CreateSession(ctx context.Context, arg CreateSessionParams) (Session, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteAccount(ctx context.Context, id int64) error
	DeleteRole(ctx context.Context, arg DeleteRoleParams) error
	DeleteUser(ctx context.Context, arg DeleteUserParams) error
	GetAccount(ctx context.Context, id int64) (Account, error)
	GetAccountByCode(ctx context.Context, code string) (Account, error)
	GetPlanCountry(ctx context.Context, arg GetPlanCountryParams) (PlanCountry, error)
	GetRole(ctx context.Context, arg GetRoleParams) (Role, error)
	GetRoles(ctx context.Context, arg GetRolesParams) ([]Role, error)
	GetSession(ctx context.Context, id uuid.UUID) (Session, error)
	GetUser(ctx context.Context, arg GetUserParams) (User, error)
	GetUserByUsername(ctx context.Context, username string) (User, error)
	GetUsers(ctx context.Context, arg GetUsersParams) ([]User, error)
	Getplan(ctx context.Context, id int64) (Plan, error)
	ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error)
	ListPlans(ctx context.Context, arg ListPlansParams) ([]Plan, error)
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error)
	UpdateAccountPlan(ctx context.Context, arg UpdateAccountPlanParams) (AccountPlan, error)
}

var _ Querier = (*Queries)(nil)