package main

import "fmt"

// PaymentProcessor определяет поведение для способов оплаты
type PaymentProcessor interface {
	ProcessPayment(amount float64) string
}

// CreditCard структура
type CreditCard struct {
	CardNumber string
}

// ProcessPayment способ оплаты для CreditCard
func (cc CreditCard) ProcessPayment(amount float64) string {
	return fmt.Sprintf(
		"Processed payment of $%.2f using Credit Card %s",
		amount,
		cc.CardNumber)
}

// PayPal структура
type PayPal struct {
	Email string
}

// ProcessPayment способ оплаты для PayPal
func (pp PayPal) ProcessPayment(amount float64) string {
	return fmt.Sprintf(
		"Processed payment of $%.2f using PayPal account %s",
		amount,
		pp.Email)
}

// BankTransfer структура
type BankTransfer struct {
	AccountNumber string
}

// ProcessPayment способ оплаты для BankTransfer
func (bt BankTransfer) ProcessPayment(amount float64) string {
	return fmt.Sprintf(
		"Processed payment of $%.2f using Bank Transfer %s",
		amount,
		bt.AccountNumber)
}

// process функция использует интерфейс PaymentProcessor interface
func process(payment PaymentProcessor, amount float64) {
	fmt.Println(payment.ProcessPayment(amount))
}

func main() {
	cc := CreditCard{CardNumber: "1234-5678-8765-4444"}
	pp := PayPal{Email: "user1@codeheim.io"}
	bt := BankTransfer{AccountNumber: "987654321"}

	process(cc, 100.95)
	process(pp, 170.75)
	process(bt, 250.50)
}
