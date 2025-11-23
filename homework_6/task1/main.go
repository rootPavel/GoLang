package main

import "fmt"

type PaymentProcessor interface {
	Process(ammount float64) string
}

type CreditCard struct {
	Owner   string
	Number  int
	Balance float64
}

func (pay *CreditCard) Process(ammount float64) string {
	receipt := "----- Start Process Credit Card -----\n"
	receipt += fmt.Sprintf("Owner: %s\nCard: %d\n", pay.Owner, pay.Number)
	if ammount < 0 {
		receipt += "❌ Ошибка операции:\n\tСумма платежа должна быть положительной!\n---------------- END -----------------\n"
		return receipt
	}

	if pay.Balance < ammount {
		receipt += fmt.Sprintf("❌ Ошибка операции:\n\tСумма платежа %.2f больше остатка средств на счету %.2f\n---------------- END -----------------\n", ammount, pay.Balance)
		return receipt
	}
	pay.Balance -= ammount
	receipt += fmt.Sprintf("✅ Платеж по карте %d в размере %.2f успешно совершен!\nОстаток средств на счету %.2f\n---------------- END -----------------\n", pay.Number, ammount, pay.Balance)
	return receipt
}

type CryptoWallet struct {
	Owner   string
	Address string
	Balance float64
}

func (pay *CryptoWallet) Process(ammount float64) string {
	receipt := "----- Start Process Crypto Wallet -----\n"
	receipt += fmt.Sprintf("Owner: %s\nAdress wallet: %s\n", pay.Owner, pay.Address)
	if ammount < 0 {
		receipt += "❌ Ошибка операции:\n\tСумма платежа должна быть положительной!\n----------------- END -----------------\n"
		return receipt
	}

	if pay.Balance < ammount {
		receipt += fmt.Sprintf("❌ Ошибка операции:\n\tСумма платежа %.8f больше остатка средств на счету %.8f\n----------------- END -----------------\n", ammount, pay.Balance)
		return receipt
	}

	pay.Balance -= ammount
	receipt += fmt.Sprintf("✅ C вашего кошелька %s списано %.8f успешно!\nОстаток средств на счету %.8f\n----------------- END -----------------\n", pay.Address, ammount, pay.Balance)
	return receipt
}

func main() {
	account1 := &CreditCard{
		Owner:   "Sergey",
		Number:  772345455222,
		Balance: 234.45,
	}
	account2 := &CryptoWallet{
		Owner:   "Oleg",
		Address: "c33md93Sqce0w5",
		Balance: 235.45456656,
	}

	accounts := []PaymentProcessor{account1, account2}

	for _, v := range accounts {
		fmt.Println(v.Process(115.12))
		fmt.Println(v.Process(-10))
		fmt.Println(v.Process(515))
	}
}
