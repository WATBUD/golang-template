// // * and & of pointer symbols in go language
package main

import "fmt"

//go語言指針符號的*和&
// &符號的意思是對變量取地址，如：變量a的地址是&a
// *符號的意思是對指針取值，如: *&a，就是a變量所在地址的值，當然也就是a的值了

func StartPointerSymbols() {
	var a int = 1
	var b *int = &a
	var c **int = &b
	var x int = *b
	fmt.Println("a = ", a)
	fmt.Println("&a = ", &a)
	fmt.Println("*&a = ", *&a)

	fmt.Println("b = ", b)
	fmt.Println("&b = ", &b)
	fmt.Println("*&b = ", *&b)
	fmt.Println("*b = ", *b)
	fmt.Println("c = ", c)
	fmt.Println("*c = ", *c)
	fmt.Println("&c = ", &c)
	fmt.Println("*&c = ", *&c)
	fmt.Println("**c = ", **c)
	fmt.Println("***&*&*&*&c = ", ***&*&*&*&*&c)
	fmt.Println("x = ", x)
}
func PointerTest() {
	a := 1222222222222222
	p := &a                //取址&
	fmt.Printf("%d\n", *p) //取值*
	//fmt.Printf("%d\n", "TTTTTTTTTTTTTTT") //取值*

}

//程序(進程)、執行緒(線程)、協程
