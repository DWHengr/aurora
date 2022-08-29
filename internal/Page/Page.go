package Page

const (
	Asc  = "asc"
	Desc = "desc"
)

type ReqPage struct {
	Page    int          `json:"page"`
	Size    int          `json:"size"`
	Orders  []OrderItem  `json:"orders"`
	Filters []FilterItem `json:"filters"`
}

type FilterItem struct {
	Column string `json:"column"`
	Value  string `json:"value"`
}

type OrderItem struct {
	Column    string `json:"column"`
	Direction string `json:"direc"` // asc|desc
}

type RespPage struct {
	Page     int         `json:"page"`
	Size     int         `json:"size"`
	Total    int64       `json:"total"`
	DataList interface{} `json:"dataList"`
}
