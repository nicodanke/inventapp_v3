package db

import (
	"context"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/nicodanke/inventapp_v3/account-service/utils"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {

	arg := CreateAccountParams{
		Code:        utils.RandomString(6),
		CompanyName: utils.RandomString(6),
		Email:       utils.RandomEmail(),
		Country:     utils.RandomString(6),
	}

	account, err := testStore.CreateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Code, account.Code)
	require.Equal(t, arg.CompanyName, account.CompanyName)
	require.Equal(t, arg.Email, account.Email)
	require.Equal(t, arg.Country, account.Country)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testStore.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Code, account2.Code)
	require.Equal(t, account1.Country, account2.Country)
	require.Equal(t, account1.Email, account2.Email)
	require.Equal(t, account1.CompanyName, account2.CompanyName)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestGetAccountByCode(t *testing.T) {
	account1 := createRandomAccount(t)
	account2, err := testStore.GetAccountByCode(context.Background(), account1.Code)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Code, account2.Code)
	require.Equal(t, account1.Country, account2.Country)
	require.Equal(t, account1.Email, account2.Email)
	require.Equal(t, account1.CompanyName, account2.CompanyName)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	account1 := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:          account1.ID,
		CompanyName: pgtype.Text{String: utils.RandomString(6), Valid: true},
	}

	account2, err := testStore.UpdateAccount(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Code, account2.Code)
	require.Equal(t, arg.CompanyName.String, account2.CompanyName)
	require.Equal(t, account1.Email, account2.Email)
	require.Equal(t, account1.Country, account2.Country)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	account1 := createRandomAccount(t)
	err := testStore.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	account2, err := testStore.GetAccount(context.Background(), account1.ID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, account2)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 0,
	}

	accounts, err := testStore.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
