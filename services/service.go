package serviceSlice

import (
	"errors"
)

func GetSlice(slice *[]int) (int, int, error) {

	if slice != nil {
		return len(*slice), cap(*slice), nil
	}
	return 0, 0, errors.New("пустой слайс или что то еще")

}

func AppendSlices(slice *[]int, i []int) {
	for _, iSlice := range i {
		*slice = append(*slice, iSlice)
	}
}
