package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T) {
	store := NewStore(testDB)
	account1 := createRandomAccount(t)
	account2 := createRandomAccount(t)

	var amount int64 = 10
	results := make(chan TransferTxResult)
	errs := make(chan error)

	go func() {
		result, err := store.TransferTx(context.Background(), TransferTxParams{
			FromAccountID: account1.ID,
			ToAccountID:   account2.ID,
			Amount:        amount,
		})
		results <- result
		errs <- err
	}()

	result := <-results
	err := <-errs

	require.NoError(t, err)
	require.NotEmpty(t, result)

	transfer := result.Transfer
	require.NotEmpty(t, transfer)
	require.Equal(t, transfer.FromAccountID, account1.ID)
	require.Equal(t, transfer.ToAccountID, account2.ID)
	require.Equal(t, transfer.Amount, amount)
	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	_, err = store.GetTransfer(context.Background(), transfer.ID)
	require.NoError(t, err)

	fromEntry := result.FromEntry
	require.NotEmpty(t, fromEntry)
	require.Equal(t, fromEntry.AccountID, account1.ID)
	require.Equal(t, fromEntry.Amount, -amount)
	require.NotZero(t, fromEntry.ID)
	require.NotZero(t, fromEntry.CreatedAt)

	_, err = store.GetEntry(context.Background(), fromEntry.ID)
	require.NoError(t, err)

	ToEntry := result.ToEntry
	require.NotEmpty(t, ToEntry)
	require.Equal(t, ToEntry.AccountID, account2.ID)
	require.Equal(t, ToEntry.Amount, amount)
	require.NotZero(t, ToEntry.ID)
	require.NotZero(t, ToEntry.CreatedAt)

	_, err = store.GetEntry(context.Background(), ToEntry.ID)
	require.NoError(t, err)

	fromAccount := result.FromAccount
	require.NotEmpty(t, fromAccount)
	require.Equal(t, fromAccount.ID, account1.ID)

	toAccount := result.ToAccount
	require.NotEmpty(t, toAccount)
	require.Equal(t, toAccount.ID, account2.ID)

	diff1 := account1.Balance - fromAccount.Balance
	diff2 := toAccount.Balance - account2.Balance
	require.Equal(t, int64(diff1), int64(diff2))
	require.True(t, diff1 > 0)
	require.True(t, diff1%amount == 0)

	updatedAccount1, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount1)

	updatedAccount2, err := testQueries.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)
	require.NotEmpty(t, updatedAccount2)
	require.Equal(t, account1.Balance-amount, updatedAccount1.Balance)
	require.Equal(t, account2.Balance+amount, updatedAccount2.Balance)

}
