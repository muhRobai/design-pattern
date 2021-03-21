package usecase

import "log"

type TruckUseCase struct{}

func newTruct() UseCaseService {
	return &TruckUseCase{}
}

func (c *TruckUseCase) PlanDelivery() error {
	log.Println("[process plan delivery by land using truck]")
	return nil
}

func (c *TruckUseCase) CreateTransport() error {
	log.Println("[process create transportation by land using truck]")
	return nil
}
