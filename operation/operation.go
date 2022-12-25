package operation

import (
	"fmt"
	"strconv"
	"strings"
)

type bookRentingSystem struct {
	books     map[string][]*bookStore          //用于search
	book2shop map[string]map[string]*bookStore // book->shop->price ，用于借、还
	rented    []*bookStore                     //用于report
}
type bookStore struct {
	shop   string
	book   string
	price  string
	status bool
}

//初始化系统
func InitBookRentingSystem(n int, entries [][]string) (*bookRentingSystem, string) {
	var (
		initError   = "Init Error"
		initSuccess = "Init Success"
		brs         = new(bookRentingSystem)
		book2shop   = make(map[string]map[string]*bookStore)
		books       = make(map[string][]*bookStore)
	)
	//书店数、图书数量检查
	if (n < 1 || n > 3*105) || (len(entries) < 1 || len(entries) > 105) {
		return nil, initError
	}

	for _, e := range entries {
		b := new(bookStore)

		if err := checkMetaData(e, n); err != nil {
			fmt.Println(err)
			return nil, initError
		}

		b.shop, b.book, b.price, b.status = e[0], e[1], e[2], true

		if _, ok := book2shop[b.book]; ok {
			if _, ok1 := book2shop[b.book][b.shop]; ok1 {
				return nil, initError
			}
			book2shop[b.book][b.shop] = b
		} else {
			book2shop[b.book] = map[string]*bookStore{b.shop: b}
		}

		books = addBook(b, books)
	}

	brs.books, brs.book2shop = books, book2shop

	return brs, initSuccess
}

//借阅
func (brs *bookRentingSystem) Rent(shop, book string) string {
	var (
		fail    = "Rent Error"
		success = "Rent Success"
	)

	if shops, ok := brs.book2shop[book]; ok {
		if bs, ok := shops[shop]; ok {
			if bs.status {
				bs.status = false
				index := getSortIndex(bs, brs.rented)
				temp := make([]*bookStore, len(brs.rented))
				copy(temp, brs.rented)

				brs.rented = append(append(brs.rented[:index], bs), temp[index:]...)
				return success
			}
		}
	}

	return fail
}

//归还
func (brs *bookRentingSystem) Drop(shop, book string) string {

	var (
		fail    = "Drop Error"
		success = "Drop Success"
	)

	if shops, ok := brs.book2shop[book]; ok {
		if bs, ok := shops[shop]; ok {
			if !bs.status {
				bs.status = true

				for k, v := range brs.rented {
					if v.shop == shop && v.book == book {
						brs.rented = append(brs.rented[:k], brs.rented[k+1:]...)
						break
					}
				}
				return success
			}
		}
	}

	return fail
}

//检索
func (brs *bookRentingSystem) Search(book string) [][]string {
	var cheapestShops = make([][]string, 0)

	for _, v := range brs.books[book] {
		if v.status {
			cheapestShops = append(cheapestShops, []string{v.shop, v.book, v.price})
		}
		if len(cheapestShops) >= 5 {
			return cheapestShops
		}
	}

	return cheapestShops
}

//已借出书籍
func (brs *bookRentingSystem) Report() [][]string {
	var cheapestBooks = make([][]string, 0)

	for _, v := range brs.rented {
		if len(cheapestBooks) < 5 {
			cheapestBooks = append(cheapestBooks, []string{v.shop, v.book})
		}
	}

	return cheapestBooks
}

func checkMetaData(e []string, n int) error {
	//shopid检查
	split := strings.Split(e[0], "_")
	shopId, err := strconv.Atoi(split[1])
	if err != nil {
		return fmt.Errorf("shopid [%v] invalid", e[0])
	}
	if shopId < 0 || shopId >= n {
		return fmt.Errorf("shopid [%v] invalid", e[0])
	}

	//bookid检查
	split = strings.Split(e[1], "_")
	bookId, err := strconv.Atoi(split[1])
	if err != nil {
		return fmt.Errorf("bookid [%v] invalid", e[1])
	}
	if bookId < 1 || bookId > 104 {
		return fmt.Errorf("bookid [%v] invalid", e[1])
	}

	//price检查
	price, err := strconv.ParseFloat(e[2], 64)
	if err != nil {
		return fmt.Errorf("price [%v] invalid", e[2])
	}
	if price < 1 || price > 104 {
		return fmt.Errorf("price [%v] invalid", e[2])
	}
	return nil
}

func addBook(bs *bookStore, books map[string][]*bookStore) map[string][]*bookStore {

	b, ok := books[bs.book]
	if !ok {
		books[bs.book] = []*bookStore{bs}
		return books
	}
	temp := make([]*bookStore, len(b))
	index := getSortIndex(bs, b)
	copy(temp, b)
	b2 := append(append(b[:index], bs), temp[index:]...)
	books[bs.book] = b2

	return books
}

func getSortIndex(bs *bookStore, bss []*bookStore) int {
	for k, v := range bss {
		out := strings.Compare(bs.price, v.price)
		if out == -1 {
			return k
		}
		if out == 0 {
			out1 := strings.Compare(bs.shop, v.shop)
			if out1 == -1 {
				return k
			}
			if out1 == 0 {
				out2 := strings.Compare(bs.book, v.book)
				if out2 == -1 {
					return k
				}
			}
		}
	}

	return len(bss)
}
