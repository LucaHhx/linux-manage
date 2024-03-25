package response

type PageResult struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"pageSize"`
}

type DxDataGridPageResult struct {
	Data       interface{} `json:"data"`
	TotalCount int64       `json:"totalCount"`
	GroupCount int         `json:"groupCount"`
}
