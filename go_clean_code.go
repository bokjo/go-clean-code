package main

import (
	"fmt"

	ch1 "github.com/bokjo/go-clean-code/ch1-clean-code"
	ch2 "github.com/bokjo/go-clean-code/ch2-meaningful-names"
)

func main() {
	fmt.Printf("########## Clean Code examples in GO ##########\n\n")

	ch1.CleanCode()
	ch2.MeaningfulNames()

	fmt.Printf("\n##################### END #####################")
}
