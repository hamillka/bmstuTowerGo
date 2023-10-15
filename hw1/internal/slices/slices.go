package slices

import "fmt"

func Sum(lst []int) int {
	s := 0
	for _, v := range lst {
		s += v
	}

	return s
}

func ReadList(lst []int) error {
	for i := range lst {
		_, err := fmt.Scanf("%d", &lst[i])
		if err != nil {
			return err
		}
	}

	return nil
}
