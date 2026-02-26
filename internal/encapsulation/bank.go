package encapsulation

// Account demonstrates ENCAPSULATION in Go.
// Encapsulation: hide internal state and expose only controlled behavior via methods.
// Go achieves this through capitalization: unexported names stay package-private.
type Account struct {
	// balance is unexported (lowercase), only this package can read/write it directly.
	// This field is unexported to achieve encapsulation.
	balance float64
}

// NewAccount is a constructor that returns an Account with an initial balance.
// Using a constructor allows us to enforce invariants (e.g., non-negative balance).
func NewAccount(initialBalance float64) *Account {
	if initialBalance < 0 {
		initialBalance = 0
	}
	return &Account{balance: initialBalance}
}

// Balance returns the current balance. Exported method provides read-only access.
// Encapsulation: external code cannot modify balance except through Deposit/Withdraw.
func (a *Account) Balance() float64 {
	return a.balance
}

// Deposit adds amount to the balance. Exported method for controlled mutation.
// Encapsulation: we validate input and update state in one place.
func (a *Account) Deposit(amount float64) {
	if amount > 0 {
		a.balance += amount
	}
}

// Withdraw subtracts amount from the balance. Returns false if insufficient funds.
// Encapsulation: business rules (e.g., no negative balance) are enforced here.
func (a *Account) Withdraw(amount float64) bool {
	if amount > 0 && amount <= a.balance {
		a.balance -= amount
		return true
	}
	return false
}
