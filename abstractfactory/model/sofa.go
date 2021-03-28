package model

type SofaItem struct {
	length int
	width  int
}

type Sofa interface {
	SetLength(key int)
	SetWidth(key int)
	GetLength() int
	GetWitdth() int
}

func (s *SofaItem) SetLength(key int) {
	s.length = key
}

func (s *SofaItem) SetWidth(key int) {
	s.width = key
}

func (s *SofaItem) GetLength() int {
	return s.length
}

func (s *SofaItem) GetWitdth() int {
	return s.width
}
