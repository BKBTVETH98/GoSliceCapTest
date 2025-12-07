package main

import (
	"fmt"
	serviceSlice "sliceCap/services"
)

// поведение емкости слайса при добавлении элементов превышающих текущую емкость
func main() {

	slice := &[]int{1, 2, 3, 4, 5} // указатель на слайс

	length, capacity, err := serviceSlice.GetSlice(slice) // получение длины и емкости слайса

	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Length: %d, Capacity: %d\n", length, capacity) // вывод длины и емкости слайса
	fmt.Println(*slice)                                        // вывод слайса

	serviceSlice.AppendSlices(slice, []int{6}) // добавление элемента в слайс

	length, capacity, _ = serviceSlice.GetSlice(slice) //	 получение длины и емкости слайса

	fmt.Printf("Length: %d, Capacity: %d\n", length, capacity)
	fmt.Println(*slice)

	serviceSlice.AppendSlices(slice, []int{7, 8, 9, 10, 11}) // добавление нескольких элементов в слайс, емкость умножится на 2

	length, capacity, _ = serviceSlice.GetSlice(slice)

	fmt.Printf("Length: %d, Capacity: %d\n", length, capacity)
	fmt.Println(*slice)

}
