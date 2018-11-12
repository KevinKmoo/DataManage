package request

type PagingDto struct {
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
}
