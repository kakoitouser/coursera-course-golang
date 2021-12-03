package main

import "fmt"

type Payer interface {
	Pay(int) error
}

type Wallet struct {
	cash int
}

func (w *Wallet) Pay(price int) error {
	if w.cash < price {
		return fmt.Errorf("Не хватает денег")
	}
	w.cash -= price
	return nil
}

type Card struct {
	Balance int
	CVV     string
	Number  string
}

func (c *Card) Pay(price int) error {
	if c.Balance < price {
		return fmt.Errorf("Не хватает денег")
	}
	c.Balance -= price
	return nil
}

type ApplePay struct {
	Money   int
	AppleId string
}

func (a *ApplePay) Pay(price int) error {
	if a.Money < price {
		return fmt.Errorf("Не хватает денег")
	}
	a.Money -= price
	return nil
}

func main() {
	buy := func(p Payer) {
		err := p.Pay(10)
		if err != nil {
			fmt.Println("Ошибка при оплате")
			return
		}
		fmt.Println("Покупка сделано")
	}

	w := &Wallet{cash: 100}
	buy(w)

	func(p Payer) {
		switch p.(type) {
		case *Wallet:
			fmt.Println("Оплата наличными")
		case *Card:
			fmt.Println("Оплата картой")
		case *ApplePay:
			fmt.Println("Оплата с эпл пэй")
		}
	}(w)
}
