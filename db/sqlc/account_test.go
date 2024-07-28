package db

import (
	"context"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
)

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account := createRandomAccount(t)
	account2, err := testStore.GetAccount(context.Background(), account.ID)
	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account2.ID, account.ID)
	require.Equal(t, account2.Owner, account.Owner)
	require.Equal(t, account2.Balance, account.Balance)
	require.Equal(t, account2.Currency, account.Currency)
	require.WithinDuration(t, account2.CreatedAt, account.CreatedAt, time.Second)
}

func TestUpdateAccount(t *testing.T) {
	account := createRandomAccount(t)

	updateAccountParams := UpdateAccountParams{
		ID:      account.ID,
		Balance: int64(gofakeit.Price(20.0, 100.00)),
	}

	updatedAccount, err := testStore.UpdateAccount(context.Background(), updateAccountParams)
	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount)

	require.Equal(t, updatedAccount.ID, account.ID)
	require.Equal(t, updatedAccount.Owner, account.Owner)
	require.Equal(t, updatedAccount.Balance, updateAccountParams.Balance)
	require.WithinDuration(t, updatedAccount.CreatedAt, account.CreatedAt, time.Second)
}

func TestDeleteAccount(t *testing.T) {
	account := createRandomAccount(t)

	err := testStore.DeleteAccount(context.Background(), account.ID)
	require.NoError(t, err)

	account2, err := testStore.GetAccount(context.Background(), account.ID)
	require.Error(t, err)
	require.EqualError(t, err, ErrRecordNotFound.Error())
	require.Empty(t, account2)
}

func TestListAccounts(t *testing.T) {
	var lastAccount Account
	for i := 0; i < 10; i++ {
		lastAccount = createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Owner:  lastAccount.Owner,
		Limit:  5,
		Offset: 0,
	}

	accounts, err := testStore.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	for _, account := range accounts {
		require.NotEmpty(t, account)
		require.Equal(t, lastAccount.Owner, account.Owner)
	}
}

func createRandomAccount(t *testing.T) Account {
	var createAccountParams CreateAccountParams

	_ = gofakeit.Struct(&createAccountParams)

	account, err := testStore.CreateAccount(context.Background(), createAccountParams)
	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, account.Owner, createAccountParams.Owner)
	require.Equal(t, account.Balance, createAccountParams.Balance)
	require.Equal(t, account.Currency, createAccountParams.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}
