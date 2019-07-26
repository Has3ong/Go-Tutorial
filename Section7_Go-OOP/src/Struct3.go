package main

import "fmt"

type productList struct{ count, price int }

func (p productList) show() int {
	return p.count * p.price
}

func (p *productList) pointer_rePurchase(count, price int) {
	p.count += count
	p.price += price
}

func (p productList) rePurchase(count, price int) {
	p.count += count
	p.price += price
}

func main() {
	p1 := productList{5, 100}
	fmt.Println(p1.show())
	p1.pointer_rePurchase(3, 300)
	fmt.Println(p1.show())

	p2 := productList{5, 100}
	fmt.Println(p2.show())
	p2.rePurchase(3, 300)
	fmt.Println(p2.show())
}
