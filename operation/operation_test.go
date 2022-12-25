package operation

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInitBookRentingSystem(t *testing.T) {
	var err string
	var initFailed = "Init Error"
	var initSuccess = "Init Success"
	/*
		entries := [][]string{{"shop_0", "book_1", "5"}, {"shop_0", "book_2", "6"}, {"shop_0", "book_3", "7"}, {"shop_1", "book_1", "4"},
				{"shop_1", "book_2", "7"}, {"shop_2", "book_1", "5"}}
		**/

	t.Run("Given_书店数量=0_When_初始化系统_Then_初始化失败", func(t *testing.T) {
		n := 0
		entries := [][]string{{"shop_0", "book_1", "5"}, {"shop_0", "book_2", "6"}, {"shop_0", "book_3", "7"}, {"shop_1", "book_1", "4"},
			{"shop_1", "book_2", "7"}, {"shop_2", "book_1", "5"}}
		_, err = InitBookRentingSystem(n, entries)
		if err != initFailed {
			t.Errorf("expected:%v,got:%v", initFailed, err)
		}
	})

	t.Run("Given_书籍数量=0_When_初始化系统_Then_初始化失败", func(t *testing.T) {
		n := 3
		var entries [][]string
		_, err = InitBookRentingSystem(n, entries)
		if err != initFailed {
			t.Errorf("expected:%v,got:%v", initFailed, err)
		}
	})

	t.Run("Given_shopid=4_When_初始化系统_Then_初始化失败", func(t *testing.T) {
		n := 3
		entries := [][]string{{"shop_3", "book_1", "5"}, {"shop_0", "book_2", "6"}, {"shop_0", "book_3", "7"}, {"shop_1", "book_1", "4"},
			{"shop_1", "book_2", "7"}, {"shop_2", "book_1", "5"}}

		_, err = InitBookRentingSystem(n, entries)
		if err != initFailed {
			t.Errorf("expected:%v,got:%v", initFailed, err)
		}
	})

	t.Run("Given_shopid=D_When_初始化系统_Then_初始化失败", func(t *testing.T) {
		n := 3
		entries := [][]string{{"shop_D", "book_1", "5"}, {"shop_0", "book_2", "6"}, {"shop_0", "book_3", "7"}, {"shop_1", "book_1", "4"},
			{"shop_1", "book_2", "7"}, {"shop_2", "book_1", "5"}}

		_, err = InitBookRentingSystem(n, entries)
		if err != initFailed {
			t.Errorf("expected:%v,got:%v", initFailed, err)
		}
	})

	t.Run("Given_bookid=105_When_初始化系统_Then_初始化失败", func(t *testing.T) {
		n := 3
		entries := [][]string{{"shop_0", "book_105", "5"}, {"shop_0", "book_2", "6"}, {"shop_0", "book_3", "7"}, {"shop_1", "book_1", "4"},
			{"shop_1", "book_2", "7"}, {"shop_2", "book_1", "5"}}

		_, err = InitBookRentingSystem(n, entries)
		if err != initFailed {
			t.Errorf("expected:%v,got:%v", initFailed, err)
		}
	})

	t.Run("Given_bookid=D_When_初始化系统_Then_初始化失败", func(t *testing.T) {
		n := 3
		entries := [][]string{{"shop_0", "book_D", "5"}, {"shop_0", "book_2", "6"}, {"shop_0", "book_3", "7"}, {"shop_1", "book_1", "4"},
			{"shop_1", "book_2", "7"}, {"shop_2", "book_1", "5"}}

		_, err = InitBookRentingSystem(n, entries)
		if err != initFailed {
			t.Errorf("expected:%v,got:%v", initFailed, err)
		}
	})

	t.Run("Given_price=105_When_初始化系统_Then_初始化失败", func(t *testing.T) {
		n := 3
		entries := [][]string{{"shop_0", "book_1", "105"}, {"shop_0", "book_2", "6"}, {"shop_0", "book_3", "7"}, {"shop_1", "book_1", "4"},
			{"shop_1", "book_2", "7"}, {"shop_2", "book_1", "5"}}

		_, err = InitBookRentingSystem(n, entries)
		if err != initFailed {
			t.Errorf("expected:%v,got:%v", initFailed, err)
		}
	})

	t.Run("Given_price=D_When_初始化系统_Then_初始化失败", func(t *testing.T) {
		n := 3
		entries := [][]string{{"shop_0", "book_1", "D"}, {"shop_0", "book_2", "6"}, {"shop_0", "book_3", "7"}, {"shop_1", "book_1", "4"},
			{"shop_1", "book_2", "7"}, {"shop_2", "book_1", "5"}}

		_, err = InitBookRentingSystem(n, entries)
		if err != initFailed {
			t.Errorf("expected:%v,got:%v", initFailed, err)
		}
	})

	t.Run("Given_{shop_0,book_1,5}*2_When_初始化系统_Then_初始化失败", func(t *testing.T) {
		n := 3
		entries := [][]string{{"shop_0", "book_1", "5"}, {"shop_0", "book_2", "6"}, {"shop_0", "book_1", "5"}, {"shop_1", "book_1", "4"},
			{"shop_1", "book_2", "7"}, {"shop_2", "book_1", "5"}}

		_, err = InitBookRentingSystem(n, entries)
		if err != initFailed {
			t.Errorf("expected:%v,got:%v", initFailed, err)
		}
	})

	t.Run("Given_满足规则的参数_When_初始化系统_Then_初始化成功", func(t *testing.T) {
		n := 3
		entries := [][]string{{"shop_0", "book_1", "5"}, {"shop_0", "book_2", "6"}, {"shop_0", "book_3", "7"}, {"shop_1", "book_1", "4"},
			{"shop_1", "book_2", "7"}, {"shop_2", "book_1", "5"}}

		_, err = InitBookRentingSystem(n, entries)
		if err != initSuccess {
			t.Errorf("expected:%v,got:%v", initSuccess, err)
		}
	})
}

