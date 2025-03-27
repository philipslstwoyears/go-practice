package main

import (
	"fmt"
	"math/big"
)

type Math struct {
	a *big.Int
	b *big.Int
}

func (m *Math) Sum() *big.Int {
	var res big.Int
	return res.Add(m.a, m.b)
}

func (m *Math) Sub() *big.Int {
	var res big.Int
	return res.Sub(m.a, m.b)
}

func (m *Math) Mul() *big.Int {
	var res big.Int
	return res.Mul(m.a, m.b)
}

func (m *Math) Div() *big.Int {
	var res big.Int
	return res.Div(m.a, m.b)
}

func main() {
	var a, b big.Int
	fmt.Scan(&a, &b)
	m := Math{a: &a, b: &b}
	fmt.Println(m.Sum())
}
