package usecase

import (
	"log"

	factoryModel "github.com/design-pattern/abstractfactory/model"
)

type Service struct {
	FactoryMethod factoryModel.FurnitureMethod
}

func (c *Service) CreateFactory() error {
	classicFacory := c.FactoryMethod.ClassicFactory().CreateChair()

	log.Println(classicFacory.GetCount())
	log.Println(classicFacory.GetLength())

	return nil
}
