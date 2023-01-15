package Practice

import (
	"fmt"
)

func PointerTest() {
	a := 1222222222222222
	p := &a                //取址&
	fmt.Printf("%d\n", *p) //取值*
	//fmt.Printf("%d\n", "TTTTTTTTTTTTTTT") //取值*

}
//程序(進程)、執行緒(線程)、協程