func TestBookRentingSystem_Rent(t *testing.T) {
	var (
		fail        = "Rent Error"
		success     = "Rent Success"
		dropSuccess = "Drop Success"
	)

	t.Run("Given_shop1存在书籍book2_When_Rent(shop_1,book_2)_Then_借阅成功", func(t *testing.T) {
		n := 3
		entries := [][]string{{"shop_0", "book_1", "5"}, {"shop_0", "book_2", "6"}, {"shop_0", "book_3", "7"}, {"shop_1", "book_1", "4"},
			{"shop_1", "book_2", "7"}, {"shop_2", "book_1", "5"}}
		brs, _ := InitBookRentingSystem(n, entries)
		ret := brs.Rent("shop_1", "book_2")
		if ret != success {
			t.Errorf("expected:%v,got:%v", success, ret)
		}
	})

	t.Run("Given_book2被借阅->归还_When_Rent(shop_1,book_2)_Then_借阅成功", func(t *testing.T) {
		n := 3
		entries := [][]string{{"shop_0", "book_1", "5"}, {"shop_0", "book_2", "6"}, {"shop_0", "book_3", "7"}, {"shop_1", "book_1", "4"},
			{"shop_1", "book_2", "7"}, {"shop_2", "book_1", "5"}}
		brs, _ := InitBookRentingSystem(n, entries)
		ret := brs.Rent("shop_1", "book_2")
		if ret != success {
			t.Errorf("expected:%v,got:%v", success, ret)
		}
		ret = brs.Drop("shop_1", "book_2")
		if ret != dropSuccess {
			t.Errorf("expected:%v,got:%v", dropSuccess, ret)
		}
		ret = brs.Rent("shop_1", "book_2")
		if ret != success {
			t.Errorf("expected:%v,got:%v", success, ret)
		}
	})

	t.Run("Given_shop1不存在书籍book3_When_Rent(shop_1,book_3)_Then_借阅失败", func(t *testing.T) {
		n := 3
		entries := [][]string{{"shop_0", "book_1", "5"}, {"shop_0", "book_2", "6"}, {"shop_0", "book_3", "7"}, {"shop_1", "book_1", "4"},
			{"shop_1", "book_2", "7"}, {"shop_2", "book_1", "5"}}
		brs, _ := InitBookRentingSystem(n, entries)
		ret := brs.Rent("shop_1", "book_3")
		if ret != fail {
			t.Errorf("expected:%v,got:%v", fail, ret)
		}
	})

	t.Run("Given_shop1书籍book2已经被借出_When_Rent(shop_1,book_2)_Then_借阅失败", func(t *testing.T) {
		n := 3
		entries := [][]string{{"shop_0", "book_1", "5"}, {"shop_0", "book_2", "6"}, {"shop_0", "book_3", "7"}, {"shop_1", "book_1", "4"},
			{"shop_1", "book_2", "7"}, {"shop_2", "book_1", "5"}}
		brs, _ := InitBookRentingSystem(n, entries)
		ret := brs.Rent("shop_1", "book_2")
		if ret != success {
			t.Errorf("expected:%v,got:%v", success, ret)
		}
		ret = brs.Rent("shop_1", "book_2")
		if ret != fail {
			t.Errorf("expected:%v,got:%v", fail, ret)
		}
	})
}

