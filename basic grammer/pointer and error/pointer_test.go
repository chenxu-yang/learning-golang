package pointer_and_error

import (
	"errors"
	"fmt"
	"testing"
)

type Bitcoin int
type Wallet struct {
	balance Bitcoin
}
type Stringer interface {
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
	//定义打印方式
}
func (w *Wallet) Deposit(amount Bitcoin) {
	//fmt.Println("address of balance in deposit is",&w.balance)
	w.balance += amount
}

var InsufficientFundsError = errors.New("cannot withdraw,insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance < amount {
		return InsufficientFundsError
	}
	w.balance -= amount
	return nil
}
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		//fmt.Println("address of balance in test is",&wallet.balance)
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(10)
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))
		assertBalance(t, wallet, Bitcoin(20))
		assertError(t, err, "cannot withdraw,insufficient funds")

	})
}
func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
func assertError(t *testing.T, got error, want string) {
	if got == nil { //nil just like null
		t.Fatal("want an error but didn't get one")
	}
	if got.Error() != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

//summary
//pointer:直接传值给函数或方法时，go会复制这些值。因此，如果函数要改变状态，需要指针指向想要更改的值
//nil: 指针可以是nil
