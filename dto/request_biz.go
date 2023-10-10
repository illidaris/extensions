package dto

import (
	"net/url"

	"github.com/spf13/cast"
)

// IBizRequest business tag reqquest
type IBizRequest interface {
	GetBizId() int64
	ToUrlQuery() url.Values
}

var _ = IBizRequest(&BizRequest{}) // check impl

// BizRequest default business tag reqquest
type BizRequest struct {
	BizId int64 `json:"bizId" form:"bizId" uri:"bizId"`
}

func (r *BizRequest) GetBizId() int64 {
	return r.BizId
}

func (r *BizRequest) ToUrlQuery() url.Values {
	u := url.Values{}
	u.Add("bizId", cast.ToString(r.BizId))
	return u
}
