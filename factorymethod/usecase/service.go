package usecase

import (
	"context"
	"errors"
	"log"

	"github.com/design-pattern/factorymethod/constant"
	"github.com/design-pattern/factorymethod/model"
)

type UseCaseService interface {
	PlanDelivery() error
	CreateTransport() error
}

type Service struct{}

func (h *Service) ProcessDelivery(ctx context.Context, req *model.DeliveryProcess) error {
	delivery := GetDeliveryType(req.Type)
	if delivery == nil {
		log.Println("[error missing delivery]")
		return errors.New("missing delivery method")
	}

	err := delivery.PlanDelivery()
	if err != nil {
		log.Println("[error process plan devlivery by", req.Type, err)
		return err
	}

	err = delivery.CreateTransport()
	if err != nil {
		log.Println("[error create transport by", req.Type, err)
		return err
	}

	return nil
}

func GetDeliveryType(key string) UseCaseService {
	if constant.GetDelivery(key) == constant.TruckDelivery {
		return newTruct()
	}

	if constant.GetDelivery(key) == constant.ShipDelivery {
		return newShip()
	}

	return nil
}
