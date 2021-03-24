package main

import "fmt"

func main() {
	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1)

	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	// 要素の数を自動でカウントする
	arr4 := [...]string{"C", "D"}
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4)

	fmt.Println(arr2[0])
	fmt.Println(arr2[1])
	fmt.Println(arr2[2])
	fmt.Println(arr2[2-1])

	arr2[2] = "C"
	fmt.Println(arr2)

	// 要素数が異なると別の型と見なされる
	// var arr5 [4]string
	// arr5 = arr1
	// fmt.Println(arr5)

	// 要素数を調べる
	fmt.Println(len(arr1))
}
