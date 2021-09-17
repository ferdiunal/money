package main

import (
	money "github.com/ferdiunal/money/src"
)

func main() {
	m := money.NewMoney(100)

	println(m.AddTax(13.8).GetAmount())
}
