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

type Service struct {
	Ship  *ShipUseCase
	Truck *TruckUseCase
}

func NewService() *Service {
	ship := ShipUseCase{}
	truck := TruckUseCase{}
	return &Service{
		Ship:  &ship,
		Truck: &truck,
	}
}

func (h *Service) ProcessDelivery(ctx context.Context, req *model.DeliveryProcess) error {
	delivery := h.GetDeliveryType(req.Type)
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

func (h *Service) GetDeliveryType(key string) UseCaseService {
	if constant.GetDelivery(key) == constant.TruckDelivery {
		if h.Ship != nil {
			return h.Ship
		}
	}

	if constant.GetDelivery(key) == constant.ShipDelivery {
		if h.Truck != nil {
			return h.Truck
		}
	}

	return nil
}
