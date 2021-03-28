package model

type ArtFactory struct{}

func (a *ArtFactory) CreateChair() Ichair {
	return &chair{
		length: 11,
		count:  11,
	}
}

func (a *ArtFactory) CreateCoffeTable() CoffeTable {
	return &Table{
		count:  11,
		height: 11,
	}
}

func (a *ArtFactory) CreateSofa() Sofa {
	return &SofaItem{
		width:  11,
		length: 11,
	}
}
