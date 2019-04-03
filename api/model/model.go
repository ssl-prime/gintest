package model

//KeyValue ...
type KeyValue struct {
	Topic string `json:"topic" valid:"required"` //database name
	Key   string `json:"key" valid:"required"`
	Value string `json:"value" valid:"required"`
}

//Resp ...
type Resp struct {
	Error  string      `json:"error"`
	Msg    interface{} `json:"msg"`
	Status int         `json:"status"`
}

//Delete ...
type Delete struct {
	Topic string `json:"topic" valid:"required"`
	Key   string `json:"key" valid:"required"`
}

//GetValue ...
type GetValue struct {
	Topic string `json:"topic" valid:"required"`
	Key   string `json:"key" valid:"required"`
}
