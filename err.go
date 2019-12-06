package main

import (
	"errors"
	"fmt"
)

//Myfunc 自定义函数
func Myfunc(a, b int) (result float32, err error) {
	err = nil
	if b == 0 {
		err = errors.New("分母不能为0")
	} else {
		result = float32(a) / float32(b)
	}
	return
}

func a() {
	fmt.Println("aaaaaa")
}
func b() {
	// fmt.Println("bbbbbb")
	// 显示调用panic触发宕机
	panic("this is a panic error....")
}
func c(i int) int {
	// 数组越界导致panic的产生，程序崩溃（内部产生panic）
	fmt.Println("cccccccc")
	arrC := [5]int{1, 2, 3, 4, 5}
	result := arrC[i]
	return result
}

func d(i int) int {
	// 设置recover
	defer func() {
		// recover() // 即使发生错误，recover后依然不影响后面的执行
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	// 数组越界导致panic的产生，程序崩溃（内部产生panic）
	arrC := [5]int{1, 2, 3, 4, 5}
	result := arrC[i]
	return result
}

func main() {

	// 1
	// err1 := fmt.Errorf("%s", "this is a normal error1")
	// fmt.Println("err1 = ", err1)
	// err2 := errors.New("this is a normal error2")
	// fmt.Println("err2 = ", err2)

	// result, err3 := Myfunc(1, 2)
	// if err3 != nil {
	// 	fmt.Println("err3 = ", err3)
	// }

	// fmt.Printf("result：%f\n", result)

	// 2
	// a()

	// 3、显示调用panic触发宕机
	// b()

	// 4、数组越界、空指针等错误
	// c(6)

	// 5、运行时一旦panic产生就会导致程序崩溃
	// GO语言为我们提供了专用于拦截运行时panic的内建函数-recover
	// recover会使程序从panic中恢复，并返回panic value。导致panic
	// 异常的函数不会继续运行，但能正常返回。
	// 在未发生panic时调用recover，recover会返回nil
	d(8)
	a()
}
