package db

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"simplebank/util"
	"testing"
)

func TestStore_TransferTx(t *testing.T) {
	store := NewStore(testDB)

	fromAccount := createRandomAccount(t)
	toAccount := createRandomAccount(t)

	n := 5

	// Run concurrent transfer transactions
	errs := make(chan error)
	results := make(chan TransferTxResult)

	for i := 0; i < n; i++ {
		go func() {
			amount := int64(util.RandomInt(0, 1000))

			result, err := store.TransferTx(context.Background(), TransferTxParams{
				FromAccountID: fromAccount.ID,
				ToAccountID:   toAccount.ID,
				Amount:        amount,
			})

			errs <- err
			results <- *result
		}()
	}

	// Check the results
	for i := 0; i < n; i++ {
		err := <-errs

		require.NoError(t, err)

		result := <-results

		require.NotEmpty(t, result)
		require.Equal(t, fromAccount.ID, result.FromAccount.ID)
		require.Equal(t, toAccount.ID, result.ToAccount.ID)

		fmt.Println(result)
	}
}
