package dto

type Pager struct {
	Page
	TotalRecord int64 `json:"total"`
	TotalPage   int   `json:"totalPage"`
}

type DataPager struct {
	Data interface{} `json:"data"`
	Pager
}

func NewDataPager(data interface{}, index, size int, total int64) *DataPager {
	res := &DataPager{}
	res.PageIndex = index
	res.PageSize = size
	res.TotalRecord = total
	res.Data = data
	res.Paginator()
	return res
}

func (p *Pager) Paginator() *Pager {
	p.TotalPage = getTotalPage(p.TotalRecord, p.GetSize())
	return p
}

func (p *DataPager) GetTotalRecord() int64 {
	return p.TotalRecord
}

func getTotalPage(total int64, pageSize int) int {
	if pageSize == 0 {
		return 1
	}
	if int(total)%pageSize == 0 {
		return int(total) / pageSize
	}
	return int(total)/pageSize + 1
}
