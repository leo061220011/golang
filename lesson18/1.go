// написать функцию, которая выводит n-ый элемент массива.
// Если n>len(arr)-1
// вернуть ошибку
// Если ошибки нет: вывести этот элемент
// func main(){
// 	arr := []int{1,2,3}
// 	i, err := myFunc(arr, 3) //"я ошибка"

package main

import (
	"errors"
	"fmt"
)

func myFunc(arr []int, n int) (int, error) {
	if n > len(arr)-1 {
		return 0, errors.New("n > размера слайса")
	}
	return arr[n], nil
}

func main() {

	arr := []int{1, 2, 3}
	i, err := myFunc(arr, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}

}
