package loans

func IsValid(loan *Loan) bool {
	if loan.PaymentDate.Unix() > loan.EndDate.Unix() || loan.PaidSum > loan.Sum {
		return false
	}

	return true
}
