package payment

import "fmt"

type PaymentMethod interface {
	Pay(total int) error
}

type CashPayment struct {
	Cash int
}

type PaymentService struct {
	PaymentMethod PaymentMethod
}

type DebitPayment struct {
	CardNumber int
}

type QrisPayment struct{}

func (c CashPayment) Pay(total int) error {
	if c.Cash < total {
		return fmt.Errorf("Uang anda tidak cukup.")
	}

	change := c.Cash - total

	fmt.Println("Pembayaran Berhasil!")
	fmt.Println("Kembalian: ", change)

	return nil
}

func (q QrisPayment) Pay(total int) error {
	fmt.Println("Pembayaran QRIS berhasil!")
	fmt.Println("Total: ", total)
	return nil
}

func (s *PaymentService) Pay(total int) error {
	return s.PaymentMethod.Pay(total)
}

func ProcessPayment(payment PaymentMethod, total int) error {
	return payment.Pay(total)
}

func (d DebitPayment) Pay(total int) error {
	fmt.Println("Pembayaran Debit Berhasil!")
	fmt.Println("Total: ", total)

	return nil
}
