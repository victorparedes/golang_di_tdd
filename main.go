package main

import (
	"fmt"
	"os"

	"github.com/victorparedes/golang_di_tdd/services"
)

func main() {
	isbn := os.Args[1]
	myServices := services.NewService()

	book, err := myServices.SearchByIsbn(isbn)
	if err != nil {
		fmt.Print(err.Error())
		return
	}

	fmt.Printf("%+v\n", book)
}
