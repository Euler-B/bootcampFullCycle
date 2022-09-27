package main

import (
	"database/sql"

	"github.com/Euler-B/bootcampFullCycle/internal/order/infra/db"
	"github.com/Euler-B/bootcampFullCycle/internal/order/usecase"
)

func main() {
	dbi, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/orders")
	if err != nil {
		panic(err)
	}
	defer dbi.Close()
	repository := db.NewOrderRepository(dbi)
	uc := usecase.NewCalculateFinalPriceUseCase(repository)
	input := usecase.OrderInputDTO{
		ID: "1234",
		Price: 100,
		Tax: 10,
	}
	ouput, err := uc.Execute(input)
	if err != nil {
		panic(err)
	}
	println(ouput.FinalPrice)
}

