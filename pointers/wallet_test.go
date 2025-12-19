package pointers

import (
	"testing"
)

// func TestWalletBalance(t *testing.T) {
// 	w := Wallet{}
// 	initialBalance := Bitcoin(103)

// 	w.Deposit(Bitcoin(1033))

// 	balance := w.Balance()
// 	fmt.Printf("address of balance in test is %p \n", &w.balance)

// 	if balance != initialBalance {
// 		t.Errorf("got %s want %s", balance, initialBalance)
// 	}
// }

// func TestWallet(t *testing.T) {

// 	t.Run("deposit", func(t *testing.T) {
// 		wallet := Wallet{}

// 		wallet.Deposit(Bitcoin(10))

// 		got := wallet.Balance()

// 		want := Bitcoin(10)

// 		if got != want {
// 			t.Errorf("got %s want %s", got, want)
// 		}
// 	})

// 	t.Run("withdraw", func(t *testing.T) {
// 		wallet := Wallet{balance: Bitcoin(20)}

// 		wallet.Withdraw(Bitcoin(10))

// 		got := wallet.Balance()

// 		want := Bitcoin(10)

//			if got != want {
//				t.Errorf("got %s want %s", got, want)
//			}
//		})
//	}
func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t testing.TB, got error, want string) {
		t.Helper()

		if got == nil {
			t.Fatal("didn't get an error but wanted one")
		}

		if got.Error() != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}
	assertNoError := func(t testing.TB, got error) {
		t.Helper()
		if got != nil {
			t.Fatal("got an error but didn't want one")
		}
	}
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		var err = wallet.Withdraw(Bitcoin(10))
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))

	})
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds.Error())
	})

}
