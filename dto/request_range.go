package dto

// IRange range
type IRange interface {
	GetBeg() int64
	GetEnd() int64
}

var _ = IRange(&Range{})

// Range range
type Range struct {
	Beg int64 `json:"beg" form:"beg"`
	End int64 `json:"end" form:"end"`
}

func (i *Range) GetBeg() int64 {
	return i.Beg
}
func (i *Range) GetEnd() int64 {
	return i.End
}
