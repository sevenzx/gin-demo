// Author:      xuan
// Date:        2023/6/25
// Description:	翻页数据

package vo

type PageVO[T any] struct {
	Records  []T   `json:"records"`
	Current  int   `json:"current"`
	PageSize int   `json:"pageSize"`
	Total    int64 `json:"total"`
}
