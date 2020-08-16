package main

import "fmt"

type totalCost func(int, int) int

func describe(cnt int, price int, fn totalCost){
  fmt.Printf("ex1 : cnt - %d, price %d, orderPrice - %d",cnt, price, fn(cnt, price))
}

func main(){

  var orderPrice totalCost
  orderPrice = func(cnt int, price int) int{
    return (cnt * price) + 100000
  }

  describe(1000, 3, orderPrice)

}
