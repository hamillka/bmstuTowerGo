package main

import (
	"fmt"
	slc "hw1/internal/slices"
	"log"
)

func main() {
	var n int
	_, err := fmt.Scanf("%d", &n)
	if err != nil {
		log.Fatal(err)
	}
	lst := make([]int, n)

	err = slc.ReadList(lst)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Сумма: %d\n", slc.Sum(lst))
}
