package model

type chair struct {
	length int
	count  int
}

type Ichair interface {
	SetLength(key int)
	SetCount(key int)
	GetLength() int
	GetCount() int
}

func (c *chair) SetLength(key int) {
	c.length = key
}

func (c *chair) SetCount(key int) {
	c.count = key
}

func (c *chair) GetLength() int {
	return c.length
}

func (c *chair) GetCount() int {
	return c.count
}
