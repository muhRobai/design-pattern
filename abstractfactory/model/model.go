package model

type Furniture struct {
	classicFactory classicItem
	artFactory     ArtFactory
	modernFactory  ModernStruct
}

var FurnitureType = map[uint8]string{
	Modern:    "Modern",
	Victorial: "Victorial",
	ArtDeco:   "Art Deco",
}

const (
	Modern    = 1
	Victorial = 2
	ArtDeco   = 3
)

type FurnitureFactory interface {
	CreateChair() Ichair
	CreateCoffeTable() CoffeTable
	CreateSofa() Sofa
}

type FurnitureMethod interface {
	ClassicFactory() FurnitureFactory
	ArtFactory() FurnitureFactory
	ModernFactory() FurnitureFactory
}

func NewFactory() FurnitureMethod {

	modern := ModernStruct{}
	art := ArtFactory{}
	classic := classicItem{}

	return &Furniture{
		classicFactory: classic,
		artFactory:     art,
		modernFactory:  modern,
	}
}

func (f *Furniture) ArtFactory() FurnitureFactory {
	return &f.artFactory
}

func (f *Furniture) ClassicFactory() FurnitureFactory {
	return &f.classicFactory
}

func (f *Furniture) ModernFactory() FurnitureFactory {
	return &f.modernFactory
}
