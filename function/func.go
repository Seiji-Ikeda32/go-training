package main

import "fmt"

func addition(x, y int) int {
	return x + y
}

//無名関数+関数を返す関数
func ReturnFunc() func() {
	return func() {
		fmt.Println("awesome Go!!")
	}
}

//関数を引数にとる関数
func CallFunc(f func()) {
	f()
}

func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

//クロージャーの実装
func Later() func(string) string {
	var store string
	return func(next string) string {
		//storeが最初空文字でreturnでsを返しているから一つ前の値を返すことになる。
		s := store
		store = next
		return s
	}
}

//クロージャーを利用したジェネレーターの実装
func intejers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	i := addition(1, 4)
	fmt.Println(i)

	i2, i3 := Div(9, 3)
	fmt.Println(i2, i3)

	//片方の返り値のみ受け取りたい場合はアンダースコア
	i4, _ := Div(9, 4)
	fmt.Println(i4)

	//無名関数+関数を返す関数
	f := ReturnFunc()
	f()

	//関数を引数にとる関数
	CallFunc(func() {
		fmt.Println("awesome Go!!")
	})

	//クロージャーの実装
	f2 := Later()
	fmt.Println(f2("hello!"))
	fmt.Println(f2("go"))
	fmt.Println(f2("is"))
	fmt.Println(f2("cool"))

	//クロージャーを利用したジェネレーターの実装
	ints := intejers()

	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())

	//違う変数で関数を呼び出すときは数がリセットされる
	otherInts := intejers()
	fmt.Println(otherInts())
	fmt.Println(otherInts())
	fmt.Println(otherInts())
	fmt.Println(otherInts())
}
