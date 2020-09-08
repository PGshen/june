package req

// 分页请求条件
type ReqCond struct {
	Size   int32                  `json:"size"`
	Page   int32                  `json:"page"`
	Sort   string                 `json:"sort"`
	Filter map[string]interface{} `json:"filter"`
}
