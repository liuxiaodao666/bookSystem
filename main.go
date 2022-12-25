package main

import (
	"fmt"
	"github.com/booksystem/operation"
)

func main() {

	n := 3
	entries := [][]string{{"shop_0", "book_1", "5"}, {"shop_0", "book_2", "6"}, {"shop_0", "book_3", "7"}, {"shop_1", "book_1", "4"},
		{"shop_1", "book_2", "7"}, {"shop_2", "book_1", "5"}}

	var rentingSystem, err = operation.InitBookRentingSystem(n, entries)
	if err != "Init Success" {
		fmt.Println(err)
	}

	fmt.Println(rentingSystem.Search("book_1"))
	fmt.Println(rentingSystem.Search("book_0"))
	fmt.Println(rentingSystem.Rent("shop_1", "book_1"))
	fmt.Println(rentingSystem.Rent("shop_1", "book_1"))
	fmt.Println(rentingSystem.Drop("shop_0", "book_1"))
	fmt.Println(rentingSystem.Drop("shop_1", "book_1"))
	fmt.Println(rentingSystem.Report())
	fmt.Println(rentingSystem.Rent("shop_1", "book_1"))
	fmt.Println(rentingSystem.Rent("shop_1", "book_2"))
	fmt.Println(rentingSystem.Report())

}
