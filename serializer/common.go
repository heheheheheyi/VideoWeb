package serializer

// Response 基础序列化器
type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Error  string      `json:"error"`
	Msg    string      `json:"msg"`
}

// DataList 基础列表结构
type DataList struct {
	Items interface{} `json:"items"`
	Total uint        `json:"total"`
}

// BuildListResponse 列表构建器
func BuildListResponse(items interface{}, total uint) Response {
	return Response{
		Status: 200,
		Data: DataList{
			Items: items,
			Total: total,
		},
	}
}
func BuildResponse(item interface{}) Response {
	return Response{
		Status: 200,
		Data:   item,
	}
}
