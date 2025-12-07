package serviceSlice

import (
	"errors"
)

func GetSlice(slice *[]int) (int, int, error) {

	if (len(*slice) == 0) && (cap(*slice) == 0) {
		return 0, 0, errors.New("пустой слайс")
	}
	if slice == nil {
		return 0, 0, errors.New("указатель на слайс равен nil")
	}
	return len(*slice), cap(*slice), nil

}

func AppendSlices(slice *[]int, i []int) {
	for _, iSlice := range i {
		*slice = append(*slice, iSlice)
	}
}
