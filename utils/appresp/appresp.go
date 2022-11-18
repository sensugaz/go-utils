package appresp

type Response struct {
	Data interface{} `json:"data,omitempty"`
	Meta interface{} `json:"meta,omitempty"`
}

func NewResponse(data interface{}, meta interface{}) *Response {
	return &Response{
		Data: data,
		Meta: meta,
	}
}
