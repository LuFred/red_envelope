package redenvelope

//CreateRequest 创建红包Request
type CreateRequest struct {
	Amount int32 `json:"amount" required:"true" description:"金额[精确到分的整数]"`
	Count  int32 `json:"count" required:"true" description:"红包个数"`
}

//CreateResponse 创建红包response
type CreateResponse struct {
	ID         int32  `json:"id" description:"红包id"`
	SecretCode string `json:"secret_code" description:"红包口令"`
	Amount     int32  `json:"amount" description:"金额[精确到分的整数]"`
	Count      int32  `json:"count" description:"红包个数"`
}
