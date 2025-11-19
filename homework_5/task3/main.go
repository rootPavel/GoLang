/*
Задача 3: Структура "Банковский счет" и методы для работы с балансом
Описание:
Создайте структуру BankAccount с полями: Owner (владелец счета) и Balance (текущий баланс).
Реализуйте методы:
Deposit(amount float64), увеличивающий баланс.
Withdraw(amount float64), уменьшающий баланс, если хватает средств, иначе выводит сообщение о недостатке средств.
Что нужно сделать:
Объявить структуру и методы.
Создать счет, пополнить его, попытаться снять деньги и вывести итоговый баланс.
*/

package main

import "fmt"

type BankAccount struct {
	Owner   string
	Balance float64
}

func (b *BankAccount) Deposit(amount float64) {
	fmt.Println("     *****")
	fmt.Printf("Ваш текущий баланс: %.2f денег\n", b.Balance)
	b.Balance += amount
	fmt.Printf("Ваш баланс пополнен на %.2f денег\n", amount)
	fmt.Printf("Ваш текущий баланс: %.2f денег\n", b.Balance)
}

func (b *BankAccount) Withdraw(amount float64) {
	fmt.Println("     *****")
	fmt.Printf("Ваш текущий баланс: %.2f денег\n", b.Balance)
	if amount < 0 {
		fmt.Printf("Вы снимаете %.2f денег\n", amount)
		fmt.Println("Сумма не может быть отрицательной!")
		return
	}
	if b.Balance < amount {
		fmt.Println("На вашем счету недостаточно денег")
		return
	}
	b.Balance -= amount
	fmt.Printf("Вы снимаете %.2f денег\n", amount)
	fmt.Printf("Ваш текущий баланс: %.2f денег\n", b.Balance)
}

func main() {
	Account1 := BankAccount{
		Owner:   "Philip",
		Balance: 156.10,
	}
	Account1.Deposit(100)
	Account1.Withdraw(200)
	Account1.Withdraw(56.10)
	Account1.Withdraw(-200)
}
