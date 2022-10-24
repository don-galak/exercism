package account

import "sync"

type Account struct {
	balance int64
	active  bool
	mutex   sync.RWMutex
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{balance: amount, active: true}
}
func (a *Account) Balance() (int64, bool) {
	a.mutex.RLock()
	defer a.mutex.RUnlock()
	return a.balance, a.active
}
func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if !a.active || a.balance+amount < 0 {
		return 0, false
	}
	a.balance += amount
	return a.balance, true
}
func (a *Account) Close() (int64, bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if !a.active {
		return 0, false
	}
	a.active = false
	balanceBeforeClosing := a.balance
	a.balance = 0
	return balanceBeforeClosing, true
}
