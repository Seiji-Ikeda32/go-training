package main

import (
	"fmt"
	"strconv"
	"time"
)

func anything(a interface{}) {
	fmt.Println(a)
}

//並行処理（サブ関数）
func sub() {
	for {
		fmt.Println("sub loop")
		time.Sleep(100 * time.Microsecond)
	}
}

func main() {
	//エラーハンドリング
	var s string = "A"

	//エラーの引数を破棄
	/*
		i, _ := strconv.Atoi(s)
		fmt.Printf("i = %T\n", i)
	*/

	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("i = %T\n", i)

	//for文 range
	arr := []string{"Go", "Ruby", "Python"}
	for i, v := range arr {
		fmt.Println(i, v)
	}

	//型アサーション
	var x interface{} = 5
	i2 := x.(int)
	fmt.Println(i2)
	fmt.Println(i2 + 3)

	anything(1)

	//並行処理（サブ関数）
	go sub()
	//並行処理（メイン関数）
	for {
		fmt.Println("main loop")
		time.Sleep(200 * time.Microsecond)
	}
}
