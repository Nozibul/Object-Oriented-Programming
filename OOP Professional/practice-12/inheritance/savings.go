// savings.go
package inheritance

type SavingsAccount struct {
	Account
	InterestRate  float64
	WithdrawLimit int
}

func (s *SavingsAccount) ApplyInterest() {
	s.Balance += s.Balance * s.InterestRate
}
