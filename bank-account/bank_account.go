package account

type Account struct {
	balance int64
	active  bool
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{amount, true}
}

func (a *Account) Balance() (int64, bool) {
	if !a.active {
		return 0, false
	}
	return a.balance, a.active
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	if !a.active || a.balance+amount < 0 {
		return 0, false
	}

	a.balance += amount
	return a.balance, true
}

func (a *Account) Close() (int64, bool) {
	if !a.active {
		return 0, false
	}
	a.active = false
	balanceBeforeClosing := a.balance
	a.balance = 0
	return balanceBeforeClosing, true
}