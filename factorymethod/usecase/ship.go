package usecase

import "log"

type ShipUseCase struct{}

func newShip() UseCaseService {
	return &ShipUseCase{}
}

func (s *ShipUseCase) PlanDelivery() error {
	log.Println("[process plan delivery by sea using ship]")
	return nil
}

func (s *ShipUseCase) CreateTransport() error {
	log.Println("[process create transport by sea using ship]")
	return nil
}
