package model

type classicItem struct{}

func (c *classicItem) CreateChair() Ichair {
	return &chair{
		length: 12,
		count:  12,
	}
}

func (c *classicItem) CreateCoffeTable() CoffeTable {
	return &Table{
		count:  12,
		height: 12,
	}
}

func (c *classicItem) CreateSofa() Sofa {
	return &SofaItem{
		width:  12,
		length: 12,
	}
}
