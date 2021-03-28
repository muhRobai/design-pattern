package model

type Table struct {
	count  int
	height int
}

type CoffeTable interface {
	SetCount(key int)
	SetHeight(key int)
	GetCount() int
	GetHeight() int
}

func (t *Table) SetCount(key int) {
	t.count = key
}

func (t *Table) SetHeight(key int) {
	t.height = key
}

func (t *Table) GetCount() int {
	return t.count
}

func (t *Table) GetHeight() int {
	return t.height
}
