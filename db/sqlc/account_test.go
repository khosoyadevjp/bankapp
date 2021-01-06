package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/khosoyadevjp/bankapp/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	arg := CreateAccountParams{
		Name:     util.RandomAccountName(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err, "Something wrong occured when creating accunt")
	require.NotEmpty(t, account, "Returned empty account")

	require.Equal(t, arg.Name, account.Name, "Name doesnt match the input value")
	require.Equal(t, arg.Balance, account.Balance, "Balance doesnt match the input value")
	require.Equal(t, arg.Currency, account.Currency, "Currency doesnt match the input value")

	require.NotZero(t, account.ID, "ID Should not be Zero")
	require.NotZero(t, account.CreatedAt, "CreatedAt should not be zero")

	return account
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	testAccount := createRandomAccount(t)
	retrievedAccount, err := testQueries.GetAccount(context.Background(), testAccount.ID)

	require.NoError(t, err)
	require.NotEmpty(t, retrievedAccount)

	require.Equal(t, testAccount.ID, retrievedAccount.ID)
	require.Equal(t, testAccount.Name, retrievedAccount.Name)
	require.Equal(t, testAccount.Balance, retrievedAccount.Balance)
	require.Equal(t, testAccount.Currency, retrievedAccount.Currency)
	require.WithinDuration(t, testAccount.CreatedAt, retrievedAccount.CreatedAt, time.Second)
}

func TestListAccounts(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 5,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}

}

func TestUpdateAccount(t *testing.T) {
	testAccount := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      testAccount.ID,
		Balance: util.RandomMoney(),
	}

	updatedAccount, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount)

	require.Equal(t, testAccount.ID, updatedAccount.ID)
	require.Equal(t, testAccount.Name, updatedAccount.Name)
	require.Equal(t, arg.Balance, updatedAccount.Balance)
	require.Equal(t, testAccount.Currency, updatedAccount.Currency)
	require.WithinDuration(t, testAccount.CreatedAt, updatedAccount.CreatedAt, time.Second)

}

func TestDeleteAccount(t *testing.T) {
	testAccount := createRandomAccount(t)
	err := testQueries.DeleteAccount(context.Background(), testAccount.ID)
	require.NoError(t, err)

	noAccount, err := testQueries.GetAccount(context.Background(), testAccount.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, noAccount)
}
