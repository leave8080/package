package main

type RequestBody struct {
	Token string `json:"token"`
}

func NewRequestBody(token string) *RequestBody {
	return &RequestBody{
		Token: token,
	}
}

type UmtrackRsp struct {
	Code      int    `json:"code"`
	Message   string `json:"message"`
	Data      *Info  `json:"data"`
	RequestId string `json:"requestId"`
	Success   bool   `json:"success"`
}

type Info struct {
	Mobile      string  `json:"mobile"`
	Score       float32 `json:"score"`
	ActiveScore float32 `json:"activeScore"`
}
