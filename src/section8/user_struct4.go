package main

import "fmt"

type ShoppingBasket struct{
  count int
  price int
}

//결제
func (b ShoppingBasket) purchase() int{
  return b.count * b.price
}

//참조 전달
func (b *ShoppingBasket) rePurcheasP(count, price int) {
  b.count = count
  b.price = price
}

//값 전달
func (b ShoppingBasket) rePurcheasD(count, price int) {
  b.count = count
  b.price = price
}

func main(){

  bs1 := ShoppingBasket{
    count : 3,
    price : 5000,
  }

  // bs2 := ShoppingBasket{5, 2000}

  fmt.Println("ex1-1 : ", bs1.purchase())
  bs1.rePurcheasP(7, 5000)
  fmt.Println("ex1-2 : ", bs1.purchase())
  bs1.rePurcheasD(3, 5000)
  fmt.Println("ex1-3 : ", bs1.purchase())

}
