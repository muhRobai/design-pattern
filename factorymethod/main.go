package main

import (
	"context"
	"log"

	"github.com/design-pattern/factorymethod/model"
	"github.com/design-pattern/factorymethod/usecase"
)

func main() {
	srv := usecase.NewService()

	err := srv.ProcessDelivery(context.Background(), &model.DeliveryProcess{
		Type: "SHIP",
	})

	if err != nil {
		log.Println(err)
	}

	err = srv.ProcessDelivery(context.Background(), &model.DeliveryProcess{
		Type: "TRUCK",
	})

	if err != nil {
		log.Println(err)
	}
}
