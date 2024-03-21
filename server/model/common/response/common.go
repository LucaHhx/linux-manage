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
	GroupCount int64       `json:"groupCount"`
}

type GroupList struct {
	Key     string      `json:"key"`
	Count   int         `json:"count"`
	Items   interface{} `json:"items"`
	Summary interface{} `json:"summary"`
}
