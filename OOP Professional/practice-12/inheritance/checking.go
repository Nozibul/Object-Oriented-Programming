package inheritance

type CheckingAccount struct {
	Account
	ServiceCharge float64
}

func (c *CheckingAccount) ApplyServiceCharge() {
	c.Balance -= c.ServiceCharge
}
