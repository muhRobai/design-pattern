package main

import (
	"context"

	"github.com/design-pattern/factorymethod/model"
	"github.com/design-pattern/factorymethod/usecase"
)

func main() {
	service := usecase.Service{}

	err := service.ProcessDelivery(context.Background(), &model.DeliveryProcess{
		Type: "SHIP",
	})

	if err != nil {
		return
	}

	err = service.ProcessDelivery(context.Background(), &model.DeliveryProcess{
		Type: "TRUCK",
	})

	if err != nil {
		return
	}
}