func TestBookRentingSystem_Drop(t *testing.T) {
	var (
		fail        = "Drop Error"
		success     = "Drop Success"
		rentSuccess = "Rent Success"
	)

	t.Run("Given_shop0借出书籍book3_When_Drop(shop_0,book_3)_Then_归还成功", func(t *testing.T) {
		n := 3
		entries := [][]string{{"shop_0", "book_1", "5"}, {"shop_0", "book_2", "6"}, {"shop_0", "book_3", "7"}, {"shop_1", "book_1", "4"},
			{"shop_1", "book_2", "7"}, {"shop_2", "book_1", "5"}}
		brs, _ := InitBookRentingSystem(n, entries)
		ret := brs.Rent("shop_0", "book_3")
		if ret != rentSuccess {
			t.Errorf("expected:%v,got:%v", rentSuccess, ret)
		}
		ret = brs.Drop("shop_0", "book_3")
		if ret != success {
			t.Errorf("expected:%v,got:%v", success, ret)
		}
	})

	t.Run("Given_shop0未借出书籍book3_When_Drop(shop_0,book_3)_Then_归还的失败", func(t *testing.T) {
		n := 3
		entries := [][]string{{"shop_0", "book_1", "5"}, {"shop_0", "book_2", "6"}, {"shop_0", "book_3", "7"}, {"shop_1", "book_1", "4"},
			{"shop_1", "book_2", "7"}, {"shop_2", "book_1", "5"}}
		brs, _ := InitBookRentingSystem(n, entries)
		ret := brs.Drop("shop_0", "book_3")
		if ret != fail {
			t.Errorf("expected:%v,got:%v", fail, ret)
		}
	})

}

func TestBookRentingSystem_Search(t *testing.T) {
	t.Run("Given_book1有副本6个_When_Search(book_1)_Then_返回最便宜的五个副本", func(t *testing.T) {
		expected := [][]string{{"shop_1", "book_1", "4"}, {"shop_5", "book_1", "4"}, {"shop_0", "book_1", "5"},
			{"shop_2", "book_1", "5"}, {"shop_3", "book_1", "5"}}
		n := 6
		entries := [][]string{{"shop_0", "book_1", "5"}, {"shop_0", "book_2", "6"}, {"shop_0", "book_3", "7"}, {"shop_1", "book_1", "4"},
			{"shop_1", "book_2", "7"}, {"shop_2", "book_1", "5"}, {"shop_3", "book_1", "5"}, {"shop_4", "book_1", "5"}, {"shop_5", "book_1", "4"}}
		brs, _ := InitBookRentingSystem(n, entries)

		got := brs.Search("book_1")
		if !reflect.DeepEqual(expected, got) {
			t.Errorf("expected:%v,got:%v", expected, got)
		}
	})

	t.Run("Given_book4有副本0个_When_Search(book_4)_Then_返回空", func(t *testing.T) {
		expected := make([][]string, 0)
		n := 6
		entries := [][]string{{"shop_0", "book_1", "5"}, {"shop_0", "book_2", "6"}, {"shop_0", "book_3", "7"}, {"shop_1", "book_1", "4"},
			{"shop_1", "book_2", "7"}, {"shop_2", "book_1", "5"}, {"shop_3", "book_1", "5"}, {"shop_4", "book_1", "5"}, {"shop_5", "book_1", "4"}}
		brs, _ := InitBookRentingSystem(n, entries)

		got := brs.Search("book_4")
		fmt.Println(got)
		if !reflect.DeepEqual(expected, got) {
			t.Errorf("expected:%v,got:%v", expected, got)
		}
	})
}

func TestBookRentingSystem_Report(t *testing.T) {
	t.Run("", func(t *testing.T) {
		expected := [][]string{{"shop_0", "book_1"}, {"shop_0", "book_3"}, {"shop_3", "book_1"},
			{"shop_4", "book_1"}}
		n := 6
		entries := [][]string{{"shop_0", "book_1", "5"}, {"shop_0", "book_2", "6"}, {"shop_0", "book_3", "5"}, {"shop_1", "book_1", "4"},
			{"shop_1", "book_2", "7"}, {"shop_2", "book_1", "5"}, {"shop_3", "book_1", "5"}, {"shop_4", "book_1", "5"}, {"shop_5", "book_1", "4"}}
		brs, _ := InitBookRentingSystem(n, entries)
		brs.Rent("shop_3", "book_1")
		brs.Rent("shop_0", "book_3")
		brs.Rent("shop_0", "book_1")
		brs.Rent("shop_4", "book_1")
		brs.Rent("shop_1", "book_2")
		brs.Drop("shop_1", "book_2")
		got:=brs.Report()
		fmt.Println(got)
		if !reflect.DeepEqual(expected, got) {
			t.Errorf("expected:%v,got:%v", expected, got)
		}

	})
}
