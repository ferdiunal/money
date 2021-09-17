package money

import "fmt"

type Money struct {
	amount    float64
	taxRate   int64
	decimals  int8
	taxAmount float64
}

type MoneyValue struct {
	Amount string
	Tax    string
}

type MoneyInterface interface {
	SetDeciamls(deciamls int8) *Money
	AddTax(percent float64) *Money
	RemoveTax(percent float64) *Money
	All() MoneyValue
	GetTax() string
}

func (m *Money) SetDeciamls(deciamls int8) *Money {
	m.decimals = deciamls

	return m
}

func (m *Money) AddTax(percent float64) *Money {
	if percent == 0 {
		percent = float64(m.taxRate)
	}
	m.taxAmount += percent

	m.amount = m.amount + percent

	return m
}

func (m *Money) RemoveTax(percent float64) *Money {
	if percent == 0 {
		percent = float64(m.taxRate)
	}

	m.taxAmount -= percent

	m.amount = m.amount - percent
	return m
}

func (m *Money) All() MoneyValue {
	return MoneyValue{
		Amount: fmt.Sprintf("%.2f", m.amount),
		Tax:    fmt.Sprintf("%.2f", m.taxAmount),
	}
}

func (m *Money) GetTax() string {
	return fmt.Sprintf("%.2f", m.taxAmount)
}

func (m *Money) GetAmount() string {
	return fmt.Sprintf("%.2f", m.amount)
}

func NewMoney(amount float64) MoneyInterface {
	return &Money{
		amount: amount,
	}
}
