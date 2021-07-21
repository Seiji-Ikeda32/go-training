package main

import (
	"fmt"
	"time"
)

//定数
const Name = "Seiji Ikeda"

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(time.Now())

	//明示的な定義
	//var 変数名 型 = 値
	var i int = 100
	fmt.Println(i)

	var s string = "Hello Go"
	fmt.Println(s)

	//暗黙的な定義
	//(最初の値で型を設定するから後から型を変えられない。　関数外で使えない。)
	//変数名 := 値
	i2 := 200
	fmt.Println(i2)

	//型を表示
	fmt.Printf("%T\n", i2)

	//配列型
	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1) //=> [3]int  ここまでで型であるため、要素数の変更はできない。

	var arr2 [3]string = [3]string{"a", "b"}
	fmt.Println(arr2)

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	arr4 := [...]int{1, 2}
	fmt.Printf("%T\n", arr4) //=> [2]int  ...は自動で配列の個数で型を定義してくれる。

	//要素数の確認
	fmt.Println(len(arr4))
}
