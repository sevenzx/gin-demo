// Author:      xuan
// Date:        2023/6/25
// Description:	page公共部分

package common

type PageRequest struct {
	Current   int  `json:"current"`
	PageSize  int  `json:"pageSize"`
	NeedTotal bool `json:"needTotal"`
}

// GetDefaultPageRequest 获取pr的默认值
func GetDefaultPageRequest() PageRequest {
	var pr = PageRequest{
		Current:   1,
		PageSize:  10,
		NeedTotal: true,
	}
	return pr
}
