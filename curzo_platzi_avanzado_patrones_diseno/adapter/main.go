package main

import "fmt"

// Payment is an interface that defines a method for payment processing.
type Payment interface {
	Pay()
}

// CashPayment implements the Payment interface for cash transactions.
type CashPayment struct{}

func (c CashPayment) Pay() {
	fmt.Println("Paying cash")
}

// ProcessPayment processes payments of any type that satisfies the Payment interface.
func ProcessPayment(p Payment) {
	p.Pay()
}

// BankPayment is a struct that requires a bank account number to process a payment.
type BankPayment struct{}

// Pay prints the bank account number to which the payment is made.
func (b BankPayment) Pay(bankAccount int) {
	fmt.Printf("Paying bank account: %d\n", bankAccount)
}

// BankPaymentAdapter adapts BankPayment to the Payment interface.
type BankPaymentAdapter struct {
	BankPayment *BankPayment // Holds an instance of BankPayment
	bankAccount int          // Bank account number to be used
}

// Pay allows BankPaymentAdapter to satisfy the Payment interface using a BankPayment method.
func (b BankPaymentAdapter) Pay() {
	b.BankPayment.Pay(b.bankAccount)
}

func main() {
	// Cash payment processing
	cash := CashPayment{}
	ProcessPayment(cash)

	// Bank payment processing using an adapter
	// The BankPayment struct does not directly satisfy the Payment interface because its Pay method requires an argument.
	// The BankPaymentAdapter struct makes it compatible with the Payment interface.
	bankAdapter := BankPaymentAdapter{BankPayment: &BankPayment{}, bankAccount: 1234567890}
	ProcessPayment(bankAdapter)

	fmt.Println("Adapter pattern demonstration")
}
