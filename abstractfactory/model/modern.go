package model

type ModernStruct struct{}

func (m *ModernStruct) CreateChair() Ichair {
	return &chair{
		length: 10,
		count:  10,
	}
}

func (m *ModernStruct) CreateCoffeTable() CoffeTable {
	return &Table{
		count:  10,
		height: 10,
	}
}

func (m *ModernStruct) CreateSofa() Sofa {
	return &SofaItem{
		width:  10,
		length: 10,
	}
}
