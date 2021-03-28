package main

import (
	factoryModel "github.com/design-pattern/abstractfactory/model"
	factoryUsecase "github.com/design-pattern/abstractfactory/usecase"
)

func main() {
	newFactrory := factoryModel.NewFactory()

	srv := factoryUsecase.Service{
		FactoryMethod: newFactrory,
	}

	srv.CreateFactory()
}
